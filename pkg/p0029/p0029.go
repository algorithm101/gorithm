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

package p0029

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if divisor == 0 || (dividend == math.MinInt32 && divisor == -1) {
		return math.MaxInt32
	}

	ret, dvd, dvs := 0, int(math.Abs(float64(dividend))), int(math.Abs(float64(divisor)))
	for dvs <= dvd {
		tmp := dvs
		mul := 1
		for dvd >= (tmp << 1) {
			tmp <<= 1
			mul <<= 1
		}

		ret += mul
		dvd -= tmp
	}

	if ((dividend < 0) && (divisor < 0)) || ((dividend >= 0) && (divisor >= 0)) {
		return ret
	}

	return 0 - ret
}
