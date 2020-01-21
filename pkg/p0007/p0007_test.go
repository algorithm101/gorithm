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

package p0007

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target int
}

var values = []result{
	{
		arg1:   123,
		target: 321,
	},
	{
		arg1:   -123,
		target: -321,
	},
	{
		arg1:   120,
		target: 21,
	},
}

type p0007TestSuite struct {
	suite.Suite
}

func (s *p0007TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, reverse(v.arg1))
	}
}

func TestP0007TestSuite(t *testing.T) {
	s := &p0007TestSuite{}
	suite.Run(t, s)
}
