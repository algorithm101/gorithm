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

package p0318

import (
	"sort"
)

type L []string

func (l L) Swap(i int, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l L) Len() int {
	return len(l)
}

func (l L) Less(i int, j int) bool {
	return len(l[i]) < len(l[j])
}

func maxProduct(words []string) int {
	_isDiff := func(s1 string, s2 string) bool {
		r1 := make([]int, 26)
		for i := 0; i < len(s1); i++ {
			r1[int(s1[i]-'a')]++
		}
		for i := 0; i < len(s2); i++ {
			if r1[int(s2[i]-'a')] > 0 {
				return false
			}
		}

		return true
	}

	sort.Sort(sort.Reverse(L(words)))

	max := 0

	for i := 0; i < len(words); i++ {
		if len(words[i])*len(words[i]) < max {
			break
		}

		for j := i + 1; j < len(words); j++ {
			if _isDiff(words[i], words[j]) {
				if len(words[i])*len(words[j]) > max {
					max = len(words[i]) * len(words[j])
					break
				}
			}
		}
	}

	return max
}
