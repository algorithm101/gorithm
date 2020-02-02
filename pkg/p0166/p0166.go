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

package p0166

import (
	"math"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	_helper := func(numerator int, denominator int) string {
		indexes := make(map[int]int)
		indexes[numerator] = 0
		ret, w := "", 1
		for numerator != 0 {
			numerator *= 10
			r := numerator / denominator
			numerator = numerator % denominator
			ret += strconv.Itoa(r)

			if i, ok := indexes[numerator]; ok {
				// Repeated
				repeated := ret[i:]
				ret = ret[0:i]
				ret = ret + "(" + repeated + ")"
				break
			}
			indexes[numerator] = w
			w++
		}
		return ret
	}
	isNegative := false
	if numerator*denominator < 0 {
		isNegative = true
	}

	numerator = int(math.Abs(float64(numerator)))
	denominator = int(math.Abs(float64(denominator)))

	first := "0"
	if numerator >= denominator {
		first = strconv.Itoa(numerator / denominator)
		numerator %= denominator
	}

	second := _helper(numerator, denominator)
	if second != "" {
		first = first + "." + second
	}

	if isNegative && first != "0" {
		first = "-" + first
	}
	return first
}
