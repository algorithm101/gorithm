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

package p0029

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0029TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   int
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   10,
		arg2:   3,
		target: 3,
	},
	{
		arg1:   7,
		arg2:   -3,
		target: -2,
	},
	{
		arg1:   1,
		arg2:   1,
		target: 1,
	},
	{
		arg1:   -2147483648,
		arg2:   -1,
		target: 2147483647,
	},
}

func (s *p0029TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, divide(v.arg1, v.arg2))
	}
}

func TestP0029TestSuite(t *testing.T) {
	s := &p0029TestSuite{}
	suite.Run(t, s)
}
