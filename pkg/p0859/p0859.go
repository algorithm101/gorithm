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

package p0859

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	diffA := make([]int, 0)

	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			diffA = append(diffA, i)
		}
	}

	if len(diffA) == 0 {
		countIndex := make([]int, 26)
		for _, v := range A {
			countIndex[byte(v)-'a']++
			if countIndex[byte(v)-'a'] >= 2 {
				return true
			}
		}
	} else if len(diffA) == 2 {
		return A[diffA[0]] == B[diffA[1]] && A[diffA[1]] == B[diffA[0]]
	}

	return false
}
