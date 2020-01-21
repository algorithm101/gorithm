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

package p0153

// TODO: 尝试使用二分
func findMin(nums []int) int {
	if 0 == len(nums) {
		return 0
	}
	if 1 == len(nums) {
		return nums[0]
	}

	last := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < last {
			return nums[i]
		}
		last = nums[i]
	}

	return nums[0]
}
