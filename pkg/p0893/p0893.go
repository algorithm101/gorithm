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

package p0893

import (
	"sort"
)

func numSpecialEquivGroups(A []string) int {
	maps := make(map[string]int, 0)

	_trans := func(s string) string {
		ret, w := make([]byte, len(s)+1), 0

		arro, wo := make([]int, len(s)/2+len(s)%2), 0
		for i := 0; i < len(s); i += 2 {
			arro[wo] = int(s[i])
			wo++
		}
		sort.Ints(arro)

		for i := 0; i < len(arro); i++ {
			ret[w] = byte(arro[i])
			w++
		}
		ret[w] = '#'
		w++

		arrl, wl := make([]int, len(s)/2), 0
		for i := 1; i < len(s); i += 2 {
			arrl[wl] = int(s[i])
			wl++
		}
		sort.Ints(arrl)

		for i := 0; i < len(arrl); i++ {
			ret[w] = byte(arrl[i])
			w++
		}

		return string(ret)
	}

	for _, s := range A {
		maps[_trans(s)] = 1
	}

	return len(maps)
}
