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

package p0868

func binaryGap(N int) int {
	last, gap := -1, 0
	for i := 0; i < 32; i++ {
		if N&(0x01<<uint(i)) != 0 {
			if last == -1 {
				last = i
			} else {
				if i-last > gap {
					gap = i - last
				}
				last = i
			}
		}
	}

	return gap
}
