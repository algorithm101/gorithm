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

package p0187

// TODO: 优化 map 中的内存使用
func findRepeatedDnaSequences(s string) []string {
	maps := make(map[string]int, len(s)-9)
	for i := 9; i < len(s); i++ {
		v := s[i-9 : i+1]
		maps[v] = maps[v] + 1
	}

	ret, w := make([]string, len(maps)), 0
	for k, v := range maps {
		if v >= 2 {
			ret[w] = k
			w++
		}
	}
	return ret[0:w]
}
