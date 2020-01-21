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

package p0709

func toLowerCase(str string) string {
	ret := []byte(str)
	for i := 0; i < len(ret); i++ {
		if ret[i] >= 'A' && ret[i] <= 'Z' {
			ret[i] = ret[i] - 'A' + 'a'
		}
	}

	return string(ret)
}
