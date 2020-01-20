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

package p0080

func removeDuplicates(nums []int) int {
	if 0 == len(nums) {
		return 0
	}

	r, w := 1, 0
	for r < len(nums) {
		switch {
		case nums[r] != nums[w]:
			fallthrough
		case (w > 0 && nums[w-1] != nums[w]) || (w == 0):
			w++
			nums[w] = nums[r]
		}
		r++
	}

	return w + 1
}
