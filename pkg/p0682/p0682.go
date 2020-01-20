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

package p0682

import (
	"strconv"
)

func calPoints(ops []string) int {
	points, scores, w := 0, make([]int, len(ops)), 0

	for _, op := range ops {
		switch op {
		case "C":
			w--
			v := scores[w]
			points -= v
		case "D":
			v := scores[w-1]
			points += 2 * v
			scores[w] = 2 * v
			w++
		case "+":
			v := scores[w-1] + scores[w-2]
			points += v
			scores[w] = v
			w++
		default:
			v, _ := strconv.Atoi(op)
			points += v
			scores[w] = v
			w++
		}
	}

	return points
}
