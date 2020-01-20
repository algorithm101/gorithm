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

package p0414

import (
	"math"
)

func thirdMax(nums []int) int {
	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64

	for _, v := range nums {
		switch {
		case v > first:
			first, second, third = v, first, second
		case v < first && v > second:
			second, third = v, second
		case v < second && v > third:
			third = v
		}
	}

	if third == math.MinInt64 {
		return first
	}

	return third
}
