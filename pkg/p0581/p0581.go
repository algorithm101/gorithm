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

package p0581

import "math"

func findUnsortedSubarray(nums []int) int {
	if 0 == len(nums) {
		return 0
	}

	n, beg, end := len(nums), -1, -2
	min, max := nums[n-1], nums[0]
	for i := 0; i < n; i++ {
		max = int(math.Max(float64(nums[i]), float64(max)))
		min = int(math.Min(float64(min), float64(nums[n-1-i])))

		if nums[i] < max {
			end = i
		}

		if nums[n-1-i] > min {
			beg = n - 1 - i
		}
	}

	return end - beg + 1
}
