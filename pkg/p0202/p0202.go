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

package p0202

func isHappy(n int) bool {
	_next := func(n int) int {
		ret := 0

		for n != 0 {
			i := n % 10
			ret += i * i
			n /= 10
		}

		return ret
	}

	exists := make(map[int]bool, 10)
	exists[n] = true
	for n != 1 {
		n = _next(n)
		_, exist := exists[n]
		if exist {
			return false
		}
		exists[n] = true
	}

	return true
}
