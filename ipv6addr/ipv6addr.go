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
	"regexp"
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

func DecompressIpv6Addr(addr string) string {

	v6, v4 := makeIpv6Fragment(addr)
	builder := constructDecompressed(v6)

	if len(v4) > 0 {
		builder += ":" + v4
	}

	return builder
}

func CompressIpv6Addr(addr string) string {

	v6, v4 := makeIpv6Fragment(addr)
	begin, end := computeMaxNullRange(v6)
	builder := constructCompressed(v6, begin, end)

	if len(v4) > 0 {
		if !strings.HasSuffix(builder, ":") {
			builder += ":"
		}
		builder += v4
	}

	return builder
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

func constructDecompressed(field []string) (result string) {
	result += appendDecompressed(field[0])
	for _, f := range field[1:] {
		result += ":" + appendDecompressed(f)
	}
	return
}

func constructCompressed(field []string, begin, end int) (result string) {
	if begin == -1 || begin == end {
		// 省略なし
		result += appendCompressed(field[0])
		for _, f := range field[1:] {
			result += ":" + appendCompressed(f)
		}
	} else if begin == 0 {
		if end == len(field)-1 {
			// 全省略
			result += "::"
		} else {
			// 前省略
			result += ":"
			for _, f := range field[end+1:] {
				result += ":" + appendCompressed(f)
			}
		}
	} else {
		if end == len(field)-1 {
			// 後省略
			for _, f := range field[:begin] {
				result += appendCompressed(f) + ":"
			}
			result += ":"
		} else {
			// 中省略
			for _, f := range field[:begin] {
				result += appendCompressed(f) + ":"
			}
			for _, f := range field[end+1:] {
				result += ":" + appendCompressed(f)
			}
		}
	}
	return
}

func appendDecompressed(text string) string {
	if isZero(text) {
		return "0000"
	} else {
		return strings.Repeat("0", 4-len(text)) + text
	}
}

func appendCompressed(text string) string {
	if isZero(text) {
		return "0"
	} else {
		return strings.TrimLeft(text, "0")
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
