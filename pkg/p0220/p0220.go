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

package p0220

// TODO: 降低复杂度, 这个复杂度为 O(kn), 降低为 O(n*lg(k))
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	_abs := func(x int) int {
		if x < 0 {
			return 0 - x
		}

		return x
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if _abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}

	return false
}
