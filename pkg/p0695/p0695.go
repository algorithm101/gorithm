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

package p0695

func _countArea(grid [][]int, i int, j int) int {
	if grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0
	left, right, up, down := 0, 0, 0, 0
	if i > 0 {
		up = _countArea(grid, i-1, j)
	}
	if j > 0 {
		left = _countArea(grid, i, j-1)
	}
	if i < len(grid)-1 {
		down = _countArea(grid, i+1, j)
	}
	if j < len(grid[i])-1 {
		right = _countArea(grid, i, j+1)
	}

	return left + right + up + down + 1
}

func maxAreaOfIsland(grid [][]int) int {
	area := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			curr := _countArea(grid, i, j)
			if curr > area {
				area = curr
			}
		}
	}

	return area
}
