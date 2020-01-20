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

package p0836

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	_max := func(x int, y int) int {
		if x > y {
			return x
		}

		return y
	}
	_min := func(x int, y int) int {
		if x > y {
			return y
		}

		return x
	}

	return _max(rec1[0], rec2[0]) < _min(rec1[2], rec2[2]) && _max(rec1[1], rec2[1]) < _min(rec1[3], rec2[3])
}
