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

package p0022

import (
	"math"
)

func generateParenthesis(n int) []string {
	ret, buf, w := make([]string, int(math.Pow(2.0, float64(2*n)))), make([]byte, 2*n), 0

	_walk := func(left int, right int, fn interface{}) {
		if 0 == left && 0 == right {
			ret[w] = string(buf)
			w++
		}

		_fn := fn.(func(int, int, interface{}))
		if left > 0 {
			buf[2*n-left-right] = '('
			_fn(left-1, right, fn)
		}
		if left < right {
			buf[2*n-left-right] = ')'
			_fn(left, right-1, fn)
		}
	}
	_walk(n, n, _walk)

	ret = ret[0:w]
	return ret
}
