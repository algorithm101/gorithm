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

package p0090

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	ret := make([][]int, 0) // TODO: 预分配容量
	ret = append(ret, []int{})

	// 排序
	sort.Ints(nums)
	lastLen := 0

	for i, num := range nums {
		j := 0
		if i > 0 && nums[i-1] == num { // 重复的数
			j = lastLen
		}

		arrs, w := make([][]int, len(ret)-j), 0
		for ; j < len(ret); j++ {
			cp := make([]int, len(ret[j]), len(ret[j])+1)
			copy(cp, ret[j])
			cp = append(cp, num)

			arrs[w] = cp
			w++
		}

		lastLen = len(ret)
		ret = append(ret, arrs...)
	}

	return ret
}
