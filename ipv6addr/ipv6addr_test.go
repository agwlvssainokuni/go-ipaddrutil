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
	"testing"
)

func TestIsIpv6Addr(t *testing.T) {
	assertIsIpv6Addr(t, false, "")
	assertIsIpv6Addr(t, true, "2001:0db8:0020:0003:1000:0100:0020:0003")
	assertIsIpv6Addr(t, true, "2001:db8:20:3:1000:100:20:3")
	assertIsIpv6Addr(t, true, "::")
	assertIsIpv6Addr(t, true, "::1.1.1.1")
}

func TestIsIpv6Addr_省略無し(t *testing.T) {
	assertIsIpv6Addr(t, true, "0:0:0:0:0:0:0:0")
	assertIsIpv6Addr(t, true, "00:00:00:00:00:00:00:00")
	assertIsIpv6Addr(t, true, "000:000:000:000:000:000:000:000")
	assertIsIpv6Addr(t, true, "0000:0000:0000:0000:0000:0000:0000:0000")
	assertIsIpv6Addr(t, true, "1111:1111:1111:1111:1111:1111:1111:1111")
	assertIsIpv6Addr(t, true, "2222:2222:2222:2222:2222:2222:2222:2222")
	assertIsIpv6Addr(t, true, "3333:3333:3333:3333:3333:3333:3333:3333")
	assertIsIpv6Addr(t, true, "4444:4444:4444:4444:4444:4444:4444:4444")
	assertIsIpv6Addr(t, true, "5555:5555:5555:5555:5555:5555:5555:5555")
	assertIsIpv6Addr(t, true, "6666:6666:6666:6666:6666:6666:6666:6666")
	assertIsIpv6Addr(t, true, "7777:7777:7777:7777:7777:7777:7777:7777")
	assertIsIpv6Addr(t, true, "8888:8888:8888:8888:8888:8888:8888:8888")
	assertIsIpv6Addr(t, true, "9999:9999:9999:9999:9999:9999:9999:9999")
	assertIsIpv6Addr(t, true, "aaaa:aaaa:aaaa:aaaa:aaaa:aaaa:aaaa:aaaa")
	assertIsIpv6Addr(t, true, "bbbb:bbbb:bbbb:bbbb:bbbb:bbbb:bbbb:bbbb")
	assertIsIpv6Addr(t, true, "cccc:cccc:cccc:cccc:cccc:cccc:cccc:cccc")
	assertIsIpv6Addr(t, true, "dddd:dddd:dddd:dddd:dddd:dddd:dddd:dddd")
	assertIsIpv6Addr(t, true, "eeee:eeee:eeee:eeee:eeee:eeee:eeee:eeee")
	assertIsIpv6Addr(t, true, "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
	assertIsIpv6Addr(t, true, "AAAA:AAAA:AAAA:AAAA:AAAA:AAAA:AAAA:AAAA")
	assertIsIpv6Addr(t, true, "BBBB:BBBB:BBBB:BBBB:BBBB:BBBB:BBBB:BBBB")
	assertIsIpv6Addr(t, true, "CCCC:CCCC:CCCC:CCCC:CCCC:CCCC:CCCC:CCCC")
	assertIsIpv6Addr(t, true, "DDDD:DDDD:DDDD:DDDD:DDDD:DDDD:DDDD:DDDD")
	assertIsIpv6Addr(t, true, "EEEE:EEEE:EEEE:EEEE:EEEE:EEEE:EEEE:EEEE")
	assertIsIpv6Addr(t, true, "FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF")
	assertIsIpv6Addr(t, false, "00000:0:0:0:0:0:0:0")
	assertIsIpv6Addr(t, false, "0:00000:0:0:0:0:0:0")
	assertIsIpv6Addr(t, false, "0:0:00000:0:0:0:0:0")
	assertIsIpv6Addr(t, false, "0:0:0:00000:0:0:0:0")
	assertIsIpv6Addr(t, false, "0:0:0:0:00000:0:0:0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:00000:0:0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:00000:0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:0:00000")
	assertIsIpv6Addr(t, false, "0")
	assertIsIpv6Addr(t, false, "0:")
	assertIsIpv6Addr(t, false, "0:0")
	assertIsIpv6Addr(t, false, "0:0:")
	assertIsIpv6Addr(t, false, "0:0:0")
	assertIsIpv6Addr(t, false, "0:0:0:")
	assertIsIpv6Addr(t, false, "0:0:0:0")
	assertIsIpv6Addr(t, false, "0:0:0:0:")
	assertIsIpv6Addr(t, false, "0:0:0:0:0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:0:")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:0:0:")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:0:0:0")
	assertIsIpv6Addr(t, false, "z:z:z:z:z:z:z:z")
	assertIsIpv6Addr(t, false, "z:0:0:0:0:0:0:0")
	assertIsIpv6Addr(t, false, "0:z:0:0:0:0:0:0")
	assertIsIpv6Addr(t, false, "0:0:z:0:0:0:0:0")
	assertIsIpv6Addr(t, false, "0:0:0:z:0:0:0:0")
	assertIsIpv6Addr(t, false, "0:0:0:0:z:0:0:0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:z:0:0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:z:0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:0:z")
}

