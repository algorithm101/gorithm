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

package p0210

import (
	"container/list"
)

// TODO: 需要进一步优化
// 1: 这是 BFS 解法, 可以尝试使用 DFS
// 2: 有向图的表示方法可以修改为矩阵
func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make(map[int]int)
	requireCourses := make(map[int][]int)

	for i := 0; i < numCourses; i++ {
		inDegree[i] = 0
	}

	for _, req := range prerequisites {
		inDegree[req[0]]++
		requireCourses[req[1]] = append(requireCourses[req[1]], req[0])
	}

	ret, w := make([]int, numCourses), 0
	l := list.New()

	for k, v := range inDegree {
		if v == 0 {
			inDegree[k] = -1
			l.PushBack(k)
		}
	}

	for l.Len() != 0 {
		e := l.Front()
		l.Remove(e)
		course := e.Value.(int)

		ret[w] = course
		w++

		for _, nextCourse := range requireCourses[course] {
			switch {
			case inDegree[nextCourse] == 1:
				l.PushBack(nextCourse)
				inDegree[nextCourse] = -1
			case inDegree[nextCourse] > 1:
				inDegree[nextCourse]--
			}
		}
	}

	for _, v := range inDegree {
		if v != -1 {
			return []int{}
		}
	}

	return ret
}
