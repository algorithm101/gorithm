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

package p0705

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string //nolint
	target string //nolint
}

var values = []result{ //nolint
	{
		arg1:   "babad",
		target: "bab",
	},
	{
		arg1:   "cbbd",
		target: "bb",
	},
}

type p0705TestSuite struct {
	suite.Suite
}

func (s *p0705TestSuite) Test() {
	// for _, v := range values {
	// 	s.Equal(v.target, longestPalindrome(v.arg1))
	// }
}

func TestP0705TestSuite(t *testing.T) {
	s := &p0705TestSuite{}
	suite.Run(t, s)
}
