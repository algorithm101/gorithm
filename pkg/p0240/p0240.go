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

package p0240

import (
	"container/list"
)

// TODO: 可继续优化
func searchMatrix(matrix [][]int, target int) bool {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return false
	}

	type Rect struct {
		x1 int
		y1 int
		x2 int
		y2 int
	}

	row, col := len(matrix)-1, len(matrix[0])-1
	l := list.New()
	l.PushBack(Rect{
		x1: 0,
		y1: 0,
		x2: row,
		y2: col,
	})
	for l.Len() != 0 {
		e := l.Front()
		l.Remove(e)
		rect := e.Value.(Rect)

		if rect.x1 > rect.x2 || rect.y1 > rect.y2 {
			continue
		}
		if rect.x1 == rect.x2 && rect.y1 == rect.y2 {
			if target == matrix[rect.x1][rect.y1] {
				return true
			}
			continue
		}

		// fmt.Println(rect)

		midx, midy := (rect.x1+rect.x2)/2, (rect.y1+rect.y2)/2
		mid := matrix[midx][midy]
		switch {
		case mid == target:
			return true
		case mid > target:
			// left top
			if midx-1 >= 0 && midy-1 >= 0 && matrix[midx-1][midy-1] >= target {
				l.PushBack(Rect{
					x1: rect.x1,
					y1: rect.y1,
					x2: midx - 1,
					y2: midy - 1,
				})
			}

			// right up
			if midx-1 >= 0 && matrix[rect.x1][midy] <= target && matrix[midx-1][rect.y2] >= target {
				n := Rect{
					x1: rect.x1,
					y1: midy,
					x2: midx - 1,
					y2: rect.y2,
				}
				l.PushBack(n)
			}

			// left down
			if midy-1 >= 0 && matrix[midx][rect.y1] <= target && matrix[rect.x2][midy-1] >= target {
				n := Rect{
					x1: midx,
					y1: rect.y1,
					x2: rect.x2,
					y2: midy - 1,
				}
				l.PushBack(n)
			}

		case mid < target:
			// right down
			if midx+1 <= row && midy+1 <= col && matrix[midx+1][midy+1] <= target {
				l.PushBack(Rect{
					x1: midx + 1,
					y1: midy + 1,
					x2: rect.x2,
					y2: rect.y2,
				})
			}

			// right up
			if midy+1 <= col && matrix[rect.x1][midy+1] <= target && matrix[midx][rect.y2] >= target {
				n := Rect{
					x1: rect.x1,
					y1: midy + 1,
					x2: midx,
					y2: rect.y2,
				}
				// fmt.Println("Push: ", rect, "  ", n)
				l.PushBack(n)
			}

			// left down
			if midx+1 <= row && matrix[midx+1][rect.y1] <= target && matrix[rect.x2][midy] >= target {
				n := Rect{
					x1: midx + 1,
					y1: rect.y1,
					x2: rect.x2,
					y2: midy,
				}
				l.PushBack(n)
			}
		}
	}

	return false
}
