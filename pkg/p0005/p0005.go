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

package p0005

import (
	"strings"
)

func minInt(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func longestPalindrome(s string) string {
	nsr := make([]rune, len(s)*2+1)
	for i := 0; i < len(nsr); i++ {
		nsr[i] = '#'
	}
	for i, c := range s {
		nsr[2*i+1] = c
	}
	ns := string(nsr)

	RL := make([]int, len(ns))
	MaxRight, pos, MaxLen := 0, 0, 0
	index := 0
	for i := 0; i < len(ns); i++ {
		if i < MaxRight {
			RL[i] = minInt(RL[2*pos-i], MaxRight-i)
		} else {
			RL[i] = 1
		}

		for i-RL[i] >= 0 && i+RL[i] < len(ns) && ns[i-RL[i]] == ns[i+RL[i]] {
			RL[i]++
		}
		if RL[i]+i-1 > MaxRight {
			MaxRight = RL[i] + i - 1
			pos = i
		}

		if MaxLen < RL[i] {
			MaxLen = RL[i]
			index = i
		}
	}
	return strings.Replace(ns[index-MaxLen+1:index+MaxLen], "#", "", -1)
}
