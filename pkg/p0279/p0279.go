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

package p0279

import (
	"math"
)

// TODO: 广度优先搜索, Lagrange's four-square theorem 解法
func numSquares(n int) int {
	r := make([]int, n+1)
	r[1] = 1

	_sqrt := func(n int) int {
		return int(math.Sqrt(float64(n)))
	}

	_isSquare := func(n int) bool {
		v := _sqrt(n)
		return n == v*v
	}
	_min := func(x int, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i := 2; i <= n; i++ {
		if _isSquare(i) {
			r[i] = 1
		} else {
			r[i] = math.MaxInt32
			for j := 1; j <= i/2; j++ {
				r[i] = _min(r[j]+r[i-j], r[i])
			}
		}

	}

	return r[len(r)-1]
}
