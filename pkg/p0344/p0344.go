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

package p0344

func reverseString(s string) string {
	sbyte := []byte(s)
	for i, j := 0, len(sbyte)-1; i < j; i, j = i+1, j-1 {
		sbyte[i], sbyte[j] = sbyte[j], sbyte[i]
	}
	return string(sbyte)
}
