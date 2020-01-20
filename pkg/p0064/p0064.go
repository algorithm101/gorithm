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

package p0064

func minPathSum(grid [][]int) int {
	if 0 == len(grid) {
		return 0
	}

	_min := func(x int, y int) int {
		if x < y {
			return x
		}
		return y
	}

	row, col := len(grid), len(grid[0])
	ways := make([][]int, row) // TODO: 优化空间复杂度
	for i := 0; i < row; i++ {
		ways[i] = make([]int, col)
	}
	ways[0][0] = grid[0][0]

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			switch {
			case i == 0 && j > 0:
				ways[i][j] = grid[i][j] + ways[i][j-1]
			case j == 0 && i > 0:
				ways[i][j] = ways[i-1][j] + grid[i][j]
			case i > 0 && j > 0:
				ways[i][j] = _min(ways[i-1][j], ways[i][j-1]) + grid[i][j]
			}
		}
	}

	return ways[row-1][col-1]
}
