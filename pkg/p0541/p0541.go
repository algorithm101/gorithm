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

package p0541

func reverseStr(s string, k int) string {
	dest := []byte(s)

	start := 0
	for start < len(s) {
		end := start + k
		if end > len(s) {
			end = len(s)
		}
		for i, j := start, end-1; i < j; i, j = i+1, j-1 {
			dest[i], dest[j] = dest[j], dest[i]
		}

		start += 2 * k
	}
	return string(dest)
}
