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

package p0018

import "sort"

func threeSumTarget(nums []int, target0 int) [][]int {
	rs := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		target := target0 - nums[i]
		j := i + 1
		k := len(nums) - 1
		movej := 0
		for j < k {
			if movej == 1 && nums[j] == nums[j-1] {
				j++
				continue
			} else if movej == 2 && nums[k] == nums[k+1] {
				k--
				continue
			}

			sum := nums[j] + nums[k]
			if sum == target {
				rs = append(rs, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				movej = 1
			} else if sum < target {
				j++
				movej = 1
			} else {
				k--
				movej = 2
			}
		}
	}

	return rs
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	rs := make([][]int, 0)

	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		rs3 := threeSumTarget(nums[i+1:], target-nums[i])
		for _, row := range rs3 {
			rs = append(rs, []int{nums[i], row[0], row[1], row[2]})
		}
	}

	return rs
}
