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

package p0034

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}

	first, last := 0, len(nums)-1
	for first <= last {
		mid := (last-first)/2 + first
		if nums[mid] > target {
			last = mid - 1
			continue
		} else if nums[mid] < target {
			first = mid + 1
			continue
		}

		if mid == 0 || nums[mid-1] != target {
			ret[0] = mid
			break
		}
		last = mid - 1
	}

	if -1 == ret[0] {
		return ret
	}

	first, last = ret[0], len(nums)-1
	for first <= last {
		mid := (last-first)/2 + first
		if nums[mid] > target {
			last = mid - 1
			continue
		} else if nums[mid] < target {
			first = mid + 1
			continue
		}

		if mid == len(nums)-1 || nums[mid+1] != target {
			ret[1] = mid
			break
		}
		first = mid + 1
	}

	return ret
}
