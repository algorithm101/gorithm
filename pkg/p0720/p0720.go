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

package p0720

import (
	"sort"
)

func longestWord(words []string) string {
	maps := make(map[string]bool, len(words))
	for _, w := range words {
		maps[w] = true
	}

	sort.Strings(words)

	ret := ""
	for i := len(words) - 1; i >= 0; i-- {
		if len(words[i]) < len(ret) {
			continue
		}

		w, exists := words[i], true
		for j := 1; j < len(w); j++ {
			if maps[w[0:j]] == false {
				exists = false
				break
			}
		}
		if exists {
			ret = w
		}
	}
	return ret
}
