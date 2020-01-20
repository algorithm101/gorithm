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

package p0205

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	smaps := make(map[byte]byte, len(s))
	dmaps := make(map[byte]byte, len(s))

	for i := 0; i < len(s); i++ {
		sc := s[i]
		tc, exists := smaps[sc]
		if exists {
			if t[i] != tc {
				return false
			}
		} else {
			if dmc, ok := dmaps[t[i]]; ok && dmc != sc {
				return false
			}
			smaps[sc] = t[i]
			dmaps[t[i]] = sc
		}
	}

	return true
}
