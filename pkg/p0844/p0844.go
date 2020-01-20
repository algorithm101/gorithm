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

package p0844

func backspaceCompare(S string, T string) bool {
	removeBackspace := func(src string) string {
		dst := make([]byte, len(src))
		end := 0

		for i := 0; i < len(src); i++ {
			switch src[i] {
			case '#':
				if end > 0 {
					end--
				}
			default:
				dst[end] = src[i]
				end++
			}
		}

		return string(dst[:end])
	}

	return removeBackspace(S) == removeBackspace(T)
}
