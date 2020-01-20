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

package p0050

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0050TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   float64
	arg2   int
	target float64
}

var values = []result{
	{
		arg1:   2.00000,
		arg2:   10,
		target: 1024.00000,
	},
	{
		arg1:   2.10000,
		arg2:   3,
		target: 9.26100,
	},
	{
		arg1:   2.00000,
		arg2:   -2,
		target: 0.25000,
	},
}

func (s *p0050TestSuite) Test() {
	for _, v := range values {
		s.InDelta(v.target, myPow(v.arg1, v.arg2), 0.000001)
	}
}

func TestP0050TestSuite(t *testing.T) {
	s := &p0050TestSuite{}
	suite.Run(t, s)
}
