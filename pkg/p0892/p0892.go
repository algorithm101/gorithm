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

package p0892

import "math"

func surfaceArea(grid [][]int) int {
	area := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				area += 6*grid[i][j] - (grid[i][j]-1)*2
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				if i > 0 && grid[i-1][j] > 0 {
					area -= int(math.Min(float64(grid[i][j]), float64(grid[i-1][j])))
				}

				if i < len(grid)-1 && grid[i+1][j] > 0 {
					area -= int(math.Min(float64(grid[i][j]), float64(grid[i+1][j])))
				}

				if j > 0 && grid[i][j-1] > 0 {
					area -= int(math.Min(float64(grid[i][j]), float64(grid[i][j-1])))
				}

				if j < len(grid[i])-1 && grid[i][j+1] > 0 {
					area -= int(math.Min(float64(grid[i][j]), float64(grid[i][j+1])))
				}
			}
		}
	}

	return area
}
