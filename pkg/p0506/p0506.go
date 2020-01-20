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

package p0506

import (
	"fmt"
	"sort"
)

func findRelativeRanks(nums []int) []string {
	sorts := make([]int, len(nums))
	copy(sorts, nums)
	sort.Sort(sort.Reverse(sort.IntSlice(sorts)))

	maps := make(map[int]int, len(sorts))
	for i, v := range sorts {
		maps[v] = i + 1
	}

	_toString := func(v int) string {
		switch v {
		case 1:
			return "Gold Medal"
		case 2:
			return "Silver Medal"
		case 3:
			return "Bronze Medal"
		default:
			return fmt.Sprintf("%d", v)
		}
	}

	ret := make([]string, len(nums))
	for i, v := range nums {
		rank := maps[v]
		ret[i] = _toString(rank)
	}

	return ret
}
