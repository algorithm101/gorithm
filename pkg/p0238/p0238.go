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

package p0238

func productExceptSelf(nums []int) []int {
	r := make([]int, len(nums))
	r[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		r[i] = r[i-1] * nums[i]
	}

	r[len(nums)-1] = r[len(nums)-2]
	right := nums[len(nums)-1]
	for i := len(r) - 2; i > 0; i-- {
		r[i] = r[i-1] * right
		right *= nums[i]
	}
	r[0] = right

	return r
}
