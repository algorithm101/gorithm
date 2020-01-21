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

package p0806

func numberOfLines(widths []int, S string) []int {
	lines, offset := 1, 0

	for i := 0; i < len(S); i++ {
		if offset+widths[S[i]-'a'] > 100 {
			lines++
			offset = widths[S[i]-'a']
		} else {
			offset += widths[S[i]-'a']
		}
	}

	return []int{lines, offset}
}
