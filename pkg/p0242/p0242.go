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

package p0242

func isAnagram(s string, t string) bool {
	counts := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counts[byte(s[i])-'a']++
	}

	for i := 0; i < len(t); i++ {
		c := counts[byte(t[i]-'a')]
		if c < 1 {
			return false
		}
		counts[byte(t[i])-'a']--
	}

	for _, c := range counts {
		if c != 0 {
			return false
		}
	}

	return true
}
