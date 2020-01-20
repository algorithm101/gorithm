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

package p0874

func robotSim(commands []int, obstacles [][]int) int {
	_isObstacle := func(x int, y int, obstacles [][]int) bool {
		for _, obstacle := range obstacles {
			if x == obstacle[0] && y == obstacle[1] {
				return true
			}
		}

		return false
	}
	_step := func(x int, y int, direct int, steps int, obstacles [][]int) (int, int) {
		n := 0
		switch direct {
		case 0:
			// y++
			for n = y + 1; n <= y+steps; n++ {
				if _isObstacle(x, n, obstacles) {
					return x, n - 1
				}
			}
			return x, n - 1
		case 1:
			// x++
			for n = x + 1; n <= x+steps; n++ {
				if _isObstacle(n, y, obstacles) {
					return n - 1, y
				}
			}
			return n - 1, y
		case 2:
			// y--
			for n = y - 1; n >= y-steps; n-- {
				if _isObstacle(x, n, obstacles) {
					return x, n + 1
				}
			}
			return x, n + 1
		case 3:
			// x--
			for n = x - 1; n >= x-steps; n-- {
				if _isObstacle(n, y, obstacles) {
					return n + 1, y
				}
			}
			return n + 1, y
		}

		panic("unreachable code")
	}

	x, y, direct, max := 0, 0, 0, 0

	for _, command := range commands {
		if command == -1 {
			direct = (direct + 1) % 4
			continue
		} else if command == -2 {
			if direct > 0 {
				direct = (direct - 1) % 4
			} else {
				direct = 3
			}
			continue
		}

		x, y = _step(x, y, direct, command, obstacles)
		if x*x+y*y > max {
			max = x*x + y*y
		}
	}
	return max
}
