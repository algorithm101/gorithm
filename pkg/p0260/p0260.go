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

package p0260

func singleNumber(nums []int) []int {
	v := nums[0]
	for i := 1; i < len(nums); i++ {
		v ^= nums[i]
	}

	x := v & (^(v - 1))

	v1, v2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&x == 0 {
			v1 ^= nums[i]
		} else {
			v2 ^= nums[i]
		}
	}

	if v1 < v2 {
		return []int{v1, v2}
	}
	return []int{v2, v1}
}
