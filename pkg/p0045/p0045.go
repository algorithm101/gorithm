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

package p0045

func jump(nums []int) int {
	if 0 == len(nums) {
		return 0
	}

	count, now, max, index := 0, 0, -1, -1
	for now < len(nums)-1 {
		if now+nums[now] >= len(nums)-1 {
			count++
			break
		}

		for i := now + 1; i <= nums[now]+now && i < len(nums); i++ {
			if nums[i]+i > max {
				max = nums[i] + i
				index = i
			}
		}
		now, max = index, -1
		count++
	}
	return count
}
