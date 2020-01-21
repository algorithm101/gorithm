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

package p0049

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	_sort := func(v string) string {
		arr := strings.Split(v, "")
		sort.Strings(arr)
		return strings.Join(arr, "")
	}

	maps := map[string][]string{}
	for _, str := range strs {
		v := _sort(str)
		maps[v] = append(maps[v], str)
	}

	keys, w := make([]string, len(maps)), 0
	for k := range maps {
		keys[w] = k
		w++
	}
	sort.Strings(keys)

	ret, w := make([][]string, len(maps)), 0
	for _, k := range keys {
		v := maps[k]
		sort.Strings(v)
		ret[w] = v
		w++
	}

	return ret
}
