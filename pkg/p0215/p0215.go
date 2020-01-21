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

package p0215

import (
	"sort"
)

// TODO: quick select implement
func findKthLargest(nums []int, k int) int {
	_f1 := func(nums []int, k int) int {
		sort.Sort(sort.Reverse(sort.IntSlice(nums)))

		return nums[k-1]
	}
	_ = _f1

	_f2 := func(nums []int, k int, fn interface{}) int {
	START:
		v, i, j := nums[0], 0, len(nums)-1
		for i != j {
			for j > i && nums[j] <= v {
				j--
			}
			for i < j && nums[i] >= v {
				i++
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		nums[0], nums[i] = nums[i], v

		// _fn := fn.(func([]int, int, interface{}) int)
		switch {
		case i+1 == k:
			return nums[i]
		case i+1 < k:
			nums, k = nums[i+1:], k-i-1
			goto START
			// return _fn(nums[i+1:], k-i-1, fn)
		default:
			nums = nums[0:i]
			goto START
			// return _fn(nums[0:i], k, fn)
		}
	}
	_ = _f2

	return _f2(nums, k, _f2)
}
