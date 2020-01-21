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

package p0150

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	_toInt := func(s string) int {
		v, _ := strconv.Atoi(s)
		return v
	}
	_toStr := func(v int) string {
		return strconv.Itoa(v)
	}
	_isOperator := func(s string) bool {
		operators := []string{"+", "-", "*", "/"}
		for i := 0; i < len(operators); i++ {
			if s == operators[i] {
				return true
			}
		}

		return false
	}
	_op := func(num1 string, num2 string, operator string) int {
		switch operator {
		case "+":
			return _toInt(num1) + _toInt(num2)
		case "-":
			return _toInt(num1) - _toInt(num2)
		case "*":
			return _toInt(num1) * _toInt(num2)
		case "/":
			return _toInt(num1) / _toInt(num2)
		}

		panic("Invalid Operator")
	}

	w := 0
	for i := 0; i < len(tokens); i++ {
		if _isOperator(tokens[i]) {
			t := _op(tokens[w-2], tokens[w-1], tokens[i])
			w -= 2
			tokens[w] = _toStr(t)
			w++
		} else {
			tokens[w] = tokens[i]
			w++
		}
	}

	return _toInt(tokens[0])
}
