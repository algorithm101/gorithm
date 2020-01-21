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

package p0492

import "math"

func constructRectangle(area int) []int {
	l := int(math.Sqrt(float64(area)))

	for ; l <= area; l++ {
		if area%l == 0 {
			w := area / l
			if w >= l {
				return []int{w, l}
			}
			return []int{l, w}
		}
	}

	return nil
}
