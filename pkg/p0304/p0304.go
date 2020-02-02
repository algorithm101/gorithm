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

//nolint
package p0304

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sum := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		sum[i] = make([]int, len(matrix[i]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			switch {
			case i == 0 && j == 0:
				sum[i][j] = matrix[i][j]
			case i == 0 && j > 0:
				sum[i][j] = sum[i][j-1] + matrix[i][j]
			case i > 0 && j == 0:
				sum[i][j] = sum[i-1][j] + matrix[i][j]
			default:
				sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i][j]
			}
		}
	}

	return NumMatrix{
		sum: sum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	switch {
	case row1 == 0 && col1 == 0:
		return this.sum[row2][col2]
	case row1 > 0 && col1 == 0:
		return this.sum[row2][col2] - this.sum[row1-1][col2]
	case row1 == 0 && col1 > 0:
		return this.sum[row2][col2] - this.sum[row2][col1-1]
	default:
		return this.sum[row2][col2] - this.sum[row2][col1-1] - this.sum[row1-1][col2] + this.sum[row1-1][col1-1]
	}
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
