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

package p0824

import (
	"strings"
)

func toGoatLatin(S string) string {
	words, l := strings.Split(S, " "), -1
	for i, word := range words {
		l += i + 1 + 2 + len(word) + 1
	}

	if l < 0 {
		panic("length error")
	}

	_isVowel := func(b byte) bool {
		vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
		for _, vowel := range vowels {
			if vowel == b {
				return true
			}
		}

		return false
	}

	ret, w := make([]byte, l), 0
	for i, word := range words {
		if i != 0 {
			ret[w] = ' '
			w++
		}

		if len(word) == 0 {
			continue
		}

		first := word[0]
		switch _isVowel(first) {
		case true:
			for j := 0; j < len(word); j++ {
				ret[w] = word[j]
				w++
			}
			ret[w] = 'm'
			ret[w+1] = 'a'
			w += 2
		default:
			for j := 1; j < len(word); j++ {
				ret[w] = word[j]
				w++
			}

			ret[w] = first
			ret[w+1] = 'm'
			ret[w+2] = 'a'
			w += 3
		}

		for j := 0; j < i+1; j++ {
			ret[w] = 'a'
			w++
		}
	}

	return string(ret)
}
