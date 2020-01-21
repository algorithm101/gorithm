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

package p0628

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)

	max1 := nums[0] * nums[1] * nums[len(nums)-1]
	max2 := nums[len(nums)-3] * nums[len(nums)-2] * nums[len(nums)-1]

	if max1 > max2 {
		return max1
	}

	return max2
}
