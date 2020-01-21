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

package p0263

func isUgly(num int) bool {
	if num == 0 {
		return false
	}

	divs := []int{2, 3, 5}
	for _, div := range divs {
		for num%div == 0 {
			num /= div
		}
	}

	return num == 1
}
