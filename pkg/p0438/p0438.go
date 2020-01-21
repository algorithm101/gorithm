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

package p0438

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	mapsp := make([]int, 26)
	for _, v := range p {
		mapsp[byte(v)-'a']++
	}

	_isEqual := func(arg1 []int, arg2 []int) bool {
		for i := 0; i < 26; i++ {
			if arg1[i] != arg2[i] {
				return false
			}
		}

		return true
	}

	mapss := make([]int, 26)
	for i := 0; i < len(p); i++ {
		mapss[s[i]-'a']++
	}

	ret := make([]int, 0)
	if _isEqual(mapss, mapsp) {
		ret = append(ret, 0)
	}

	for i := len(p); i < len(s); i++ {
		mapss[s[i-len(p)]-'a']--

		mapss[s[i]-'a']++
		if _isEqual(mapss, mapsp) {
			ret = append(ret, i-len(p)+1)
		}
	}

	return ret
}
