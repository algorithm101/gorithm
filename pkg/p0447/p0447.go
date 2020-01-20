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

package p0447

func numberOfBoomerangs(points [][]int) int {
	count := 0

	for i, point := range points {
		dist := make(map[int]int)
		for j, p := range points {
			if i != j {
				distance := (point[0]-p[0])*(point[0]-p[0]) + (point[1]-p[1])*(point[1]-p[1])
				dist[distance] = dist[distance] + 1
			}
		}

		for _, v := range dist {
			if v >= 2 {
				count += v * (v - 1)
			}
		}
	}

	return count
}
