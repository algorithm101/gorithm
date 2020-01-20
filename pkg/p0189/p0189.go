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

package p0189

func rotate(nums []int, k int) {
	k = k % len(nums)

	_reverse := func(nums []int, i int, j int) {
		for ; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	_reverse(nums, 0, len(nums)-k-1)
	_reverse(nums, len(nums)-k, len(nums)-1)
	_reverse(nums, 0, len(nums)-1)
}
