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

package p0078

func subsets(nums []int) [][]int {
	ret, w := make([][]int, 1<<uint(len(nums))), 0
	ret[w] = []int{}
	w++

	for _, num := range nums {
		l := w
		for i := 0; i < l; i++ {
			arr := ret[i]
			cp := make([]int, len(arr), len(arr)+1)
			copy(cp, arr)
			cp = append(cp, num)

			ret[w] = cp
			w++
		}
	}

	return ret
}
