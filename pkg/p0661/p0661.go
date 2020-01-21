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

package p0661

func imageSmoother(M [][]int) [][]int {
	_calc := func(m [][]int, x int, y int) int {
		sum, count := 0, 0
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				if i >= 0 && j >= 0 && i < len(m) && j < len(m[i]) {
					count++
					sum += m[i][j]
				}
			}
		}

		return sum / count
	}

	ret := make([][]int, len(M))
	for i := 0; i < len(M); i++ {
		ret[i] = make([]int, len(M[i]))
	}

	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[i]); j++ {
			ret[i][j] = _calc(M, i, j)
		}
	}

	return ret
}
