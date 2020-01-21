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

package p0728

func selfDividingNumbers(left int, right int) []int {
	_isSelfDivid := func(n int) bool {
		r := n
		for n != 0 {
			m := n % 10
			if m == 0 || r%m != 0 {
				return false
			}

			n = n / 10
		}
		return true
	}

	ret := make([]int, 0)

	for i := left; i <= right; i++ {
		if _isSelfDivid(i) {
			ret = append(ret, i)
		}
	}

	return ret
}
