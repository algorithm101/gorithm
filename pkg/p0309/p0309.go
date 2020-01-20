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

package p0309

func maxProfit(prices []int) int {
	_max := func(x int, y int) int {
		if x > y {
			return x
		}
		return y
	}

	if 0 == len(prices) {
		return 0
	}

	sdp, bdp := make([]int, len(prices)), make([]int, len(prices))
	sdp[0] = 0
	bdp[0] = 0 - prices[0]

	for i := 1; i < len(prices); i++ {
		sdp[i] = _max(sdp[i-1], bdp[i-1]+prices[i])
		if i >= 2 {
			bdp[i] = _max(bdp[i-1], sdp[i-2]-prices[i])
		} else {
			bdp[i] = _max(bdp[i-1], 0-prices[i])
		}
	}

	return sdp[len(prices)-1]
}
