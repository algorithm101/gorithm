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

package p0036

func isValidSudoku(board [][]byte) bool {
	_isRowValid := func(board [][]byte) bool {
		for i := 0; i < len(board); i++ {
			mark := make([]byte, 9)
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					continue
				}
				num := int(board[i][j] - '1')
				if mark[num] == 1 {
					return false
				}
				mark[num] = 1
			}
		}

		return true
	}
	if !_isRowValid(board) {
		return false
	}

	_isColumnValid := func(board [][]byte) bool {
		for j := 0; j < len(board); j++ {
			mark := make([]byte, 9)
			for i := 0; i < 9; i++ {
				if board[i][j] == '.' {
					continue
				}
				num := int(board[i][j] - '1')
				if mark[num] == 1 {
					return false
				}
				mark[num] = 1
			}
		}

		return true
	}
	if !_isColumnValid(board) {
		return false
	}

	_isValid33 := func(board [][]byte, row int, column int) bool {
		mark := make([]byte, 9)
		for i := row; i < row+3; i++ {
			for j := column; j < column+3; j++ {
				if board[i][j] == '.' {
					continue
				}
				num := int(board[i][j] - '1')
				if mark[num] == 1 {
					return false
				}
				mark[num] = 1
			}
		}
		return true
	}
	for i := 0; i < 8; i += 3 {
		for j := 0; j < 8; j += 3 {
			if !_isValid33(board, i, j) {
				return false
			}
		}
	}
	return true
}
