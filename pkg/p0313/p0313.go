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

package p0313

import (
	"math"
)

func nthSuperUglyNumber(n int, primes []int) int {
	r, w, pointer := make([]int, n), 1, make([]int, len(primes))
	r[0] = 1
	for w != n {
		min, index := math.MaxInt32, 0
		for i := 0; i < len(pointer); i++ {
			v := primes[i] * r[pointer[i]]
			if v < min {
				min = v
				index = i
			} else if min == v {
				pointer[i]++
			}
		}
		pointer[index]++
		r[w] = min
		w++
	}

	return r[n-1]
}