func TestIsIpv6Addr_前省略(t *testing.T) {
	assertIsIpv6Addr(t, true, "::1")
	assertIsIpv6Addr(t, true, "::1:1")
	assertIsIpv6Addr(t, true, "::1:1:1")
	assertIsIpv6Addr(t, true, "::1:1:1:1")
	assertIsIpv6Addr(t, true, "::1:1:1:1:1")
	assertIsIpv6Addr(t, true, "::1:1:1:1:1:1")
	assertIsIpv6Addr(t, true, "::1:1:1:1:1:1:1")
	assertIsIpv6Addr(t, false, "::1:1:1:1:1:1:1:1")
	assertIsIpv6Addr(t, false, "::z:z:z:z:z:z:z")
	assertIsIpv6Addr(t, false, "::z:1:1:1:1:1:1")
	assertIsIpv6Addr(t, false, "::1:z:1:1:1:1:1")
	assertIsIpv6Addr(t, false, "::1:1:z:1:1:1:1")
	assertIsIpv6Addr(t, false, "::1:1:1:z:1:1:1")
	assertIsIpv6Addr(t, false, "::1:1:1:1:z:1:1")
	assertIsIpv6Addr(t, false, "::1:1:1:1:1:z:1")
	assertIsIpv6Addr(t, false, "::1:1:1:1:1:1:z")
}

func TestIsIpv6Addr_後省略(t *testing.T) {
	assertIsIpv6Addr(t, true, "1::")
	assertIsIpv6Addr(t, true, "1:1::")
	assertIsIpv6Addr(t, true, "1:1:1::")
	assertIsIpv6Addr(t, true, "1:1:1:1::")
	assertIsIpv6Addr(t, true, "1:1:1:1:1::")
	assertIsIpv6Addr(t, true, "1:1:1:1:1:1::")
	assertIsIpv6Addr(t, true, "1:1:1:1:1:1:1::")
	assertIsIpv6Addr(t, false, "1:1:1:1:1:1:1:1::")
	assertIsIpv6Addr(t, false, "z:z:z:z:z:z:z::")
	assertIsIpv6Addr(t, false, "z:1:1:1:1:1:1::")
	assertIsIpv6Addr(t, false, "1:z:1:1:1:1:1::")
	assertIsIpv6Addr(t, false, "1:1:z:1:1:1:1::")
	assertIsIpv6Addr(t, false, "1:1:1:z:1:1:1::")
	assertIsIpv6Addr(t, false, "1:1:1:1:z:1:1::")
	assertIsIpv6Addr(t, false, "1:1:1:1:1:z:1::")
	assertIsIpv6Addr(t, false, "1:1:1:1:1:1:z::")
}

func TestIsIpv6Addr_中省略(t *testing.T) {

	assertIsIpv6Addr(t, true, "1::1")
	assertIsIpv6Addr(t, true, "1::1:1")
	assertIsIpv6Addr(t, true, "1::1:1:1")
	assertIsIpv6Addr(t, true, "1::1:1:1:1")
	assertIsIpv6Addr(t, true, "1::1:1:1:1:1")
	assertIsIpv6Addr(t, true, "1::1:1:1:1:1:1")
	assertIsIpv6Addr(t, false, "1::1:1:1:1:1:1:1")
	assertIsIpv6Addr(t, false, "z::z:z:z:z:z:z")
	assertIsIpv6Addr(t, false, "z::1:1:1:1:1:1")
	assertIsIpv6Addr(t, false, "1::z:1:1:1:1:1")
	assertIsIpv6Addr(t, false, "1::1:z:1:1:1:1")
	assertIsIpv6Addr(t, false, "1::1:1:z:1:1:1")
	assertIsIpv6Addr(t, false, "1::1:1:1:z:1:1")
	assertIsIpv6Addr(t, false, "1::1:1:1:1:z:1")
	assertIsIpv6Addr(t, false, "1::1:1:1:1:1:z")

	assertIsIpv6Addr(t, true, "2:2::2")
	assertIsIpv6Addr(t, true, "2:2::2:2")
	assertIsIpv6Addr(t, true, "2:2::2:2:2")
	assertIsIpv6Addr(t, true, "2:2::2:2:2:2")
	assertIsIpv6Addr(t, true, "2:2::2:2:2:2:2")
	assertIsIpv6Addr(t, false, "2:2::2:2:2:2:2:2")
	assertIsIpv6Addr(t, false, "z:z::z:z:z:z:z")
	assertIsIpv6Addr(t, false, "z:2::2:2:2:2:2")
	assertIsIpv6Addr(t, false, "2:z::2:2:2:2:2")
	assertIsIpv6Addr(t, false, "2:2::z:2:2:2:2")
	assertIsIpv6Addr(t, false, "2:2::2:z:2:2:2")
	assertIsIpv6Addr(t, false, "2:2::2:2:z:2:2")
	assertIsIpv6Addr(t, false, "2:2::2:2:2:z:2")
	assertIsIpv6Addr(t, false, "2:2::2:2:2:2:z")

	assertIsIpv6Addr(t, true, "3:3:3::3")
	assertIsIpv6Addr(t, true, "3:3:3::3:3")
	assertIsIpv6Addr(t, true, "3:3:3::3:3:3")
	assertIsIpv6Addr(t, true, "3:3:3::3:3:3:3")
	assertIsIpv6Addr(t, false, "3:3:3::3:3:3:3:3")
	assertIsIpv6Addr(t, false, "z:z:z::z:z:z:z")
	assertIsIpv6Addr(t, false, "z:3:3::3:3:3:3")
	assertIsIpv6Addr(t, false, "3:z:3::3:3:3:3")
	assertIsIpv6Addr(t, false, "3:3:z::3:3:3:3")
	assertIsIpv6Addr(t, false, "3:3:3::z:3:3:3")
	assertIsIpv6Addr(t, false, "3:3:3::3:z:3:3")
	assertIsIpv6Addr(t, false, "3:3:3::3:3:z:3")
	assertIsIpv6Addr(t, false, "3:3:3::3:3:3:z")

	assertIsIpv6Addr(t, true, "4:4:4:4::4")
	assertIsIpv6Addr(t, true, "4:4:4:4::4:4")
	assertIsIpv6Addr(t, true, "4:4:4:4::4:4:4")
	assertIsIpv6Addr(t, false, "4:4:4:4::4:4:4:4")
	assertIsIpv6Addr(t, false, "z:z:z:z::z:z:z")
	assertIsIpv6Addr(t, false, "z:4:4:4::4:4:4")
	assertIsIpv6Addr(t, false, "4:z:4:4::4:4:4")
	assertIsIpv6Addr(t, false, "4:4:z:4::4:4:4")
	assertIsIpv6Addr(t, false, "4:4:4:z::4:4:4")
	assertIsIpv6Addr(t, false, "4:4:4:4::z:4:4")
	assertIsIpv6Addr(t, false, "4:4:4:4::4:z:4")
	assertIsIpv6Addr(t, false, "4:4:4:4::4:4:z")

	assertIsIpv6Addr(t, true, "5:5:5:5:5::5")
	assertIsIpv6Addr(t, true, "5:5:5:5:5::5:5")
	assertIsIpv6Addr(t, false, "5:5:5:5:5::5:5:5")
	assertIsIpv6Addr(t, false, "z:z:z:z:z::z:z")
	assertIsIpv6Addr(t, false, "z:5:5:5:5::5:5")
	assertIsIpv6Addr(t, false, "5:z:5:5:5::5:5")
	assertIsIpv6Addr(t, false, "5:5:z:5:5::5:5")
	assertIsIpv6Addr(t, false, "5:5:5:z:5::5:5")
	assertIsIpv6Addr(t, false, "5:5:5:5:z::5:5")
	assertIsIpv6Addr(t, false, "5:5:5:5:5::z:5")
	assertIsIpv6Addr(t, false, "5:5:5:5:5::5:z")

	assertIsIpv6Addr(t, true, "6:6:6:6:6:6::6")
	assertIsIpv6Addr(t, false, "6:6:6:6:6:6::6:6")
	assertIsIpv6Addr(t, false, "z:z:z:z:z:z::z")
	assertIsIpv6Addr(t, false, "z:6:6:6:6:6::6")
	assertIsIpv6Addr(t, false, "6:z:6:6:6:6::6")
	assertIsIpv6Addr(t, false, "6:6:z:6:6:6::6")
	assertIsIpv6Addr(t, false, "6:6:6:z:6:6::6")
	assertIsIpv6Addr(t, false, "6:6:6:6:z:6::6")
	assertIsIpv6Addr(t, false, "6:6:6:6:6:z::6")
	assertIsIpv6Addr(t, false, "6:6:6:6:6:6::z")

	assertIsIpv6Addr(t, false, "7:7:7:7:7:7:7::7")
}

