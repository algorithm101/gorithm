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

package p0605

func canPlaceFlowers(flowerbed []int, n int) bool {
	last, count := -1, 0

	for i := 0; i < len(flowerbed); i++ {
		switch {
		case flowerbed[i] == 1:
			last = i
		case flowerbed[i] == 0 && ((last == -1) || (i-last >= 2)):
			if i+1 >= len(flowerbed) || flowerbed[i+1] == 0 {
				count++
				last = i
			}
		}
	}

	return count >= n
}
