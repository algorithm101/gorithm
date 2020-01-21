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

package p0063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if 0 == len(obstacleGrid) {
		return 0
	}

	row, col := len(obstacleGrid), len(obstacleGrid[0])
	ways := make([][]int, row) // TODO: 优化空间复杂度
	for i := 0; i < row; i++ {
		ways[i] = make([]int, col)
	}
	ways[0][0] = 1

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if obstacleGrid[i][j] == 1 {
				ways[i][j] = 0
				continue
			}

			switch {
			case i == 0 && j > 0:
				ways[i][j] = ways[i][j-1]
			case j == 0 && i > 0:
				ways[i][j] = ways[i-1][j]
			case i > 0 && j > 0:
				ways[i][j] = ways[i-1][j] + ways[i][j-1]
			}
		}
	}

	return ways[row-1][col-1]
}
