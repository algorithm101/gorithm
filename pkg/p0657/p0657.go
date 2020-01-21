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

package p0657

func judgeCircle(moves string) bool {
	hor, ver := 0, 0

	for _, step := range moves {
		switch step {
		case 'L':
			ver++
		case 'R':
			ver--
		case 'U':
			hor++
		case 'D':
			hor--
		}
	}

	return hor == 0 && ver == 0
}
