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

package p0299

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	countA := func(secret string, guess string) int {
		count := 0
		for i := 0; i < len(secret); i++ {
			if secret[i] == guess[i] {
				count++
			}
		}
		return count
	}
	countB := func(secret string, guess string) int {
		count := 0

		m1, m2 := make(map[byte]int), make(map[int]bool)
		for i := 0; i < len(secret); i++ {
			if secret[i] == guess[i] {
				m2[i] = true
			} else {
				m1[secret[i]] = m1[secret[i]] + 1
			}
		}

		for i := 0; i < len(guess); i++ {
			if !m2[i] && m1[guess[i]] > 0 {
				m1[guess[i]] = m1[guess[i]] - 1
				count++
			}
		}

		return count
	}
	return fmt.Sprintf("%dA%dB", countA(secret, guess), countB(secret, guess))
}
