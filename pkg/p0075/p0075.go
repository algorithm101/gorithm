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

package p0075

func sortColors(nums []int) {
	if 0 == len(nums) {
		return
	}

	start, cursor, end := 0, 0, len(nums)-1
	for cursor <= end {
		switch {
		case nums[cursor] == 0 && cursor == start:
			start++
			cursor++
		case nums[cursor] == 0 && cursor != start:
			nums[start], nums[cursor] = nums[cursor], nums[start]
			start++
		case nums[cursor] == 2:
			nums[cursor], nums[end] = nums[end], nums[cursor]
			end--
		default:
			cursor++
		}
	}
}
