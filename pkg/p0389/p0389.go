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

package p0389

func findTheDifference(s string, t string) byte {
	maps := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		count := maps[s[i]]
		maps[s[i]] = count + 1
	}

	for i := 0; i < len(t); i++ {
		count := maps[t[i]]
		if count < 1 {
			return t[i]
		}
		maps[t[i]] = count - 1
	}

	return 0
}
