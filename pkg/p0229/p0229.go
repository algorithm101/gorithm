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

package p0229

// TODO: 用求超过一半的数的相同解法处理
func majorityElement(nums []int) []int {
	quickSelect := func(nums []int, k int) int {
	START:
		if 0 == len(nums) {
			panic("quickSelect: nums length should not be zero")
		}

		pivot, i, j := nums[0], 0, len(nums)-1
		for i != j {
			for i < j && nums[j] >= pivot {
				j--
			}
			for i < j && nums[i] <= pivot {
				i++
			}

			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		nums[0], nums[i] = nums[i], pivot

		switch {
		case i+1 == k:
			return pivot
		case i+1 < k:
			nums = nums[i+1:]
			k = k - i - 1
			goto START
		default:
			nums = nums[0:i]
			goto START
		}
	}

	_helper := func(nums []int, k int) (int, bool) {
		if k > len(nums) {
			return 0, false
		}

		v, count := quickSelect(nums, k), 0
		for _, num := range nums {
			if num == v {
				count++
			}
		}

		if count > len(nums)/3 {
			return v, true
		}
		return 0, false
	}

	r := []int{}
	if v, ok := _helper(nums, len(nums)/3+1); ok {
		r = append(r, v)
	}
	if v, ok := _helper(nums, len(nums)*2/3+1); ok {
		if len(r) == 0 || r[0] != v {
			r = append(r, v)
		}
	}

	return r
}
