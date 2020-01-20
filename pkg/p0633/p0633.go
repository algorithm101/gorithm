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

package p0633

import (
	"math"
)

func judgeSquareSum(c int) bool {
	_isSqrtable := func(v int) bool {
		sqrt := int(math.Sqrt(float64(v)))
		return (sqrt * sqrt) == v
	}

	sqrt := int(math.Sqrt(float64(c)))
	for i := 0; i <= sqrt; i++ {
		diff := c - i*i
		if _isSqrtable(diff) {
			return true
		}
	}

	return false
}
