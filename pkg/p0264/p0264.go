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

package p0264

func nthUglyNumber(n int) int {
	index1, index2, index3 := 0, 0, 0
	nth, w := make([]int, n), 1
	nth[0] = 1

	_min := func(x int, y int, z int) int {
		r := x
		if r > y {
			r = y
		}
		if r > z {
			r = z
		}

		return r
	}

	for n > 1 {
		n--

		m := _min(nth[index1]*2, nth[index2]*3, nth[index3]*5)
		nth[w] = m
		w++

		if m == nth[index1]*2 {
			index1++
		}
		if m == nth[index2]*3 {
			index2++
		}
		if m == nth[index3]*5 {
			index3++
		}
	}

	return nth[len(nth)-1]
}
