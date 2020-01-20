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

package p0415

import (
	"math"
)

func addStrings(num1 string, num2 string) string {
	ret := make([]byte, int(math.Max(float64(len(num1)), float64(len(num2))))+1)
	carry := 0
	for i, j, k := len(num1)-1, len(num2)-1, len(ret)-1; i >= 0 || j >= 0; i, j, k = i-1, j-1, k-1 {
		v1 := 0
		if i >= 0 {
			v1 = int(num1[i] - '0')
		}

		v2 := 0
		if j >= 0 {
			v2 = int(num2[j] - '0')
		}

		v := v1 + v2 + carry
		if v > 9 {
			carry = 1
			v %= 10
		} else {
			carry = 0
		}

		ret[k] = byte(v) + '0'
	}

	if carry == 1 {
		ret[0] = '1'
		return string(ret)
	}
	return string(ret[1:])
}
