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

package p0200

import (
	"container/list"
)

func numIslands(grid [][]byte) int {
	if 0 == len(grid) {
		return 0
	}

	count := 0
	rowm, colm := len(grid), len(grid[0])

	type Node struct {
		row int
		col int
	}
	nodes := list.New()
	for i := 0; i < rowm; i++ {
		for j := 0; j < colm; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '2'
				nodes.PushBack(Node{row: i, col: j})
				for nodes.Len() != 0 {
					e := nodes.Front()
					nodes.Remove(e)
					n := e.Value.(Node)
					if n.row > 0 && grid[n.row-1][n.col] == '1' {
						grid[n.row-1][n.col] = '2'
						nodes.PushBack(Node{row: n.row - 1, col: n.col})
					}
					if n.row < rowm-1 && grid[n.row+1][n.col] == '1' {
						grid[n.row+1][n.col] = '2'
						nodes.PushBack(Node{row: n.row + 1, col: n.col})
					}
					if n.col > 0 && grid[n.row][n.col-1] == '1' {
						grid[n.row][n.col-1] = '2'
						nodes.PushBack(Node{row: n.row, col: n.col - 1})
					}
					if n.col < colm-1 && grid[n.row][n.col+1] == '1' {
						grid[n.row][n.col+1] = '2'
						nodes.PushBack(Node{row: n.row, col: n.col + 1})
					}
				}
				count++
			}
		}
	}

	for i := 0; i < rowm; i++ {
		for j := 0; j < colm; j++ {
			if grid[i][j] == '2' {
				grid[i][j] = '1'
			}
		}
	}

	return count
}
