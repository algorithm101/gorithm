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

package p0766

func isToeplitzMatrix(matrix [][]int) bool {
	for r := 0; r < len(matrix); r++ {
		dest, c := matrix[r][0], 0
		for r+c < len(matrix) && c < len(matrix[r+c]) {
			if matrix[r+c][c] != dest {
				return false
			}
			c++
		}
	}

	for c := 1; c < len(matrix[0]); c++ {
		dest, r := matrix[0][c], 0
		for r < len(matrix) && r+c < len(matrix[r]) {
			if matrix[r][c+r] != dest {
				return false
			}
			r++
		}
	}
	return true
}
