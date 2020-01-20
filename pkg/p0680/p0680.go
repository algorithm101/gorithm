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

package p0680

func validPalindrome(s string) bool {
	_isPanlindrome := func(s string, start int, end int) bool {
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}

		return true
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return _isPanlindrome(s, i, j-1) || _isPanlindrome(s, i+1, j)
		}
	}

	return true
}
