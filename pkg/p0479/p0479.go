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

package p0479

import (
	"math"
)

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}

	_palindrome := func(n int) int {
		r := n
		for n != 0 {
			r = r*10 + n%10
			n = n / 10
		}

		return r
	}
	upper := int(math.Pow10(n)) - 1
	lower := int(math.Pow10(n - 1))

	for i := upper; i >= lower; i-- {
		j := _palindrome(i)
		for k := upper; k*k > j; k-- {
			if j%k == 0 {
				return j % 1337
			}
		}
	}

	return 0
}
