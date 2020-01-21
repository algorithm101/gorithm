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

package p0089

func grayCode(n int) []int {
	_f1 := func(n int) []int {
		ret, w := make([]int, 1<<uint(n)), 0
		ret[w] = 0
		w++

		for i := 0; i < n; i++ {
			for j := w - 1; j >= 0; j-- {
				r := 1<<uint(i) + ret[j]
				ret[w] = r
				w++
			}
		}

		return ret
	}
	_ = _f1

	_f2 := func(n int) []int {
		ret, w := make([]int, 1<<uint(n)), 0
		for i := 0; i < (1 << uint(n)); i++ {
			ret[w] = i ^ (i >> 1)
			w++
		}

		return ret
	}
	_ = _f2

	return _f2(n)
}
