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

package p0001

func twoSum(nums []int, target int) []int {
	marks := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		numTarget := target - num
		index, ok := marks[numTarget]
		if ok {
			return []int{i, index}
		}

		marks[num] = i
	}

	return nil
}
