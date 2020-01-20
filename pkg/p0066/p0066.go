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

package p0066

func plusOne(digits []int) []int {
	ret := make([]int, len(digits)+1)
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		now := digits[i] + carry
		if now > 9 {
			carry = 1
			now = now % 10
		} else {
			carry = 0
		}
		ret[i+1] = now
	}
	if carry != 0 {
		ret[0] = carry
	} else {
		ret = ret[1:]
	}

	return ret
}
