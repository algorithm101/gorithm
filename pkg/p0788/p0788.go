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

package p0788

func rotatedDigits(N int) int {
	_isHaoNumber := func(n int) bool {
		r, tmp, mul := n, 0, 1
		for n != 0 {
			digit := n % 10
			switch digit {
			case 0, 1, 8:
			case 2:
				digit = 5
			case 5:
				digit = 2
			case 6:
				digit = 9
			case 9:
				digit = 6
			default:
				return false
			}
			tmp = digit*mul + tmp
			mul = mul * 10
			n = n / 10
		}

		return r != tmp
	}

	count := 0
	for i := 1; i <= N; i++ {
		if _isHaoNumber(i) {
			count++
		}
	}

	return count
}
