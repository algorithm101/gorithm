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

package p0496

func nextGreaterElement(findNums []int, nums []int) []int {
	nexts, w := make([]int, len(nums)), 0
	maps := make(map[int]int, len(nums))
	for _, v := range nums {
		for w > 0 && nexts[w-1] < v {
			maps[nexts[w-1]] = v
			w--
		}
		nexts[w] = v
		w++

	}

	r := make([]int, len(findNums))
	for i, k := range findNums {
		if v, exists := maps[k]; exists {
			r[i] = v
		} else {
			r[i] = -1
		}
	}
	return r
}
