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

package p0139

func wordBreak(s string, wordDict []string) bool {
	// timeout
	_f1 := func(s string, wordDict []string) bool {
		wordMap := make(map[string]bool, len(wordDict))
		for _, word := range wordDict {
			wordMap[word] = true
		}

		_walk := func(s string, fn interface{}) bool {
			if _, ok := wordMap[s]; ok {
				return true
			}

			_fn := fn.(func(string, interface{}) bool)
			for i := 0; i < len(s); i++ {
				w := s[0 : i+1]
				if _, ok := wordMap[w]; ok {
					if _fn(s[i+1:], fn) {
						return true
					}
				}
			}

			return false
		}

		return _walk(s, _walk)
	}
	_ = _f1

	_f2 := func(s string, wordDict []string) bool {
		wordMap := make(map[string]bool, len(wordDict))
		for _, word := range wordDict {
			wordMap[word] = true
		}

		// dp[i] expect that s[0:i+1] whether can wordBreak
		dp := make([]bool, len(s))
		dp[0] = wordMap[s[0:1]]
		for i := 1; i < len(s); i++ {
			if wordMap[s[0:i+1]] {
				dp[i] = true
				continue
			}

			for j := 0; j < i; j++ {
				if dp[j] && wordMap[s[j+1:i+1]] {
					// if dp[j] {
					dp[i] = true
				}
			}
		}
		return dp[len(s)-1]
	}
	_ = _f2

	return _f2(s, wordDict)
}
