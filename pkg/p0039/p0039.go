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

package p0039

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ret := make([][]int, 0)

	_walk := func(candidates []int, target int, arr []int, index int, fn interface{}) {
		if target == 0 {
			r := make([]int, len(arr))
			copy(r, arr)
			ret = append(ret, r)
			return
		}

		_fn := fn.(func([]int, int, []int, int, interface{}))
		for i := index; i < len(candidates); i++ {
			if candidates[i] <= target {
				arr = append(arr, candidates[i])
				_fn(candidates, target-candidates[i], arr, i, fn)
				arr = arr[0 : len(arr)-1]
			}
		}
	}

	_walk(candidates, target, []int{}, 0, _walk)

	return ret
}
