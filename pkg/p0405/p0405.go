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

package p0405

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	ret := []byte("00000000")
	chars := "0123456789abcdef"
	for i := 0; i < 32; i += 4 {
		ret[7-i/4] = chars[num&0xF]
		num = num >> 4
	}

	for i := 0; i < len(ret); i++ {
		if ret[i] != '0' {
			return string(ret[i:])
		}
	}

	return ""
}
