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

package p0047

import "sort"

func permuteUnique(nums []int) [][]int {
	if 0 == len(nums) {
		return make([][]int, 0)
	}

	_an := func(n int) int {
		r := 1
		for n > 0 {
			r *= n
			n--
		}
		return r
	}
	ret, w := make([][]int, _an(len(nums))), 0

	sort.Ints(nums)
	_walk := func(pos int, fn interface{}) {
		if pos == len(nums) {
			r := make([]int, len(nums))
			copy(r, nums)
			ret[w] = r
			w++
		}

		_fn := fn.(func(int, interface{}))
		visited := make(map[int]int)
		for i := pos; i < len(nums); i++ {
			if _, ok := visited[nums[i]]; !ok {
				nums[pos], nums[i] = nums[i], nums[pos]
				_fn(pos+1, fn)
				nums[i], nums[pos] = nums[pos], nums[i]

				visited[nums[i]] = 1
			}
		}
	}
	_walk(0, _walk)

	ret = ret[0:w]
	return ret
}
