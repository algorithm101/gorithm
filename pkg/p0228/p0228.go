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

package p0228

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	if 0 == len(nums) {
		return []string{}
	}

	_toString := func(first int, last int) string {
		if first == last {
			return fmt.Sprintf("%d", first)
		}
		return fmt.Sprintf("%d->%d", first, last)
	}

	r := make([]string, 0)
	first, last := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == last+1 {
			last = nums[i]
		} else {
			r = append(r, _toString(first, last))
			first, last = nums[i], nums[i]
		}
	}

	r = append(r, _toString(first, last))
	return r
}
