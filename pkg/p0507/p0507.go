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

package p0507

import (
	"math"
)

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	r, min := 1, int(math.Sqrt(float64(num)))
	for i := 2; i <= min; i++ {
		if num%i == 0 {
			r += i
			r += num / i
			// min = num / i
		}

		if r > num {
			return false
		}
	}

	return r == num
}
