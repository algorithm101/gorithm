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

package p0747

func dominantIndex(nums []int) int {
	if 0 == len(nums) {
		return -1
	}
	if 1 == len(nums) {
		return 0
	}

	max, max1, index := nums[0], nums[1], 0
	if nums[0] < nums[1] {
		max, max1, index = nums[1], nums[0], 1
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > max {
			max, max1, index = nums[i], max, i
		} else if nums[i] > max1 {
			max1 = nums[i]
		}
	}

	if max >= 2*max1 {
		return index
	}

	return -1
}
