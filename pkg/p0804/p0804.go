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

package p0804

func uniqueMorseRepresentations(words []string) int {
	mos := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	counts := make(map[string]int)

	target, w := make([]byte, 48), 0
	for _, word := range words {
		for i := 0; i < len(target); i++ {
			target[i] = ' '
		}
		w = 0

		for i := 0; i < len(word); i++ {
			mosv := mos[word[i]-'a']
			for j := 0; j < len(mosv); j++ {
				target[w] = mosv[j]
				w++
			}
		}

		counts[string(target)] = 1
	}

	return len(counts)
}
