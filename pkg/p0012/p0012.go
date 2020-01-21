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

package p0012

func intToRoman(num int) string {
	ret := ""

	maps := []struct {
		value int
		roman string
	}{
		{
			value: 1000,
			roman: "M",
		},
		{
			value: 900,
			roman: "CM",
		},
		{
			value: 500,
			roman: "D",
		},
		{
			value: 400,
			roman: "CD",
		},
		{
			value: 100,
			roman: "C",
		},
		{
			value: 90,
			roman: "XC",
		},
		{
			value: 50,
			roman: "L",
		},
		{
			value: 40,
			roman: "XL",
		},
		{
			value: 10,
			roman: "X",
		},
		{
			value: 9,
			roman: "IX",
		},
		{
			value: 5,
			roman: "V",
		},
		{
			value: 4,
			roman: "IV",
		},
		{
			value: 1,
			roman: "I",
		},
	}

	for num > 0 {
		for _, v := range maps {
			if num >= v.value {
				ret += v.roman
				num -= v.value
				break
			}
		}
	}

	return ret
}
