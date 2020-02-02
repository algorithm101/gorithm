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

package p0213

func rob(nums []int) int {
	_max := func(x int, y int) int {
		if x < y {
			return y
		}
		return x
	}

	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return _max(nums[0], nums[1])
	}

	_rob := func(nums []int) int {
		max1, max2 := nums[0], _max(nums[0], nums[1])
		for i := 2; i < len(nums); i++ {
			max2, max1 = _max(max1+nums[i], max2), max2
		}

		return _max(max1, max2)
	}

	return _max(_rob(nums[0:len(nums)-1]), _rob(nums[1:]))
}
