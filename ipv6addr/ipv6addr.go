/*
 * Copyright 2017 agwlvssainokuni
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ipv6addr

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

var ipv6pattern = "^" + //
	"(?i)" + // [正規表現のオプション] 大文字小文字を区別しない
	"(" +
	// (1) IPv4混在なし
	"(" +
	// ・省略なし
	"[0-9a-f]{1,4}(:[0-9a-f]{1,4}){7}" +
	// ・全省略
	"|::" +
	// ・前省略
	"|:(:[0-9a-f]{1,4}){1,7}" +
	// ・後省略
	"|([0-9a-f]{1,4}:){1,7}:" +
	// ・中省略
	"|([0-9a-f]{1,4}:){1}(:[0-9a-f]{1,4}){1,6}" + //
	"|([0-9a-f]{1,4}:){2}(:[0-9a-f]{1,4}){1,5}" + //
	"|([0-9a-f]{1,4}:){3}(:[0-9a-f]{1,4}){1,4}" + //
	"|([0-9a-f]{1,4}:){4}(:[0-9a-f]{1,4}){1,3}" + //
	"|([0-9a-f]{1,4}:){5}(:[0-9a-f]{1,4}){1,2}" + //
	"|([0-9a-f]{1,4}:){6}(:[0-9a-f]{1,4}){1}" + //
	")" +
	// (2) IPv4混在あり]
	"|(" +
	// ・省略なし
	"[0-9a-f]{1,4}(:[0-9a-f]{1,4}){5}" +
	// ・全省略
	"|:" +
	// ・前省略
	"|:(:[0-9a-f]{1,4}){1,5}" +
	// ・後省略
	"|([0-9a-f]{1,4}:){1,5}" +
	// ・中省略
	"|([0-9a-f]{1,4}:){1}(:[0-9a-f]{1,4}){1,4}" + //
	"|([0-9a-f]{1,4}:){2}(:[0-9a-f]{1,4}){1,3}" + //
	"|([0-9a-f]{1,4}:){3}(:[0-9a-f]{1,4}){1,2}" + //
	"|([0-9a-f]{1,4}:){4}(:[0-9a-f]{1,4}){1}" +
	// ・共通末尾 (IPv4部)
	"):(25[0-5]|2[0-4][0-9]|[01]?[0-9]{1,2})(\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9]{1,2})){3}" + //
	")" +
	//
	"$"

var ipv6regexp *regexp.Regexp

var tailv4pattern = "^.*:[0-9]{1,3}(\\.[0-9]{1,3}){3}$"
var tailv4regexp *regexp.Regexp

func init() {
	ipv6regexp = regexp.MustCompile(ipv6pattern)
	tailv4regexp = regexp.MustCompile(tailv4pattern)
}

func IsIpv6Addr(addr string) bool {
	return ipv6regexp.MatchString(addr)
}

func Ipv6Bytes(addr string) ([16]byte, error) {
	result := [16]byte{}
	v6, v4 := makeIpv6Fragment(addr)
	for i, p := range v6 {
		if isZero(p) {
			result[i*2] = 0
			result[i*2+1] = 0
		} else if u, err := strconv.ParseUint(p, 16, 8*2); err != nil {
			return result, err
		} else {
			result[i*2] = byte(u >> 8)
			result[i*2+1] = byte(u & 0x00ff)
		}
	}
	if v4 != "" {
		for i, p := range strings.Split(v4, ".") {
			if u, err := strconv.ParseUint(p, 10, 8); err != nil {
				return result, err
			} else {
				result[len(v6)*2+i] = byte(u)
			}
		}
	}
	return result, nil
}

func Ipv6Mask(length uint) (mask [16]byte) {
	mlen := length / 8
	for i := uint(0); i < mlen; i++ {
		mask[i] = 0xFF
	}
	if r := length % 8; r != 0 {
		b := 1<<8 - 1<<(8-r)
		mask[mlen] = byte(b)
	}
	return
}

func DecompressIpv6Addr(addr string) string {

	buf := bytes.NewBuffer(make([]byte, 0, 4*8+7))

	v6, v4 := makeIpv6Fragment(addr)
	constructDecompressed(buf, v6)

	if len(v4) > 0 {
		buf.WriteString(":")
		buf.WriteString(v4)
	}

	return buf.String()
}

func CompressIpv6Addr(addr string) string {

	buf := bytes.NewBuffer(make([]byte, 0, 4*8+7))

	v6, v4 := makeIpv6Fragment(addr)
	begin, end := computeMaxNullRange(v6)
	constructCompressed(buf, v6, begin, end)

	if len(v4) > 0 {
		if buf.Bytes()[buf.Len()-1] != byte(':') {
			buf.WriteString(":")
		}
		buf.WriteString(v4)
	}

	return buf.String()
}

func makeIpv6Fragment(addr string) (v6 []string, v4 string) {
	if tailv4regexp.MatchString(addr) {

		colon := strings.LastIndex(addr, ":")
		v4 = addr[colon+1:]

		addrv6 := addr[:colon]
		if addrv6 == ":" {
			v6 = padding(6, nil, nil)
		} else if strings.HasPrefix(addrv6, "::") {
			v6 = padding(6, nil, strings.Split(addrv6[2:], ":"))
		} else if strings.HasSuffix(addrv6, ":") {
			v6 = padding(6, strings.Split(addrv6[:len(addrv6)-1], ":"), nil)
		} else if strings.Index(addrv6, "::") >= 0 {
			part := strings.Split(addrv6, "::")
			v6 = padding(6, strings.Split(part[0], ":"), strings.Split(part[1], ":"))
		} else {
			v6 = padding(6, strings.Split(addrv6, ":"), nil)
		}
	} else {
		addrv6 := addr
		if addrv6 == "::" {
			v6 = padding(8, nil, nil)
		} else if strings.HasPrefix(addrv6, "::") {
			v6 = padding(8, nil, strings.Split(addrv6[2:], ":"))
		} else if strings.HasSuffix(addrv6, "::") {
			v6 = padding(8, strings.Split(addrv6[:len(addrv6)-2], ":"), nil)
		} else if strings.Index(addrv6, "::") >= 0 {
			part := strings.Split(addrv6, "::")
			v6 = padding(8, strings.Split(part[0], ":"), strings.Split(part[1], ":"))
		} else {
			v6 = padding(8, strings.Split(addrv6, ":"), nil)
		}
	}
	return
}

func padding(size int, prefix, suffix []string) []string {
	result := make([]string, size)
	if prefix != nil {
		copy(result, prefix)
	}
	if suffix != nil {
		copy(result[size-len(suffix):], suffix)
	}
	return result
}

func computeMaxNullRange(field []string) (begin int, end int) {

	curBegin := -1
	curEnd := -1
	begin = -1
	end = -1

	for i, f := range field {

		if !isZero(f) {
			curBegin = -1
			curEnd = -1
			continue
		}

		if curBegin < 0 {
			curBegin = i
			curEnd = i
		} else {
			curEnd = i
		}

		if begin < 0 {
			begin = curBegin
			end = curEnd
		} else {
			if end-begin < curEnd-curBegin {
				begin = curBegin
				end = curEnd
			}
		}
	}

	return
}

func constructDecompressed(buf *bytes.Buffer, field []string) {
	appendDecompressed(buf, field[0])
	for _, f := range field[1:] {
		buf.WriteString(":")
		appendDecompressed(buf, f)
	}
}

func constructCompressed(buf *bytes.Buffer, field []string, begin, end int) {
	if begin == -1 || begin == end {
		// 省略なし
		appendCompressed(buf, field[0])
		for _, f := range field[1:] {
			buf.WriteString(":")
			appendCompressed(buf, f)
		}
	} else if begin == 0 {
		if end == len(field)-1 {
			// 全省略
			buf.WriteString("::")
		} else {
			// 前省略
			buf.WriteString(":")
			for _, f := range field[end+1:] {
				buf.WriteString(":")
				appendCompressed(buf, f)
			}
		}
	} else {
		if end == len(field)-1 {
			// 後省略
			for _, f := range field[:begin] {
				appendCompressed(buf, f)
				buf.WriteString(":")
			}
			buf.WriteString(":")
		} else {
			// 中省略
			for _, f := range field[:begin] {
				appendCompressed(buf, f)
				buf.WriteString(":")
			}
			for _, f := range field[end+1:] {
				buf.WriteString(":")
				appendCompressed(buf, f)
			}
		}
	}
}

func appendDecompressed(buf *bytes.Buffer, text string) {
	if isZero(text) {
		buf.WriteString("0000")
	} else {
		for i := 0; i < 4-len(text); i++ {
			buf.WriteString("0")
		}
		buf.WriteString(text)
	}
}

func appendCompressed(buf *bytes.Buffer, text string) {
	if isZero(text) {
		buf.WriteString("0")
	} else {
		nzIndex := 0
		for i, ch := range text {
			if ch != '0' {
				nzIndex = i
				break
			}
		}
		buf.WriteString(text[nzIndex:])
	}
}

func isZero(text string) bool {
	if len(text) <= 0 {
		return true
	}
	for _, ch := range text {
		if ch != '0' {
			return false
		}
	}
	return true
}
