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

package p0219

func containsNearbyDuplicate(nums []int, k int) bool {
	maps := make(map[int]bool, len(nums))
	for i := 0; i < len(nums) && i <= k; i++ {
		_, ok := maps[nums[i]]
		if ok {
			return true
		}
		maps[nums[i]] = true
	}

	for i := k + 1; i < len(nums); i++ {
		delete(maps, nums[i-k-1])
		_, ok := maps[nums[i]]
		if ok {
			return true
		}
		maps[nums[i]] = true
	}

	return false
}
