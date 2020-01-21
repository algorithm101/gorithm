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

package p0216

func combinationSum3(k int, n int) [][]int {
	_sum := func(arr []int) int {
		sum := 0
		for _, v := range arr {
			sum += v
		}
		return sum
	}

	ret, arr := make([][]int, 0), make([]int, 0)
	_walk := func(start int, fn interface{}) {
		if len(arr) == k && _sum(arr) == n {
			cp := make([]int, len(arr))
			copy(cp, arr)
			ret = append(ret, cp)
			return
		}

		sum, l := _sum(arr), len(arr)
		_fn := fn.(func(int, interface{}))

		for i := start; i <= 9; i++ {
			if sum+i > n {
				break
			}

			arr = append(arr, i)
			_fn(i+1, fn)
			arr = arr[0:l]
		}
	}
	_walk(1, _walk)

	return ret
}
