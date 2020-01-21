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

package p0557

func reverseWords(s string) string {
	begin, end := 0, 0
	ret := make([]byte, len(s)+1)

	for i := 0; i < len(s); i++ {
		ret[i] = s[i]
	}
	ret[len(s)] = ' '

	for end < len(ret) {
		for begin < len(ret) && ret[begin] == ' ' {
			begin++
		}

		end = begin + 1
		for end < len(ret) && ret[end] != ' ' {
			end++
		}

		if end >= len(ret) || begin == end {
			break
		}

		for i, j := begin, end-1; i < j; i, j = i+1, j-1 {
			ret[i], ret[j] = ret[j], ret[i]
		}
		begin = end + 1
	}

	return string(ret[:len(ret)-1])
}
