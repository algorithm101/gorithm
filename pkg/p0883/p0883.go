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

package p0883

func projectionArea(grid [][]int) int {
	area := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				area++
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		max := 0
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > max {
				max = grid[i][j]
			}
		}
		area += max
	}

	for j := 0; j < len(grid); j++ {
		max := 0
		for i := 0; i < len(grid); i++ {
			if grid[i][j] > max {
				max = grid[i][j]
			}
		}
		area += max
	}

	return area
}
