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

package p0006

func convert(s string, numRows int) string {
	dst := make([]byte, len(s))
	next := func(maxRows int, row int, prev int) int {
		if maxRows <= 1 {
			return prev + 1
		}

		if ((prev - row) % ((maxRows - 1) * 2)) != 0 {
			return prev + (row-1)*2
		}

		if row == maxRows {
			row = 1
		}
		return prev + (maxRows-row)*2
	}

	j := 0
	for row := 1; row <= numRows; row++ {
		prev := row
		for prev <= len(s) && j < len(s) {
			dst[j] = s[prev-1]
			j++
			prev = next(numRows, row, prev)
		}
	}

	return string(dst)
}
