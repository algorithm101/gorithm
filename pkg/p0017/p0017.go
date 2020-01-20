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

package p0017

func letterCombinations(digits string) []string {
	if 0 == len(digits) {
		return []string{}
	}

	letters := map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}
	w := 1
	for i := 0; i < len(digits); i++ {
		w *= len(letters[int(digits[i]-'0')])
	}
	ret, buf := make([]string, w), make([]byte, len(digits))
	w = 0

	_walk := func(i int, fn interface{}) {
		if i >= len(digits) {
			ret[w] = string(buf)
			w++
			return
		}
		_fn := fn.(func(int, interface{}))
		letter := letters[int(digits[i]-'0')]
		for j := 0; j < len(letter); j++ {
			buf[i] = letter[j]
			_fn(i+1, fn)
		}
	}
	_walk(0, _walk)

	return ret
}
