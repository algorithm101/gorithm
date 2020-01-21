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

package p0762

func countPrimeSetBits(L int, R int) int {
	_countSetBits := func(n int) int {
		count := 0
		for n != 0 {
			count++
			n = n & (n - 1)
		}
		return count
	}
	_isPrime := func(n int) bool {
		primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
		for _, v := range primes {
			if n == v {
				return true
			}
		}

		return false
	}

	count := 0
	for i := L; i <= R; i++ {
		if _isPrime(_countSetBits(i)) {
			count++
		}
	}

	return count
}
