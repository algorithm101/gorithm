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

package p0409

func longestPalindrome(s string) int {
	maps := make(map[rune]int, 100)
	for _, v := range s {
		count := maps[v]
		maps[v] = count + 1
	}

	single := 0
	length := 0
	for _, v := range maps {
		if v%2 == 1 {
			single = 1
			length += v - 1
		} else {
			length += v
		}
	}

	return length + single
}
