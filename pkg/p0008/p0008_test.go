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

package p0008

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0008TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   string
	target int
}

var values = []result{
	{
		arg1:   "42",
		target: 42,
	},
	{
		arg1:   "   -42",
		target: -42,
	},
	{
		arg1:   "4193 with words",
		target: 4193,
	},
	{
		arg1:   "words and 987",
		target: 0,
	},
	{
		arg1:   "-91283472332",
		target: -2147483648,
	},
	{
		arg1:   "3.14159",
		target: 3,
	},
	{
		arg1:   "9223372036854775808",
		target: 2147483647,
	},
	{
		arg1:   "2147483648",
		target: 2147483647,
	},
}

func (s *p0008TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, myAtoi(v.arg1))
	}
}

func TestP0008TestSuite(t *testing.T) {
	s := &p0008TestSuite{}
	suite.Run(t, s)
}
