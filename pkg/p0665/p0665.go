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

package p0665

func checkPossibility(nums []int) bool {
	i, j := 0, len(nums)-1
	for i < j && nums[i] <= nums[i+1] {
		i++
	}

	for j > i && nums[j-1] <= nums[j] {
		j--
	}

	if j-i > 1 {
		return false
	}

	if i == 0 || j == len(nums)-1 {
		return true
	}

	if nums[i-1] <= nums[j] || nums[i] <= nums[j+1] {
		return true
	}

	return false
}
