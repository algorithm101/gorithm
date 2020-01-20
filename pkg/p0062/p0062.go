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

package p0062

func uniquePaths(m int, n int) int {
	m, n = n, m
	ways := make([][]int, m) // TODO: 继续优化空间复杂度
	for i := 0; i < m; i++ {
		ways[i] = make([]int, n)
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row == 0 || col == 0 {
				ways[row][col] = 1
			} else {
				ways[row][col] = ways[row-1][col] + ways[row][col-1]
			}
		}
	}

	return ways[m-1][n-1]
}
