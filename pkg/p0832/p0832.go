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

package p0832

func flipAndInvertImage(A [][]int) [][]int {
	for r := 0; r < len(A); r++ {
		for i, j := 0, len(A[r])-1; i < j; i, j = i+1, j-1 {
			A[r][i], A[r][j] = A[r][j], A[r][i]
		}
	}

	for r := 0; r < len(A); r++ {
		for c := 0; c < len(A[r]); c++ {
			A[r][c] = 1 - A[r][c]
		}
	}

	return A
}
