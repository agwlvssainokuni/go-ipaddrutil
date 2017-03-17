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

package ipv4addr

import (
	"regexp"
)

var ipv4pattern = "^(?i)" + //
	"(25[0-5]|2[0-4][0-9]|[01]?[0-9]{1,2})(\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9]{1,2})){3}" + //
	"$"

var ipv4regexp *regexp.Regexp

func init() {
	ipv4regexp = regexp.MustCompile(ipv4pattern)
}

func IsIpv4Addr(addr string) bool {
	return ipv4regexp.MatchString(addr)
}
