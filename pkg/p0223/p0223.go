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

package p0223

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	area := (C-A)*(D-B) + (G-E)*(H-F)

	_max := func(x int, y int) int {
		if x < y {
			return y
		}
		return x
	}
	_min := func(x int, y int) int {
		if x < y {
			return x
		}
		return y
	}
	_abs := func(x int) int {
		if x < 0 {
			return 0 - x
		}
		return x
	}

	x1 := _max(A, E)
	x2 := _min(C, G)

	if x1 >= x2 {
		return area
	}

	y1 := _max(B, F)
	y2 := _min(H, D)
	if y1 >= y2 {
		return area
	}

	return area - _abs(x2-x1)*_abs(y2-y1)
}
