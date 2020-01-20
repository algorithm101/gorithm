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

package p0059

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}

	direct, row, column := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		ret[row][column] = i
		switch direct {
		case 0:
			// ret[row][column] = i
			if column >= n-1 || ret[row][column+1] != 0 {
				direct = 1
				row++
			} else {
				column++
			}
		case 1:
			// ret[row][column] = i
			if row >= n-1 || ret[row+1][column] != 0 {
				direct = 2
				column--
			} else {
				row++
			}
		case 2:
			if column <= 0 || ret[row][column-1] != 0 {
				direct = 3
				row--
			} else {
				column--
			}
		case 3:
			if row <= 0 || ret[row-1][column] != 0 {
				direct = 0
				column++
			} else {
				row--
			}
		}
	}

	return ret
}
