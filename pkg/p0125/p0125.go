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

package p0125

import (
	"strings"
)

func isPalindrome(s string) bool {
	isAlphanumeric := func(c byte) bool {
		return (c >= byte('a') && c <= byte('z')) ||
			(c >= byte('A') && c <= byte('Z')) ||
			(c >= byte('0') && c <= byte('9'))
	}

	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlphanumeric(s[i]) {
			i++
		}
		for i < j && !isAlphanumeric(s[j]) {
			j--
		}

		if s[i] != s[j] {
			return false
		}
		i, j = i+1, j-1
	}

	return true
}
