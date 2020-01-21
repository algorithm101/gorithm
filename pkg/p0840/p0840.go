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

package p0840

func numMagicSquaresInside(grid [][]int) int {
	_isMagicSquare := func(grid [][]int, i int, j int) int {
		bits := make([]int, 10)

		for row := i; row < int(i+3); row++ {
			sum := 0
			for col := j; col < int(j+3); col++ {
				if grid[row][col] > 9 || grid[row][col] < 1 {
					return 0
				}

				if bits[grid[row][col]] == 1 {
					return 0
				}
				bits[grid[row][col]] = 1
				sum += grid[row][col]
			}

			if sum != 15 {
				return 0
			}
		}

		for col := j; col < int(j+3); col++ {
			sum := 0
			for row := i; row < int(i+3); row++ {
				sum += grid[row][col]
			}

			if sum != 15 {
				return 0
			}
		}

		// 正斜对角
		{
			sum := 0
			sum += grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2]
			if sum != 15 {
				return 0
			}
		}

		// 反斜对角
		{
			sum := 0
			sum += grid[i+2][j] + grid[i+1][j+1] + grid[i][j+2]
			if sum != 15 {
				return 0
			}
		}

		return 1
	}
	count := 0

	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[i])-2; j++ {
			count += _isMagicSquare(grid, i, j)
		}
	}

	return count
}
