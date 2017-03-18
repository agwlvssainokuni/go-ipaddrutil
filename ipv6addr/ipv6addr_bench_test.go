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

func BenchmarkComp10(b *testing.B)     { comp(b, 10) }
func BenchmarkComp100(b *testing.B)    { comp(b, 100) }
func BenchmarkComp1000(b *testing.B)   { comp(b, 1000) }
func BenchmarkComp10000(b *testing.B)  { comp(b, 10000) }
func BenchmarkComp100000(b *testing.B) { comp(b, 100000) }

func BenchmarkDecomp10(b *testing.B)     { decomp(b, 10) }
func BenchmarkDecomp100(b *testing.B)    { decomp(b, 100) }
func BenchmarkDecomp1000(b *testing.B)   { decomp(b, 1000) }
func BenchmarkDecomp10000(b *testing.B)  { decomp(b, 10000) }
func BenchmarkDecomp100000(b *testing.B) { decomp(b, 100000) }

func comp(b *testing.B, count int) {
	addr := "1234:0::000:abcd"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for j := 0; j < count; j++ {
			CompressIpv6Addr(addr)
		}
	}
}

func decomp(b *testing.B, count int) {
	addr := "1234:0::000:abcd"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for j := 0; j < count; j++ {
			DecompressIpv6Addr(addr)
		}
	}
}
