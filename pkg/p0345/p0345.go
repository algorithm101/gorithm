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

package p0345

func reverseVowels(s string) string {
	isVowel := func(b byte) bool {
		vowels := "aeiouAEIOU"
		for i := 0; i < len(vowels); i++ {
			if b == vowels[i] {
				return true
			}
		}

		return false
	}

	sbyte := []byte(s)
	for i, j := 0, len(sbyte)-1; i < j; i, j = i+1, j-1 {
		for !isVowel(sbyte[i]) && i < j {
			i++
		}
		for !isVowel(sbyte[j]) && i < j {
			j--
		}

		if i >= j {
			break
		}

		sbyte[i], sbyte[j] = sbyte[j], sbyte[i]
	}
	return string(sbyte)
}
