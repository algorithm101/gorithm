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

package p0067

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	diffCount := len(a) - len(b)
	for diffCount > 0 {
		diffCount--
		b = "0" + b
	}

	ret := make([]rune, len(a)+1)
	carry := '0'
	for i := len(a) - 1; i >= 0; i-- {
		v := string(a[i]) + string(b[i]) + string(carry)
		switch {
		case v == "100" || v == "010" || v == "001":
			ret[i+1] = '1'
			carry = '0'
		case v == "000":
			ret[i+1] = '0'
			carry = '0'
		case v == "011" || v == "101" || v == "110":
			ret[i+1] = '0'
			carry = '1'
		case v == "111":
			ret[i+1] = '1'
			carry = '1'
		}
	}

	if carry == '1' {
		ret[0] = '1'
		return string(ret)
	}

	return string(ret[1:])
}
