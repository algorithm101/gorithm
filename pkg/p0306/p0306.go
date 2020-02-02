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

package p0306

func isAdditiveNumber(num string) bool {
	_add := func(n1 string, n2 string) string {
		if len(n1) < len(n2) {
			n1, n2 = n2, n1
		}
		r, w, plus := make([]byte, len(n1)+1), len(n1), 0

		for i, j := len(n1)-1, len(n2)-1; i >= 0 || j >= 0; {
			v := plus
			switch {
			case i < 0:
				v += int(n2[j] - '0')
				j--
			case j < 0:
				v += int(n1[i] - '0')
				i--
			default:
				v = v + int(n2[j]-'0') + int(n1[i]-'0')
				i--
				j--
			}

			plus = v / 10
			v = v % 10
			r[w] = byte(v) + '0'
			w--
		}
		if plus > 0 {
			r[w] = byte(plus) + '0'
			w--
		}

		return string(r[w+1:])
	}

	_is := func(v1 string, v2 string, num string) bool {
	START:
		sum := _add(v1, v2)
		if len(sum) > len(num) || (len(v1) > 1 && v1[0] == '0') || (len(v2) > 1 && v2[0] == '0') {
			return false
		}

		switch {
		case len(sum) > len(num):
			return false
		case len(sum) == len(num):
			return sum == num
		default:
			if num[0:len(sum)] == sum {
				v1, v2 = v2, sum
				num = num[len(sum):]
				goto START
			}
			return false
		}
	}

	for i := 1; i < len(num)/2+1; i++ {
		for j := i + 1; j < len(num)-i+1; j++ {
			if _is(num[0:i], num[i:j], num[j:]) {
				return true
			}
		}
	}

	return false
}
