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

package p0054

func spiralOrder(matrix [][]int) []int {
	if 0 == len(matrix) {
		return make([]int, 0)
	}

	ret, w := make([]int, len(matrix)*len(matrix[0])), 0
	row, column := 0, 0

	for 2*row < len(matrix) {
		// 顶
		for j := column; j < len(matrix[row])-column; j++ {
			ret[w] = matrix[row][j]
			w++
		}

		// 处理只剩一行的情况, 在这种情况下会超出
		if w == len(ret) {
			break
		}

		// 右边
		for i := row + 1; i < len(matrix)-row-1; i++ {
			ret[w] = matrix[i][len(matrix[row])-1-column]
			w++
		}

		// 处理只剩一行的情况, 在这种情况下会超出
		if w == len(ret) {
			break
		}

		// 底
		for j := len(matrix[row]) - column - 1; j >= column; j-- {
			ret[w] = matrix[len(matrix)-1-row][j]
			w++
		}

		// 处理只剩一行的情况, 在这种情况下会超出
		if w == len(ret) {
			break
		}

		// 左边
		for i := len(matrix) - 1 - row - 1; i > row; i-- {
			ret[w] = matrix[i][column]
			w++
		}

		row, column = row+1, column+1
	}

	return ret
}
