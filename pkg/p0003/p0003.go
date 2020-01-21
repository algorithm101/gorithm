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

package p0003

// 动态规划
func lengthOfLongestSubstring(s string) int {
	length := 0
	startIndex := 0
	cachedIndex := make(map[rune]int)

	for i, c := range s {
		lastIndex, exists := cachedIndex[c]
		if exists && lastIndex >= startIndex {
			startIndex = lastIndex + 1
		}

		currLength := i - startIndex + 1
		if currLength > length {
			length = currLength
		}
		cachedIndex[c] = i
	}

	return length
}
