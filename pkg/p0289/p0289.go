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

package p0289

func gameOfLife(board [][]int) {
	liveCount := func(board [][]int, i int, j int) int {
		count := 0
		for r := i - 1; r <= i+1; r++ {
			for c := j - 1; c <= j+1; c++ {
				if r == i && c == j {
					continue
				}
				if r >= 0 && c >= 0 && r < len(board) && c < len(board[r]) {
					if board[r][c] == 1 || board[r][c] == -1 {
						count++
					}
				}
			}
		}
		return count
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			c := liveCount(board, i, j)
			if board[i][j] == 1 {
				if c < 2 || c > 3 {
					board[i][j] = -1
				}
			} else if board[i][j] == 0 {
				if c == 3 {
					board[i][j] = 100
				}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == -1 {
				board[i][j] = 0
			} else if board[i][j] == 100 {
				board[i][j] = 1
			}
		}
	}
}
