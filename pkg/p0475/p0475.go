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

package p0475

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	max, j := 0.0, 0
	for _, house := range houses {
		for j < len(heaters)-1 && heaters[j+1] < house {
			j++
		}

		min1, min2 := math.Abs(float64(house-heaters[j])), math.Abs(float64(house-heaters[j]))
		if j+1 < len(heaters) {
			min2 = math.Abs(float64(heaters[j+1] - house))
		}
		max = math.Max(max, math.Min(min1, min2))
	}
	return int(max)
}
