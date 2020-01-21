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

package p0696

import (
	"math"
)

func countBinarySubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	count, prevCount, curr, currCount := 0, 0, s[0], 1
	for i := 1; i < len(s); i++ {
		switch {
		case s[i] == curr:
			currCount++
		default:
			count += int(math.Min(float64(prevCount), float64(currCount)))
			curr = s[i]
			prevCount = currCount
			currCount = 1
		}
	}

	count += int(math.Min(float64(prevCount), float64(currCount)))
	return count
}
