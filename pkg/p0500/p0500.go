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

package p0500

import (
	"strings"
)

func findWords(words []string) []string {
	rows := []string{
		"qwertyuiopQWERTYUIOP",
		"ASDFGHJKLasdfghjkl",
		"ZXCVBNMzxcvbnm",
	}
	_getRow := func(w rune) int {
		for i, row := range rows {
			if strings.Index(row, string(w)) != -1 {
				return i
			}
		}

		return -1
	}

	_isSameRow := func(s string) bool {
		if len(s) == 0 {
			return true
		}

		row := _getRow(rune(s[0]))
		for _, w := range s {
			if row != _getRow(w) {
				return false
			}
		}

		return true
	}

	r := make([]string, 0)
	for _, word := range words {
		if _isSameRow(word) {
			r = append(r, word)
		}
	}
	return r
}
