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
package p0003

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target int
}

var values = []result{
	{
		arg1:   "abcabcbb",
		target: 3,
	},
	{
		arg1:   "bbbbb",
		target: 1,
	},
	{
		arg1:   "pwwkew",
		target: 3,
	},
}

type p0003TestSuite struct {
	suite.Suite
}

func (s *p0003TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, lengthOfLongestSubstring(v.arg1))
	}
}

func TestP0003TestSuite(t *testing.T) {
	s := &p0003TestSuite{}
	suite.Run(t, s)
}

func init() {

}
