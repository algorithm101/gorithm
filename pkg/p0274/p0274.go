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

package p0274

import (
	"sort"
)

// 二分查找见: p0275
func hIndex(citations []int) int {
	sort.Ints(citations)

	for i := 0; i < len(citations); i++ {
		nums := len(citations) - i
		if citations[i] >= nums {
			return nums
		}
	}

	return 0
}
