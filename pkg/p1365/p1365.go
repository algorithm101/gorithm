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

package p1365

func smallerNumbersThanCurrent(nums []int) []int {
	v := [101]int{}
	for _, n := range nums {
		v[n]++
	}

	last := v[0]
	v[0] = 0
	for i := 1; i < len(v); i++ {
		v[i], last = v[i-1]+last, v[i]
	}

	r := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		r[i] = v[nums[i]]
	}
	return r
}
