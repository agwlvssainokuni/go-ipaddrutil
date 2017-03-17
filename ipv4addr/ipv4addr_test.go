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
	"testing"
)

func TestIsIpv4Addr(t *testing.T) {
	assertIsIpv4Addr(t, true, "0.0.0.0")
	assertIsIpv4Addr(t, true, "00.00.00.00")
	assertIsIpv4Addr(t, true, "000.000.000.000")
	assertIsIpv4Addr(t, true, "255.255.255.255")
	assertIsIpv4Addr(t, true, "127.0.0.1")
	assertIsIpv4Addr(t, true, "10.0.0.1")
	assertIsIpv4Addr(t, true, "192.168.0.1")
	assertIsIpv4Addr(t, false, "")
	assertIsIpv4Addr(t, false, "0")
	assertIsIpv4Addr(t, false, "0.")
	assertIsIpv4Addr(t, false, "0.0")
	assertIsIpv4Addr(t, false, "0.0.")
	assertIsIpv4Addr(t, false, "0.0.0")
	assertIsIpv4Addr(t, false, "0.0.0.")
	assertIsIpv4Addr(t, false, "0000.0.0.0")
	assertIsIpv4Addr(t, false, "0.0000.0.0")
	assertIsIpv4Addr(t, false, "0.0.0000.0")
	assertIsIpv4Addr(t, false, "0.0.0.0000")
	assertIsIpv4Addr(t, false, "z.z.z.z")
	assertIsIpv4Addr(t, false, "z.0.0.0")
	assertIsIpv4Addr(t, false, "0.z.0.0")
	assertIsIpv4Addr(t, false, "0.0.z.0")
	assertIsIpv4Addr(t, false, "0.0.0.z")

	assertIsIpv4Addr(t, true, "0.0.0.0")
	assertIsIpv4Addr(t, true, "9.0.0.0")
	assertIsIpv4Addr(t, true, "00.0.0.0")
	assertIsIpv4Addr(t, true, "99.0.0.0")
	assertIsIpv4Addr(t, true, "000.0.0.0")
	assertIsIpv4Addr(t, true, "099.0.0.0")
	assertIsIpv4Addr(t, true, "100.0.0.0")
	assertIsIpv4Addr(t, true, "199.0.0.0")
	assertIsIpv4Addr(t, true, "200.0.0.0")
	assertIsIpv4Addr(t, true, "209.0.0.0")
	assertIsIpv4Addr(t, true, "240.0.0.0")
	assertIsIpv4Addr(t, true, "249.0.0.0")
	assertIsIpv4Addr(t, true, "250.0.0.0")
	assertIsIpv4Addr(t, true, "255.0.0.0")
	assertIsIpv4Addr(t, false, "256.0.0.0")
	assertIsIpv4Addr(t, false, "260.0.0.0")
	assertIsIpv4Addr(t, false, "270.0.0.0")
	assertIsIpv4Addr(t, false, "280.0.0.0")
	assertIsIpv4Addr(t, false, "290.0.0.0")

	assertIsIpv4Addr(t, true, "0.0.0.0")
	assertIsIpv4Addr(t, true, "0.0.0.9")
	assertIsIpv4Addr(t, true, "0.0.0.00")
	assertIsIpv4Addr(t, true, "0.0.0.99")
	assertIsIpv4Addr(t, true, "0.0.0.000")
	assertIsIpv4Addr(t, true, "0.0.0.099")
	assertIsIpv4Addr(t, true, "0.0.0.100")
	assertIsIpv4Addr(t, true, "0.0.0.199")
	assertIsIpv4Addr(t, true, "0.0.0.200")
	assertIsIpv4Addr(t, true, "0.0.0.209")
	assertIsIpv4Addr(t, true, "0.0.0.240")
	assertIsIpv4Addr(t, true, "0.0.0.249")
	assertIsIpv4Addr(t, true, "0.0.0.250")
	assertIsIpv4Addr(t, true, "0.0.0.255")
	assertIsIpv4Addr(t, false, "0.0.0.256")
	assertIsIpv4Addr(t, false, "0.0.0.260")
	assertIsIpv4Addr(t, false, "0.0.0.270")
	assertIsIpv4Addr(t, false, "0.0.0.280")
	assertIsIpv4Addr(t, false, "0.0.0.290")
}

func assertIsIpv4Addr(t *testing.T, expected bool, addr string) {
	actual := IsIpv4Addr(addr)
	if actual != expected {
		t.Errorf("IsIpv4Addr(%s) must be %v, but %v", addr, expected, actual)
	}
}
