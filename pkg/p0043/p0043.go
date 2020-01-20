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

package p0043

func multiply(num1 string, num2 string) string {
	ret := make([]byte, len(num1)+len(num2))
	for i := 0; i < len(ret); i++ {
		ret[i] = '0'
	}

	carry := 0
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			n2, k := int(num2[j]-'0'), int(ret[i+j+1]-'0')
			r := n1*n2 + k + carry

			carry = r / 10
			ret[i+j+1] = byte(r%10) + '0'
		}

		ret[i] = byte(carry) + '0'
		carry = 0
	}

	i := 0
	for ret[i] == '0' && i < len(ret)-1 {
		i++
	}

	if i == len(ret) {
		return "0"
	}

	return string(ret[i:])
}