func TestIsIpv6Addr_後IPv6省略無し(t *testing.T) {
	assertIsIpv6Addr(t, true, "0:0:0:0:0:0:0.0.0.0")
	assertIsIpv6Addr(t, true, "00:00:00:00:00:00:00.00.00.00")
	assertIsIpv6Addr(t, true, "000:000:000:000:000:000:000.000.000.000")
	assertIsIpv6Addr(t, true, "0000:0000:0000:0000:0000:0000:000.000.000.000")
	assertIsIpv6Addr(t, true, "1111:1111:1111:1111:1111:1111:1.1.1.1")
	assertIsIpv6Addr(t, true, "2222:2222:2222:2222:2222:2222:2.2.2.2")
	assertIsIpv6Addr(t, true, "3333:3333:3333:3333:3333:3333:3.3.3.3")
	assertIsIpv6Addr(t, true, "4444:4444:4444:4444:4444:4444:4.4.4.4")
	assertIsIpv6Addr(t, true, "5555:5555:5555:5555:5555:5555:5.5.5.5")
	assertIsIpv6Addr(t, true, "6666:6666:6666:6666:6666:6666:6.6.6.6")
	assertIsIpv6Addr(t, true, "7777:7777:7777:7777:7777:7777:7.7.7.7")
	assertIsIpv6Addr(t, true, "8888:8888:8888:8888:8888:8888:8.8.8.8")
	assertIsIpv6Addr(t, true, "9999:9999:9999:9999:9999:9999:9.9.9.9")
	assertIsIpv6Addr(t, true, "aaaa:aaaa:aaaa:aaaa:aaaa:aaaa:10.10.10.10")
	assertIsIpv6Addr(t, true, "bbbb:bbbb:bbbb:bbbb:bbbb:bbbb:11.11.11.11")
	assertIsIpv6Addr(t, true, "cccc:cccc:cccc:cccc:cccc:cccc:12.12.12.12")
	assertIsIpv6Addr(t, true, "dddd:dddd:dddd:dddd:dddd:dddd:13.13.13.13")
	assertIsIpv6Addr(t, true, "eeee:eeee:eeee:eeee:eeee:eeee:14.14.14.14")
	assertIsIpv6Addr(t, true, "ffff:ffff:ffff:ffff:ffff:ffff:15.15.15.15")
	assertIsIpv6Addr(t, true, "AAAA:AAAA:AAAA:AAAA:AAAA:AAAA:10.10.10.10")
	assertIsIpv6Addr(t, true, "BBBB:BBBB:BBBB:BBBB:BBBB:BBBB:11.11.11.11")
	assertIsIpv6Addr(t, true, "CCCC:CCCC:CCCC:CCCC:CCCC:CCCC:12.12.12.12")
	assertIsIpv6Addr(t, true, "DDDD:DDDD:DDDD:DDDD:DDDD:DDDD:13.13.13.13")
	assertIsIpv6Addr(t, true, "EEEE:EEEE:EEEE:EEEE:EEEE:EEEE:14.14.14.14")
	assertIsIpv6Addr(t, true, "FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:15.15.15.15")
	assertIsIpv6Addr(t, false, "00000:0:0:0:0:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:00000:0:0:0:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:00000:0:0:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:00000:0:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:00000:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:00000:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:256.256.256.256")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:256.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:0.256.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:0.0.256.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:0.0.0.256")
	assertIsIpv6Addr(t, false, "0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "z:z:z:z:z:z:z.z.z.z")
	assertIsIpv6Addr(t, false, "z:0:0:0:0:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:z:0:0:0:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:z:0:0:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:z:0:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:z:0:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:z:0.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:z.0.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:0.z.0.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:0.0.z.0")
	assertIsIpv6Addr(t, false, "0:0:0:0:0:0:0.0.0.z")
}

func TestIsIpv6Addr_後IPv4前省略(t *testing.T) {
	assertIsIpv6Addr(t, true, "::1:1.1.1.1")
	assertIsIpv6Addr(t, true, "::1:1:1.1.1.1")
	assertIsIpv6Addr(t, true, "::1:1:1:1.1.1.1")
	assertIsIpv6Addr(t, true, "::1:1:1:1:1.1.1.1")
	assertIsIpv6Addr(t, true, "::1:1:1:1:1:1.1.1.1")
	assertIsIpv6Addr(t, false, "::1:1:1:1:1:1:1.1.1.1")
	assertIsIpv6Addr(t, false, "::z:z:z:z:z:1.1.1.1")
	assertIsIpv6Addr(t, false, "::z:1:1:1:1:1.1.1.1")
	assertIsIpv6Addr(t, false, "::1:z:1:1:1:1.1.1.1")
	assertIsIpv6Addr(t, false, "::1:1:z:1:1:1.1.1.1")
	assertIsIpv6Addr(t, false, "::1:1:1:z:1:1.1.1.1")
	assertIsIpv6Addr(t, false, "::1:1:1:1:z:1.1.1.1")
}

func TestIsIpv6Addr_後IPv4後省略(t *testing.T) {
	assertIsIpv6Addr(t, true, "1::1.1.1.1")
	assertIsIpv6Addr(t, true, "1:1::1.1.1.1")
	assertIsIpv6Addr(t, true, "1:1:1::1.1.1.1")
	assertIsIpv6Addr(t, true, "1:1:1:1::1.1.1.1")
	assertIsIpv6Addr(t, true, "1:1:1:1:1::1.1.1.1")
	assertIsIpv6Addr(t, false, "1:1:1:1:1:1::1.1.1.1")
	assertIsIpv6Addr(t, false, "z:z:z:z:z::1.1.1.1")
	assertIsIpv6Addr(t, false, "z:1:1:1:1::1.1.1.1")
	assertIsIpv6Addr(t, false, "1:z:1:1:1::1.1.1.1")
	assertIsIpv6Addr(t, false, "1:1:z:1:1::1.1.1.1")
	assertIsIpv6Addr(t, false, "1:1:1:z:1::1.1.1.1")
	assertIsIpv6Addr(t, false, "1:1:1:1:z::1.1.1.1")
}

func TestIsIpv6Addr_後IPv4中省略(t *testing.T) {

	assertIsIpv6Addr(t, true, "1::1:1.1.1.1")
	assertIsIpv6Addr(t, true, "1::1:1:1.1.1.1")
	assertIsIpv6Addr(t, true, "1::1:1:1:1.1.1.1")
	assertIsIpv6Addr(t, true, "1::1:1:1:1:1.1.1.1")
	assertIsIpv6Addr(t, false, "1::1:1:1:1:1:1.1.1.1")
	assertIsIpv6Addr(t, false, "z::z:z:z:z:1.1.1.1")
	assertIsIpv6Addr(t, false, "z::1:1:1:1:1.1.1.1")
	assertIsIpv6Addr(t, false, "1::z:1:1:1:1.1.1.1")
	assertIsIpv6Addr(t, false, "1::1:z:1:1:1.1.1.1")
	assertIsIpv6Addr(t, false, "1::1:1:z:1:1.1.1.1")
	assertIsIpv6Addr(t, false, "1::1:1:1:z:1.1.1.1")

	assertIsIpv6Addr(t, true, "2:2::2:2.2.2.2")
	assertIsIpv6Addr(t, true, "2:2::2:2:2.2.2.2")
	assertIsIpv6Addr(t, true, "2:2::2:2:2:2.2.2.2")
	assertIsIpv6Addr(t, false, "2:2::2:2:2:2:2.2.2.2")
	assertIsIpv6Addr(t, false, "z:z::z:z:z:2.2.2.2")
	assertIsIpv6Addr(t, false, "z:2::2:2:2:2.2.2.2")
	assertIsIpv6Addr(t, false, "2:z::2:2:2:2.2.2.2")
	assertIsIpv6Addr(t, false, "2:2::z:2:2:2.2.2.2")
	assertIsIpv6Addr(t, false, "2:2::2:z:2:2.2.2.2")
	assertIsIpv6Addr(t, false, "2:2::2:2:z:2.2.2.2")

	assertIsIpv6Addr(t, true, "3:3:3::3:3.3.3.3")
	assertIsIpv6Addr(t, true, "3:3:3::3:3:3.3.3.3")
	assertIsIpv6Addr(t, false, "3:3:3::3:3:3:3.3.3.3")
	assertIsIpv6Addr(t, false, "z:z:z::z:z:3.3.3.3")
	assertIsIpv6Addr(t, false, "z:3:3::3:3:3.3.3.3")
	assertIsIpv6Addr(t, false, "3:z:3::3:3:3.3.3.3")
	assertIsIpv6Addr(t, false, "3:3:z::3:3:3.3.3.3")
	assertIsIpv6Addr(t, false, "3:3:3::z:3:3.3.3.3")
	assertIsIpv6Addr(t, false, "3:3:3::3:z:3.3.3.3")

	assertIsIpv6Addr(t, true, "4:4:4:4::4:4.4.4.4")
	assertIsIpv6Addr(t, false, "4:4:4:4::4:4:4.4.4.4")
	assertIsIpv6Addr(t, false, "z:z:z:z::z:4.4.4.4")
	assertIsIpv6Addr(t, false, "z:4:4:4::4:4.4.4.4")
	assertIsIpv6Addr(t, false, "4:z:4:4::4:4.4.4.4")
	assertIsIpv6Addr(t, false, "4:4:z:4::4:4.4.4.4")
	assertIsIpv6Addr(t, false, "4:4:4:z::4:4.4.4.4")
	assertIsIpv6Addr(t, false, "4:4:4:4::z:4.4.4.4")

	assertIsIpv6Addr(t, false, "5:5:5:5:5::5:5.5.5.5")
}

func TestDecompressIpv6Addr(t *testing.T) {
	assertDecompressIpv6Addr(t, "2001:0db8:0020:0003:1000:0100:0020:0003", "2001:db8:20:3:1000:100:20:3")
	assertDecompressIpv6Addr(t, "2001:0db8:0000:0000:1234:0000:0000:9abc", "2001:db8::1234:0:0:9abc")
	assertDecompressIpv6Addr(t, "2001:0db8:0000:0000:0000:0000:0000:9abc", "2001:db8::9abc")
	assertDecompressIpv6Addr(t, "0001:0000:0000:0001:0000:0000:0000:0001", "1::1:0:0:0:1")
	assertDecompressIpv6Addr(t, "0001:0000:0000:0001:0000:0000:0001:0001", "1:0:0:1::1:1")
}

func TestCompressIpv6Addr(t *testing.T) {
	assertCompressIpv6Addr(t, "2001:db8:20:3:1000:100:20:3", "2001:0db8:0020:0003:1000:0100:0020:0003")
	assertCompressIpv6Addr(t, "2001:db8::1234:0:0:9abc", "2001:0db8:0000:0000:1234:0000:0000:9abc")
	assertCompressIpv6Addr(t, "2001:db8::9abc", "2001:0db8:0000:0000:0000:0000:0000:9abc")
	assertCompressIpv6Addr(t, "1:0:0:1::1", "1::1:0:0:0:1")
	assertCompressIpv6Addr(t, "1::1:0:0:1:1", "1:0:0:1::1:1")
}

func TestIpAddressCompression_全省略(t *testing.T) {
	testIpv6AddressCompression(t, "::", "0000:0000:0000:0000:0000:0000:0000:0000")
}

func TestIpAddressCompression_省略無(t *testing.T) {
	testIpv6AddressCompression(t, "1:2:3:4:5:6:7:8", "0001:0002:0003:0004:0005:0006:0007:0008")
	testIpv6AddressCompression(t, "11:22:33:44:55:66:77:88", "0011:0022:0033:0044:0055:0066:0077:0088")
	testIpv6AddressCompression(t, "111:222:333:444:555:666:777:888", "0111:0222:0333:0444:0555:0666:0777:0888")
	testIpv6AddressCompression(t, "1111:2222:3333:4444:5555:6666:7777:8888", "1111:2222:3333:4444:5555:6666:7777:8888")
}

func TestIpAddressCompression_前省略(t *testing.T) {
	testIpv6AddressCompression(t, "::1", "0000:0000:0000:0000:0000:0000:0000:0001")
	testIpv6AddressCompression(t, "::1:1", "0000:0000:0000:0000:0000:0000:0001:0001")
	testIpv6AddressCompression(t, "::1:1:1", "0000:0000:0000:0000:0000:0001:0001:0001")
	testIpv6AddressCompression(t, "::1:1:1:1", "0000:0000:0000:0000:0001:0001:0001:0001")
	testIpv6AddressCompression(t, "::1:1:1:1:1", "0000:0000:0000:0001:0001:0001:0001:0001")
	testIpv6AddressCompression(t, "::1:1:1:1:1:1", "0000:0000:0001:0001:0001:0001:0001:0001")

	// testIpv6AddressCompression(t, "::1:1:1:1:1:1:1", "0000:0001:0001:0001:0001:0001:0001:0001")
	assertDecompressIpv6Addr(t, "0000:0001:0001:0001:0001:0001:0001:0001", "::1:1:1:1:1:1:1")
	assertDecompressIpv6Addr(t, "0000:0001:0001:0001:0001:0001:0001:0001", "0000:0001:0001:0001:0001:0001:0001:0001")
	assertCompressIpv6Addr(t, "0:1:1:1:1:1:1:1", "0000:0001:0001:0001:0001:0001:0001:0001")
	assertCompressIpv6Addr(t, "0:1:1:1:1:1:1:1", "::1:1:1:1:1:1:1")
}

func TestIpAddressCompression_後省略(t *testing.T) {
	testIpv6AddressCompression(t, "1::", "0001:0000:0000:0000:0000:0000:0000:0000")
	testIpv6AddressCompression(t, "1:1::", "0001:0001:0000:0000:0000:0000:0000:0000")
	testIpv6AddressCompression(t, "1:1:1::", "0001:0001:0001:0000:0000:0000:0000:0000")
	testIpv6AddressCompression(t, "1:1:1:1::", "0001:0001:0001:0001:0000:0000:0000:0000")
	testIpv6AddressCompression(t, "1:1:1:1:1::", "0001:0001:0001:0001:0001:0000:0000:0000")
	testIpv6AddressCompression(t, "1:1:1:1:1:1::", "0001:0001:0001:0001:0001:0001:0000:0000")

	// testIpv6AddressCompression(t, "1:1:1:1:1:1:1::", "0001:0001:0001:0001:0001:0001:0001:0000")
	assertDecompressIpv6Addr(t, "0001:0001:0001:0001:0001:0001:0001:0000", "1:1:1:1:1:1:1::")
	assertDecompressIpv6Addr(t, "0001:0001:0001:0001:0001:0001:0001:0000", "0001:0001:0001:0001:0001:0001:0001:0000")
	assertCompressIpv6Addr(t, "1:1:1:1:1:1:1:0", "0001:0001:0001:0001:0001:0001:0001:0000")
	assertCompressIpv6Addr(t, "1:1:1:1:1:1:1:0", "1:1:1:1:1:1:1::")
}

func TestIpAddressCompression_中省略(t *testing.T) {

	testIpv6AddressCompression(t, "1::1", "0001:0000:0000:0000:0000:0000:0000:0001")
	testIpv6AddressCompression(t, "1::1:1", "0001:0000:0000:0000:0000:0000:0001:0001")
	testIpv6AddressCompression(t, "1::1:1:1", "0001:0000:0000:0000:0000:0001:0001:0001")
	testIpv6AddressCompression(t, "1::1:1:1:1", "0001:0000:0000:0000:0001:0001:0001:0001")
	testIpv6AddressCompression(t, "1::1:1:1:1:1", "0001:0000:0000:0001:0001:0001:0001:0001")
	// testIpv6AddressCompression(t, "1::1:1:1:1:1:1", "0001:0000:0001:0001:0001:0001:0001:0001")
	assertDecompressIpv6Addr(t, "0001:0000:0001:0001:0001:0001:0001:0001", "1::1:1:1:1:1:1")
	assertDecompressIpv6Addr(t, "0001:0000:0001:0001:0001:0001:0001:0001", "0001:0000:0001:0001:0001:0001:0001:0001")
	assertCompressIpv6Addr(t, "1:0:1:1:1:1:1:1", "0001:0000:0001:0001:0001:0001:0001:0001")
	assertCompressIpv6Addr(t, "1:0:1:1:1:1:1:1", "1::1:1:1:1:1:1")

	testIpv6AddressCompression(t, "2:2::2", "0002:0002:0000:0000:0000:0000:0000:0002")
	testIpv6AddressCompression(t, "2:2::2:2", "0002:0002:0000:0000:0000:0000:0002:0002")
	testIpv6AddressCompression(t, "2:2::2:2:2", "0002:0002:0000:0000:0000:0002:0002:0002")
	testIpv6AddressCompression(t, "2:2::2:2:2:2", "0002:0002:0000:0000:0002:0002:0002:0002")
	// testIpv6AddressCompression(t, "2:2::2:2:2:2:2", "0002:0002:0000:0002:0002:0002:0002:0002")
	assertDecompressIpv6Addr(t, "0002:0002:0000:0002:0002:0002:0002:0002", "2:2::2:2:2:2:2")
	assertDecompressIpv6Addr(t, "0002:0002:0000:0002:0002:0002:0002:0002", "0002:0002:0000:0002:0002:0002:0002:0002")
	assertCompressIpv6Addr(t, "2:2:0:2:2:2:2:2", "0002:0002:0000:0002:0002:0002:0002:0002")
	assertCompressIpv6Addr(t, "2:2:0:2:2:2:2:2", "2:2::2:2:2:2:2")

	testIpv6AddressCompression(t, "3:3:3::3", "0003:0003:0003:0000:0000:0000:0000:0003")
	testIpv6AddressCompression(t, "3:3:3::3:3", "0003:0003:0003:0000:0000:0000:0003:0003")
	testIpv6AddressCompression(t, "3:3:3::3:3:3", "0003:0003:0003:0000:0000:0003:0003:0003")
	// testIpv6AddressCompression(t, "3:3:3::3:3:3:3", "0003:0003:0003:0000:0003:0003:0003:0003")
	assertDecompressIpv6Addr(t, "0003:0003:0003:0000:0003:0003:0003:0003", "3:3:3::3:3:3:3")
	assertDecompressIpv6Addr(t, "0003:0003:0003:0000:0003:0003:0003:0003", "0003:0003:0003:0000:0003:0003:0003:0003")
	assertCompressIpv6Addr(t, "3:3:3:0:3:3:3:3", "0003:0003:0003:0000:0003:0003:0003:0003")
	assertCompressIpv6Addr(t, "3:3:3:0:3:3:3:3", "3:3:3::3:3:3:3")

	testIpv6AddressCompression(t, "4:4:4:4::4", "0004:0004:0004:0004:0000:0000:0000:0004")
	testIpv6AddressCompression(t, "4:4:4:4::4:4", "0004:0004:0004:0004:0000:0000:0004:0004")
	// testIpv6AddressCompression(t, "4:4:4:4::4:4:4", "0004:0004:0004:0004:0000:0004:0004:0004")
	assertDecompressIpv6Addr(t, "0004:0004:0004:0004:0000:0004:0004:0004", "4:4:4:4::4:4:4")
	assertDecompressIpv6Addr(t, "0004:0004:0004:0004:0000:0004:0004:0004", "0004:0004:0004:0004:0000:0004:0004:0004")
	assertCompressIpv6Addr(t, "4:4:4:4:0:4:4:4", "0004:0004:0004:0004:0000:0004:0004:0004")
	assertCompressIpv6Addr(t, "4:4:4:4:0:4:4:4", "4:4:4:4::4:4:4")

	testIpv6AddressCompression(t, "5:5:5:5:5::5", "0005:0005:0005:0005:0005:0000:0000:0005")
	// testIpv6AddressCompression(t, "5:5:5:5:5::5:5", "0005:0005:0005:0005:0005:0000:0005:0005")
	assertDecompressIpv6Addr(t, "0005:0005:0005:0005:0005:0000:0005:0005", "5:5:5:5:5::5:5")
	assertDecompressIpv6Addr(t, "0005:0005:0005:0005:0005:0000:0005:0005", "0005:0005:0005:0005:0005:0000:0005:0005")
	assertCompressIpv6Addr(t, "5:5:5:5:5:0:5:5", "0005:0005:0005:0005:0005:0000:0005:0005")
	assertCompressIpv6Addr(t, "5:5:5:5:5:0:5:5", "5:5:5:5:5::5:5")

	// testIpv6AddressCompression(t, "6:6:6:6:6:6::6", "0006:0006:0006:0006:0006:0006:0000:0006")
	assertDecompressIpv6Addr(t, "0006:0006:0006:0006:0006:0006:0000:0006", "6:6:6:6:6:6::6")
	assertDecompressIpv6Addr(t, "0006:0006:0006:0006:0006:0006:0000:0006", "0006:0006:0006:0006:0006:0006:0000:0006")
	assertCompressIpv6Addr(t, "6:6:6:6:6:6:0:6", "0006:0006:0006:0006:0006:0006:0000:0006")
	assertCompressIpv6Addr(t, "6:6:6:6:6:6:0:6", "6:6:6:6:6:6::6")
}

func TestIpAddressCompression_後IPv4全省略(t *testing.T) {
	testIpv6AddressCompression(t, "::1.1.1.1", "0000:0000:0000:0000:0000:0000:1.1.1.1")
}

func TestIpAddressCompression_後IPv4省略無(t *testing.T) {
	testIpv6AddressCompression(t, "1:2:3:4:5:6:1.1.1.1", "0001:0002:0003:0004:0005:0006:1.1.1.1")
	testIpv6AddressCompression(t, "11:22:33:44:55:66:1.1.1.1", "0011:0022:0033:0044:0055:0066:1.1.1.1")
	testIpv6AddressCompression(t, "111:222:333:444:555:666:1.1.1.1", "0111:0222:0333:0444:0555:0666:1.1.1.1")
	testIpv6AddressCompression(t, "1111:2222:3333:4444:5555:6666:1.1.1.1", "1111:2222:3333:4444:5555:6666:1.1.1.1")
}

func TestIpAddressCompression_後IPv4前省略(t *testing.T) {
	testIpv6AddressCompression(t, "::1:1.1.1.1", "0000:0000:0000:0000:0000:0001:1.1.1.1")
	testIpv6AddressCompression(t, "::1:1:1.1.1.1", "0000:0000:0000:0000:0001:0001:1.1.1.1")
	testIpv6AddressCompression(t, "::1:1:1:1.1.1.1", "0000:0000:0000:0001:0001:0001:1.1.1.1")
	testIpv6AddressCompression(t, "::1:1:1:1:1.1.1.1", "0000:0000:0001:0001:0001:0001:1.1.1.1")
	// testIpv6AddressCompression(t, "::1:1:1:1:1:1.1.1.1", "0000:0001:0001:0001:0001:0001:1.1.1.1")
	assertDecompressIpv6Addr(t, "0000:0001:0001:0001:0001:0001:1.1.1.1", "::1:1:1:1:1:1.1.1.1")
	assertDecompressIpv6Addr(t, "0000:0001:0001:0001:0001:0001:1.1.1.1", "0000:0001:0001:0001:0001:0001:1.1.1.1")
	assertCompressIpv6Addr(t, "0:1:1:1:1:1:1.1.1.1", "0000:0001:0001:0001:0001:0001:1.1.1.1")
	assertCompressIpv6Addr(t, "0:1:1:1:1:1:1.1.1.1", "::1:1:1:1:1:1.1.1.1")
}

func TestIpAddressCompression_後IPv4後省略(t *testing.T) {
	testIpv6AddressCompression(t, "1::1.1.1.1", "0001:0000:0000:0000:0000:0000:1.1.1.1")
	testIpv6AddressCompression(t, "1:1::1.1.1.1", "0001:0001:0000:0000:0000:0000:1.1.1.1")
	testIpv6AddressCompression(t, "1:1:1::1.1.1.1", "0001:0001:0001:0000:0000:0000:1.1.1.1")
	testIpv6AddressCompression(t, "1:1:1:1::1.1.1.1", "0001:0001:0001:0001:0000:0000:1.1.1.1")
	// testIpv6AddressCompression(t, "1:1:1:1:1::1.1.1.1", "0001:0001:0001:0001:0001:0000:1.1.1.1")
	assertDecompressIpv6Addr(t, "0001:0001:0001:0001:0001:0000:1.1.1.1", "1:1:1:1:1::1.1.1.1")
	assertDecompressIpv6Addr(t, "0001:0001:0001:0001:0001:0000:1.1.1.1", "0001:0001:0001:0001:0001:0000:1.1.1.1")
	assertCompressIpv6Addr(t, "1:1:1:1:1:0:1.1.1.1", "0001:0001:0001:0001:0001:0000:1.1.1.1")
	assertCompressIpv6Addr(t, "1:1:1:1:1:0:1.1.1.1", "1:1:1:1:1::1.1.1.1")
}

func TestIpAddressCompression_後IPv4中省略(t *testing.T) {

	testIpv6AddressCompression(t, "1::1:1.1.1.1", "0001:0000:0000:0000:0000:0001:1.1.1.1")
	testIpv6AddressCompression(t, "1::1:1:1.1.1.1", "0001:0000:0000:0000:0001:0001:1.1.1.1")
	testIpv6AddressCompression(t, "1::1:1:1:1.1.1.1", "0001:0000:0000:0001:0001:0001:1.1.1.1")
	// testIpv6AddressCompression(t, "1::1:1:1:1:1.1.1.1", "0001:0000:0001:0001:0001:0001:1.1.1.1")
	assertDecompressIpv6Addr(t, "0001:0000:0001:0001:0001:0001:1.1.1.1", "1::1:1:1:1:1.1.1.1")
	assertDecompressIpv6Addr(t, "0001:0000:0001:0001:0001:0001:1.1.1.1", "0001:0000:0001:0001:0001:0001:1.1.1.1")
	assertCompressIpv6Addr(t, "1:0:1:1:1:1:1.1.1.1", "0001:0000:0001:0001:0001:0001:1.1.1.1")
	assertCompressIpv6Addr(t, "1:0:1:1:1:1:1.1.1.1", "1::1:1:1:1:1.1.1.1")

	testIpv6AddressCompression(t, "2:2::2:2.2.2.2", "0002:0002:0000:0000:0000:0002:2.2.2.2")
	testIpv6AddressCompression(t, "2:2::2:2:2.2.2.2", "0002:0002:0000:0000:0002:0002:2.2.2.2")
	// testIpv6AddressCompression(t, "2:2::2:2:2:2.2.2.2", "0002:0002:0000:0002:0002:0002:2.2.2.2")
	assertDecompressIpv6Addr(t, "0002:0002:0000:0002:0002:0002:2.2.2.2", "2:2::2:2:2:2.2.2.2")
	assertDecompressIpv6Addr(t, "0002:0002:0000:0002:0002:0002:2.2.2.2", "0002:0002:0000:0002:0002:0002:2.2.2.2")
	assertCompressIpv6Addr(t, "2:2:0:2:2:2:2.2.2.2", "0002:0002:0000:0002:0002:0002:2.2.2.2")
	assertCompressIpv6Addr(t, "2:2:0:2:2:2:2.2.2.2", "2:2::2:2:2:2.2.2.2")

	testIpv6AddressCompression(t, "3:3:3::3:3.3.3.3", "0003:0003:0003:0000:0000:0003:3.3.3.3")
	// testIpv6AddressCompression(t, "3:3:3::3:3:3.3.3.3", "0003:0003:0003:0000:0003:0003:3.3.3.3")
	assertDecompressIpv6Addr(t, "0003:0003:0003:0000:0003:0003:3.3.3.3", "3:3:3::3:3:3.3.3.3")
	assertDecompressIpv6Addr(t, "0003:0003:0003:0000:0003:0003:3.3.3.3", "0003:0003:0003:0000:0003:0003:3.3.3.3")
	assertCompressIpv6Addr(t, "3:3:3:0:3:3:3.3.3.3", "0003:0003:0003:0000:0003:0003:3.3.3.3")
	assertCompressIpv6Addr(t, "3:3:3:0:3:3:3.3.3.3", "3:3:3::3:3:3.3.3.3")

	// testIpv6AddressCompression(t, "4:4:4:4::4:4.4.4.4", "0004:0004:0004:0004:0000:0004:4.4.4.4")
	assertDecompressIpv6Addr(t, "0004:0004:0004:0004:0000:0004:4.4.4.4", "4:4:4:4::4:4.4.4.4")
	assertDecompressIpv6Addr(t, "0004:0004:0004:0004:0000:0004:4.4.4.4", "0004:0004:0004:0004:0000:0004:4.4.4.4")
	assertCompressIpv6Addr(t, "4:4:4:4:0:4:4.4.4.4", "0004:0004:0004:0004:0000:0004:4.4.4.4")
	assertCompressIpv6Addr(t, "4:4:4:4:0:4:4.4.4.4", "4:4:4:4::4:4.4.4.4")
}

func testIpv6AddressCompression(t *testing.T, comped, decomped string) {
	assertDecompressIpv6Addr(t, decomped, comped)
	assertDecompressIpv6Addr(t, decomped, decomped)
	assertCompressIpv6Addr(t, comped, decomped)
	assertCompressIpv6Addr(t, comped, comped)
}

func assertIsIpv6Addr(t *testing.T, expected bool, addr string) {
	actual := IsIpv6Addr(addr)
	if actual != expected {
		t.Errorf("IsIpv6Addr(%s) must be %v, but %v", addr, expected, actual)
	}
}

func assertDecompressIpv6Addr(t *testing.T, expected, addr string) {
	actual := DecompressIpv6Addr(addr)
	if actual != expected {
		t.Errorf("DecompressIpv6Addr(%s) must be %v, but %v", addr, expected, actual)
	}
}

func assertCompressIpv6Addr(t *testing.T, expected, addr string) {
	actual := CompressIpv6Addr(addr)
	if actual != expected {
		t.Errorf("CompressIpv6Addr(%s) must be %v, but %v", addr, expected, actual)
	}
}
