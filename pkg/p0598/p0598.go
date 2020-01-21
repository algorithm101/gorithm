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

package p0598

import (
	"math"
)

func maxCount(m int, n int, ops [][]int) int {
	if 0 == len(ops) {
		return m * n
	}

	minm, minn := math.MaxInt32, math.MaxInt32
	for _, op := range ops {
		if op[0] < minm {
			minm = op[0]
		}

		if op[1] < minn {
			minn = op[1]
		}
	}

	return minm * minn
}
