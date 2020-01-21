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

package p0017

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0017TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   string
	target []string
}

var values = []result{
	{
		arg1:   "23",
		target: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	},
}

func (s *p0017TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, letterCombinations(v.arg1))
	}
}

func TestP0017TestSuite(t *testing.T) {
	s := &p0017TestSuite{}
	suite.Run(t, s)
}
