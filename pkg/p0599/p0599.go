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

package p0599

import (
	"math"
)

func findRestaurant(list1 []string, list2 []string) []string {
	ret, minIndex := make([]string, 0), math.MaxInt32

	maps1 := make(map[string]int, len(list1))
	for i, v := range list1 {
		maps1[v] = i
	}

	for i, v := range list2 {
		if ni, exists := maps1[v]; exists {
			index := ni + i
			if index < minIndex {
				ret, minIndex = []string{v}, index
			} else if index == minIndex {
				ret = append(ret, v)
			}
		}
	}

	return ret
}
