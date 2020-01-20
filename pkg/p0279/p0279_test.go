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

package p0279

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0279TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   int
	target int
}

var values = []result{
	{
		arg1:   12,
		target: 3,
	},
	{
		arg1:   13,
		target: 2,
	},
	{
		arg1:   6405,
		target: 3,
	},
	{
		arg1:   6922,
		target: 2,
	},
}

func (s *p0279TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, numSquares(v.arg1))
	}
}

func TestP0279TestSuite(t *testing.T) {
	s := &p0279TestSuite{}
	suite.Run(t, s)
}
