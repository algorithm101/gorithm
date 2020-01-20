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

package p0241

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(input string) []int {
	if 0 == len(input) {
		return []int{}
	}

	_toInt := func(v string) int {
		r, _ := strconv.Atoi(v)
		return r
	}
	apply := func(operator string, op1 int, op2 int) int {
		switch operator {
		case "+":
			return op1 + op2
		case "-":
			return op1 - op2
		case "*":
			return op1 * op2
		}

		panic(fmt.Errorf("Unsupported operator: %s", operator))
	}

	r := make([]int, 0)

	for i := 0; i < len(input); i++ {
		if !(input[i] >= '0' && input[i] <= '9') {
			left := diffWaysToCompute(input[0:i])
			right := diffWaysToCompute(input[i+1:])

			for _, op1 := range left {
				for _, op2 := range right {
					r = append(r, apply(string(input[i]), op1, op2))
				}
			}
		}
	}

	if 0 == len(r) {
		r = append(r, _toInt(input))
	}
	return r
}
