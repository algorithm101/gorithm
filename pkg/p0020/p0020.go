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

package p0020

import "container/list"

func isValid(s string) bool {
	stack := list.New()

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack.PushBack(c)
		} else {
			e := stack.Back()
			if e == nil {
				return false
			}

			lastE := e.Value.(rune)
			if (c == ')' && lastE != '(') ||
				(c == ']' && lastE != '[') ||
				(c == '}' && lastE != '{') {
				return false
			}

			stack.Remove(e)
		}
	}

	return stack.Len() == 0
}
