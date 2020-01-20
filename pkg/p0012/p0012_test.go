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

package p0012

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0012TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   int
	target string
}

var values = []result{
	{
		arg1:   3,
		target: "III",
	},
	{
		arg1:   4,
		target: "IV",
	},
	{
		arg1:   9,
		target: "IX",
	},
	{
		arg1:   58,
		target: "LVIII",
	},
	{
		arg1:   1994,
		target: "MCMXCIV",
	},
	{
		arg1:   20,
		target: "XX",
	},
}

func (s *p0012TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, intToRoman(v.arg1))
	}
}

func TestP0012TestSuite(t *testing.T) {
	s := &p0012TestSuite{}
	suite.Run(t, s)
}
