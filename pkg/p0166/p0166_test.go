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

package p0166

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0166TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   int
	arg2   int
	target string
}

var values = []result{
	{
		arg1:   1,
		arg2:   2,
		target: "0.5",
	},
	{
		arg1:   2,
		arg2:   1,
		target: "2",
	},
	{
		arg1:   2,
		arg2:   3,
		target: "0.(6)",
	},
	{
		arg1:   -50,
		arg2:   8,
		target: "-6.25",
	},
}

func (s *p0166TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, fractionToDecimal(v.arg1, v.arg2))
	}
}

func TestP0166TestSuite(t *testing.T) {
	s := &p0166TestSuite{}
	suite.Run(t, s)
}
