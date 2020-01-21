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

package p0448

func findDisappearedNumbers(nums []int) []int {
	_abs := func(v int) int {
		if v < 0 {
			return 0 - v
		}
		return v
	}

	for _, v := range nums {
		v = _abs(v)
		i := nums[v-1]
		if i > 0 {
			nums[v-1] = 0 - i
		}
	}

	ret := make([]int, 0)
	for i, v := range nums {
		if v > 0 {
			ret = append(ret, i+1)
		}
	}

	return ret
}
