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

package p0077

func combine(n int, k int) [][]int {
	_an := func(n int) int {
		r := 1
		for i := 1; i <= n; i++ {
			r *= i
		}
		return r
	}
	_cn := func(n int, k int) int {
		return _an(n) / (_an(k) * _an(n-k))
	}
	l := _cn(n, k)
	ret, w, arr := make([][]int, l), 0, make([]int, k)

	_walk := func(n int, last int, size int, fn interface{}) {
		if size == k {
			cp := make([]int, k)
			copy(cp, arr)
			ret[w] = cp
			w++
			return
		}

		_fn := fn.(func(int, int, int, interface{}))
		for i := last; i <= n; i++ {
			arr[size] = i
			_fn(n, i+1, size+1, fn)
		}
	}

	_walk(n, 1, 0, _walk)
	return ret
}
