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

package p0566

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if 0 == len(nums) {
		return nums
	}

	or, oc := len(nums), len(nums[0])
	if c*r != or*oc {
		return nums
	}

	ret := make([][]int, r)
	for i := 0; i < len(ret); i++ {
		ret[i] = make([]int, c)
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			index := (i)*oc + j
			nr, nc := index/c, index%c
			ret[nr][nc] = nums[i][j]
		}
	}

	return ret
}
