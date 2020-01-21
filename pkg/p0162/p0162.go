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

package p0162

// TODO: 应该有更优美的实现方式
func findPeakElement(nums []int) int {
	if 1 == len(nums) {
		return 0
	}

	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		switch {
		case mid == len(nums)-1 && nums[mid] > nums[mid-1]:
			fallthrough
		case mid == 0 && nums[mid] > nums[mid+1]:
			fallthrough
		case nums[mid] > nums[mid+1] && nums[mid] > nums[mid-1]:
			return mid
		}

		if mid == len(nums)-1 || nums[mid] < nums[mid+1] {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}
