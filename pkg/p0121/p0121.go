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

package p0121

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0
	prevMaxPrice, prevMinPrice := prices[0], prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > prevMaxPrice {
			prevMaxPrice = prices[i]

			if prevMaxPrice-prevMinPrice > profit {
				profit = prevMaxPrice - prevMinPrice
			}
		} else if prices[i] < prevMinPrice {
			prevMinPrice = prices[i]
			prevMaxPrice = prices[i]
		}
	}
	return profit
}
