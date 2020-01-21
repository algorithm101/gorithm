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

package p0207

// TODO: 需要进一步优化
// 1: 这是 BFS 解法, 可以尝试使用 DFS
// 2: 有向图的表示方法可以修改为矩阵
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree, degreeCount := make(map[int]int), 0
	requireCourses := make(map[int][]int)

	for i := 0; i < numCourses; i++ {
		inDegree[i] = 0
	}

	for _, req := range prerequisites {
		inDegree[req[0]]++
		degreeCount++
		requireCourses[req[1]] = append(requireCourses[req[1]], req[0])
	}

	_filterInDegree0 := func(inDegree map[int]int) (int, bool) {
		for k, v := range inDegree {
			if v == 0 {
				// need remove it, so it should be process next time
				inDegree[k] = -1
				return k, true
			}
		}

		return -1, false
	}
	for degreeCount != 0 {
		course, ok := _filterInDegree0(inDegree)
		if !ok {
			return false
		}

		for _, nextCourse := range requireCourses[course] {
			degreeCount--
			inDegree[nextCourse]--
		}
	}

	return true
}
