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

package p0697

import (
	"math"
)

func findShortestSubArray(nums []int) int {
	mapcount, mapindex, min, du := make(map[int]int), make(map[int]int), math.MaxInt32, 0

	if len(nums) == 0 {
		return 0
	}

	min, du = 1, 1
	mapcount[nums[0]] = 1
	mapindex[nums[0]] = 0

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		count := mapcount[num]
		if count == 0 {
			mapcount[num] = 1
			mapindex[num] = i
		} else {
			count++
			mapcount[num] = count
			if count > du {
				du = count
				min = i - mapindex[num] + 1
			} else if count == du {
				if i-mapindex[num]+1 < min {
					min = i - mapindex[num] + 1
				}
			}
		}
	}

	return min
}
