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

package p0771

func numJewelsInStones(J string, S string) int {
	has, count := make([]int, 256), 0
	for i := 0; i < len(J); i++ {
		has[J[i]] = 1
	}

	for i := 0; i < len(S); i++ {
		count += has[S[i]]
	}
	return count
}
