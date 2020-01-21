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

package p0227

import (
	"bytes"
	"container/list"
	"strconv"
)

// TODO: 实现为先计算乘除, 后计算加减
func calculate(s string) int {
	_strToInt := func(v string) int {
		r, _ := strconv.Atoi(v)
		return r
	}

	_splitToField := func(s string) []string {
		r := make([]string, 0)

		buf := bytes.Buffer{}
		for i := 0; i < len(s); i++ {
			switch {
			case s[i] == ' ':
				if buf.Len() > 0 {
					r = append(r, buf.String())
					buf.Reset()
				}
			case s[i] == '+':
				fallthrough
			case s[i] == '-':
				fallthrough
			case s[i] == '*':
				fallthrough
			case s[i] == '/':
				if buf.Len() > 0 {
					r = append(r, buf.String())
					buf.Reset()
				}
				r = append(r, string(s[i]))
			default:
				buf.WriteByte(s[i])
			}
		}

		if buf.Len() > 0 {
			r = append(r, buf.String())
		}

		return r
	}

	fields := _splitToField(s)
	_toPostfixExp := func(fields []string) []string {
		r := make([]string, 0)
		s := list.New()

		for _, field := range fields {
			switch {
			case field == "*" || field == "/":
				for s.Len() > 0 {
					e := s.Back()

					if e.Value.(string) == "+" || e.Value.(string) == "-" {
						break
					}

					s.Remove(e)

					r = append(r, e.Value.(string))
				}
				s.PushBack(field)
			case field == "+" || field == "-":
				for s.Len() > 0 {
					e := s.Back()
					s.Remove(e)

					r = append(r, e.Value.(string))
				}
				s.PushBack(field)
			default:
				r = append(r, field)
			}
		}

		for s.Len() > 0 {
			e := s.Back()
			s.Remove(e)

			r = append(r, e.Value.(string))
		}
		return r
	}
	fields = _toPostfixExp(fields)

	operandStack := list.New()
	for _, field := range fields {
		switch {
		case field == "+":
			e1 := operandStack.Back()
			operandStack.Remove(e1)
			e2 := operandStack.Back()
			operandStack.Remove(e2)

			operandStack.PushBack(e1.Value.(int) + e2.Value.(int))
		case field == "-":
			e1 := operandStack.Back()
			operandStack.Remove(e1)
			e2 := operandStack.Back()
			operandStack.Remove(e2)

			operandStack.PushBack(e2.Value.(int) - e1.Value.(int))
		case field == "*":
			e1 := operandStack.Back()
			operandStack.Remove(e1)
			e2 := operandStack.Back()
			operandStack.Remove(e2)

			operandStack.PushBack(e1.Value.(int) * e2.Value.(int))
		case field == "/":
			e1 := operandStack.Back()
			operandStack.Remove(e1)
			e2 := operandStack.Back()
			operandStack.Remove(e2)

			operandStack.PushBack(e2.Value.(int) / e1.Value.(int))
		default:
			operandStack.PushBack(_strToInt(field))
		}
	}
	return operandStack.Front().Value.(int)
}
