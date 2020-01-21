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

package p0383

func canConstruct(ransomNote string, magazine string) bool {
	maps := make(map[byte]int, len(ransomNote))
	for i := 0; i < len(ransomNote); i++ {
		count := maps[ransomNote[i]]
		maps[ransomNote[i]] = count + 1
	}

	chars := len(ransomNote)
	for i := 0; i < len(magazine); i++ {
		count := maps[magazine[i]]
		if count > 0 {
			maps[magazine[i]] = count - 1
			chars--

			if chars == 0 {
				return true
			}
		}
	}

	return chars == 0
}
