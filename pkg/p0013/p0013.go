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

package p0013

func romanToInt(s string) int {
	maps1 := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	maps2 := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	ret := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 {
			char2 := s[i : i+2]
			nums2, ok := maps2[char2]
			if ok {
				i++
				ret += nums2
				continue
			}
		}

		num1 := maps1[rune(s[i])]
		ret += num1
	}

	return ret
}
