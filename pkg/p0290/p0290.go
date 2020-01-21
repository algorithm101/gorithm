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

package p0290

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	if len(pattern) != len(strs) {
		return false
	}

	maps := make(map[byte]string, len(pattern))
	rmaps := make(map[string]byte, len(pattern))

	for i := 0; i < len(pattern); i++ {
		v, exists := maps[byte(pattern[i])]
		if exists && v != strs[i] {
			return false
		}

		rc, exists := rmaps[strs[i]]
		if exists && rc != byte(pattern[i]) {
			return false
		}

		maps[byte(pattern[i])] = strs[i]
		rmaps[strs[i]] = byte(pattern[i])
	}

	return true
}
