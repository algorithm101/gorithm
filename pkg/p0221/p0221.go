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

package p0221

// TODO: 优化 dp 空间使用
func maximalSquare(matrix [][]byte) int {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return 0
	}

	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	for j := 0; j < col; j++ {
		dp[0][j] = int(matrix[0][j] - '0')
	}
	for i := 0; i < row; i++ {
		dp[i][0] = int(matrix[i][0] - '0')
	}

	_min := func(x int, y int, z int) int {
		m := x
		if y < m {
			m = y
		}
		if z < m {
			m = z
		}

		return m
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = _min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	max := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}

	return max * max
}
