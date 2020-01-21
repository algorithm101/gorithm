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

package p0504

import (
	"fmt"
	"math"
)

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	negative := false
	if num < 0 {
		negative = true
		num = int(math.Abs(float64(num)))
	}

	dest, w := make([]byte, 20), 19

	for num != 0 {
		v := num % 7
		dest[w] = fmt.Sprintf("%d", v)[0]
		w--
		num = num / 7
	}

	if negative {
		dest[w] = '-'
		w--
	}

	return string(dest[w+1:])
}
