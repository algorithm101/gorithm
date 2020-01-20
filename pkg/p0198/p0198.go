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

package p0198

func rob(nums []int) int {
	if 0 == len(nums) {
		return 0
	}

	if len(nums) < 2 {
		return nums[0]
	}

	_max := func(v1 int, v2 int) int {
		if v1 < v2 {
			return v2
		}
		return v1
	}

	prevMax, currMax := nums[0], _max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		m := _max(prevMax+nums[i], currMax)
		prevMax, currMax = currMax, m
	}

	return _max(currMax, prevMax)
}
