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

package p0071

func simplifyPath(path string) string {
	p, w := make([]byte, len(path)+1), 0
	p[w] = '/'
	w++

	i, n := 0, len(path)
	for i < n {
		switch {
		case path[i] == '/':
			i++
		case path[i] == '.' && (i+1 == n || (path[i+1] == '/')):
			i += 2
		case path[i] == '.' && (i+1 < n && path[i+1] == '.') && (i+2 == n || path[i+2] == '/'):
			i += 3
			if w > 1 {
				w--
				for w > 1 && p[w] != '/' {
					w--
				}
			}
		default:
			if w > 1 {
				p[w] = '/'
				w++
			}

			for i < n && path[i] != '/' {
				p[w] = path[i]
				w++
				i++
			}
		}
	}

	return string(p[0:w])
}
