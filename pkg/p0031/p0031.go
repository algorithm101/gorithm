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

package p0031

import (
	"sort"
)

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}

	if i < 0 {
		sort.Ints(nums)
		return
	}

	for j := len(nums) - 1; j > i; j-- {
		if nums[j] > nums[i] {
			nums[j], nums[i] = nums[i], nums[j]
			sort.Ints(nums[i+1:])
			return
		}
	}
}
