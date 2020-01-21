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

package p0458

import (
	"math"
)

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if minutesToDie >= minutesToTest {
		return 0
	}

	statusCount := math.Floor(float64(minutesToTest/minutesToDie) + 1)
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(statusCount))))
}
