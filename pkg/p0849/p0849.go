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

package p0849

func maxDistToClosest(seats []int) int {
	maxDist, lastSeatedIndex := 0, -1
	for i, v := range seats {
		if v == 1 {
			if lastSeatedIndex == -1 {
				maxDist = i
			} else {
				currMaxDist := (i - lastSeatedIndex) / 2
				if currMaxDist > maxDist {
					maxDist = currMaxDist
				}
			}
			lastSeatedIndex = i
		}
	}
	if lastSeatedIndex != len(seats)-1 {
		currMaxDist := len(seats) - lastSeatedIndex - 1
		if currMaxDist > maxDist {
			maxDist = currMaxDist
		}
	}

	return maxDist
}
