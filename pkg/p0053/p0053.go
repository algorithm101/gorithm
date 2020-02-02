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

package p0053

import (
	"math"
)

func maxSubArray(nums []int) int {
	_max := func(v1 int, v2 int) int {
		if v1 > v2 {
			return v1
		}
		return v2
	}

	_f1 := func(nums []int) int {
		max, curr := nums[0], 0
		for i := 0; i < len(nums); i++ {
			curr += nums[i]
			if curr > max {
				max = curr
			}

			if curr < 0 {
				curr = 0
			}
		}
		return max
	}
	_ = _f1

	_f2 := func(nums []int) int {
		if 1 == len(nums) {
			return nums[0]
		} else if 0 == len(nums) {
			return math.MinInt32
		}

		mid := len(nums) / 2
		maxLeft := maxSubArray(nums[:mid])
		maxRight := maxSubArray(nums[mid:])

		maxLeftMid := 0 - 0xFFFFFFFF
		currMaxLeftMid := 0
		for i := mid - 1; i >= 0; i-- {
			currMaxLeftMid += nums[i]
			if currMaxLeftMid > maxLeftMid {
				maxLeftMid = currMaxLeftMid
			}
		}

		maxRightMid := 0 - 0xFFFFFFFF
		currMaxRightMid := 0
		for i := mid; i < len(nums); i++ {
			currMaxRightMid += nums[i]
			if currMaxRightMid > maxRightMid {
				maxRightMid = currMaxRightMid
			}
		}

		maxMidSum := maxLeftMid + maxRightMid

		return _max(_max(maxLeft, maxRight), maxMidSum)
	}
	_ = _f2

	_f3 := func(nums []int) int {
		max, last := nums[0], 0

		for i := 0; i < len(nums); i++ {
			m := _max(nums[i], last+nums[i])
			if m > max {
				max = m
			}
			last = m
		}

		return max
	}
	_ = _f3

	return _f3(nums)
}
