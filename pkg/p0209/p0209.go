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

package p0209

import (
	"math"
)

// TODO: 实现 O(n*lg(n)) 的解法
func minSubArrayLen(s int, nums []int) int {
	sum, min, last := 0, math.MaxInt32, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum >= s {
			// forward last
			for sum >= s {
				sum -= nums[last]
				last++
			}

			if i-last+2 < min {
				min = i - last + 2
			}
		}
	}

	if min == math.MaxInt32 {
		return 0
	}
	return min
}
