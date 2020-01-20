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

package p0152

func maxProduct(nums []int) int {
	if 0 == len(nums) {
		return -1
	}

	_max := func(x int, y int) int {
		if x < y {
			return y
		}
		return x
	}
	_min := func(x int, y int) int {
		if x < y {
			return x
		}
		return y
	}

	max := nums[0]
	// 连续到当前位置的 最大/最小值
	dpmax, dpmin := make([]int, len(nums)), make([]int, len(nums))
	dpmax[0] = nums[0]
	dpmin[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dpmax[i] = _max(_max(nums[i], dpmax[i-1]*nums[i]), dpmin[i-1]*nums[i])
		dpmin[i] = _min(_min(dpmin[i-1]*nums[i], nums[i]), dpmax[i-1]*nums[i])

		max = _max(dpmax[i], max)
	}

	return max
}
