// Copyright (c) 2018 soren yang
//
// Licensed under the MIT License
// you may not use this file except in complicance with the License.
// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0050

// 快速幂
func myPow(x float64, n int) float64 {
	_fastpow := func(x float64, n int) float64 {
		r, base := 1.0, x
		for n != 0 {
			if n&0x01 == 1 {
				r *= base
			}
			base *= base
			n = n >> 1
		}

		return r
	}

	if n == 0 {
		return 1
	}

	if n < 0 {
		r := _fastpow(x, 0-n)
		return 1.0 / r
	}
	return _fastpow(x, n)
}
