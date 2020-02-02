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

package p0819

func mostCommonWord(paragraph string, banned []string) string {
	bannedMaps := make(map[string]bool, len(banned))
	for _, v := range banned {
		bannedMaps[v] = true
	}

	_isAlpha := func(b byte) bool {
		return 'a' <= b && b <= 'z'
	}
	last := -1
	usedMaps := make(map[string]int)

	paragraphBytes := make([]byte, len(paragraph)+1)
	copy(paragraphBytes, []byte(paragraph))
	paragraphBytes[len(paragraph)] = '.'

	for i, b := range paragraphBytes {
		if b >= 'A' && b <= 'Z' {
			b = b - 'A' + 'a'
			paragraphBytes[i] = b
		}

		if _isAlpha(b) {
			if last == -1 {
				last = i
			}
		} else {
			if last != -1 {
				s := string(paragraphBytes[last:i])
				if !bannedMaps[s] {
					count := usedMaps[s]
					usedMaps[s] = count + 1
				}
				last = -1
			}
		}
	}

	maxCount, maxString := 0, ""
	for k, v := range usedMaps {
		if v > maxCount {
			maxCount = v
			maxString = k
		}
	}

	return maxString
}
