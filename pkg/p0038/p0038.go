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

package p0038

func countAndSay(n int) string {
	ret := "1"
	now := ""
	for n > 1 {
		n--
		for i := 0; i < len(ret); i++ {
			switch {
			case ((i+1 < len(ret)) && ret[i+1] == ret[i]) && ((i+2 < len(ret)) && ret[i+2] == ret[i]):
				now = now + "3" + string(ret[i])
				i += 2
			case ((i+1 < len(ret)) && ret[i+1] == ret[i]):
				now = now + "2" + string(ret[i])
				i++
			default:
				now = now + "1" + string(ret[i])
			}
		}
		ret = now
		now = ""
	}
	return ret
}
