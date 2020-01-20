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

package p0860

func lemonadeChange(bills []int) bool {
	count5, count10, count20 := 0, 0, 0
	for _, v := range bills {
		switch v {
		case 5:
			count5++
		case 10:
			if count5 <= 0 {
				return false
			}
			count5--
			count10++
		case 20:
			if (count10 == 0 && count5 < 3) ||
				(count5 < 1) {
				return false
			}

			count5--
			if count10 > 0 {
				count10--
			} else {
				count5 -= 2
			}

			count20++
		}
	}
	return true
}
