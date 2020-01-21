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

package p0079

func exist(board [][]byte, word string) bool {
	if 0 == len(word) {
		return true
	}
	if 0 == len(board) {
		return false
	}

	_exist := func(board [][]byte, word string, index int, i int, j int, fn interface{}) bool {
		_fn := fn.(func([][]byte, string, int, int, int, interface{}) bool)
		if index >= len(word) {
			return true
		}

		if i > 0 && board[i-1][j] == word[index] {
			board[i-1][j] = '0'
			if _fn(board, word, index+1, i-1, j, fn) {
				return true
			}
			board[i-1][j] = word[index]
		}

		if i < len(board)-1 && board[i+1][j] == word[index] {
			board[i+1][j] = '0'
			if _fn(board, word, index+1, i+1, j, fn) {
				return true
			}
			board[i+1][j] = word[index]
		}

		if j > 0 && board[i][j-1] == word[index] {
			board[i][j-1] = '0'
			if _fn(board, word, index+1, i, j-1, fn) {
				return true
			}
			board[i][j-1] = word[index]
		}

		if j < len(board[i])-1 && board[i][j+1] == word[index] {
			board[i][j+1] = '0'
			if _fn(board, word, index+1, i, j+1, fn) {
				return true
			}
			board[i][j+1] = word[index]
		}

		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				board[i][j] = '0'
				if _exist(board, word, 1, i, j, _exist) {
					return true
				}
				board[i][j] = word[0]
			}
		}
	}

	return false
}
