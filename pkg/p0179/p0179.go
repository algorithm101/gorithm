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

package p0179

import (
	"sort"
	"strconv"
	"strings"
)

// Numbers ...
type Numbers []string

// Less ...
func (n Numbers) Less(i, j int) bool {
	v1 := n[i] + n[j]
	v2 := n[j] + n[i]
	return v1 < v2
}

// Len ...
func (n Numbers) Len() int {
	return len(n)
}

// Swap ...
func (n Numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func largestNumber(nums []int) string {
	numstrs := make([]string, len(nums))
	for i, num := range nums {
		numstrs[i] = strconv.Itoa(num)
	}

	sort.Sort(sort.Reverse(Numbers(numstrs)))

	r, i := strings.Join(numstrs, ""), 0
	for i < len(r) && r[i] == '0' {
		i++
	}
	if i >= len(r) {
		return "0"
	}
	return r[i:]
}
