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

package p0048

func rotate(matrix [][]int) {
	l := len(matrix) - 1

	for i := 0; i <= l-1; i++ {
		for j := i; j <= l-i-1; j++ {
			matrix[i][j], matrix[j][l-i], matrix[l-i][l-j], matrix[l-j][i] = matrix[l-j][i], matrix[i][j], matrix[j][l-i], matrix[l-i][l-j]
		}
	}
}
