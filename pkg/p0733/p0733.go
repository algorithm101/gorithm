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

package p0733

import (
	"container/list"
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}

	stack, oldColor := list.New(), image[sr][sc]
	stack.PushBack([]int{sr, sc})
	for stack.Len() != 0 {
		e := stack.Front()
		stack.Remove(e)

		v := e.Value.([]int)
		r, c := v[0], v[1]

		image[r][c] = newColor
		if r > 0 && image[r-1][c] == oldColor {
			image[r-1][c] = -1
			stack.PushBack([]int{r - 1, c})
		}
		if r < len(image)-1 && image[r+1][c] == oldColor {
			image[r+1][c] = -1
			stack.PushBack([]int{r + 1, c})
		}
		if c > 0 && image[r][c-1] == oldColor {
			image[r][c-1] = -1
			stack.PushBack([]int{r, c - 1})
		}
		if c < len(image[r])-1 && image[r][c+1] == oldColor {
			image[r][c+1] = -1
			stack.PushBack([]int{r, c + 1})
		}
	}

	return image
}
