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

package p0169

func majorityElement(nums []int) int {
	ret, count := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if 0 == count {
			ret = nums[i]
			count = 1
		} else {
			if ret == nums[i] {
				count++
			} else {
				count--
			}
		}
	}

	return ret
}
