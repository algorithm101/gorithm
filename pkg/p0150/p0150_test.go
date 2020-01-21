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

package p0150

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0150TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []string
	target int
}

var values = []result{
	{
		arg1:   []string{"2", "1", "+", "3", "*"},
		target: 9,
	},
	{
		arg1:   []string{"4", "13", "5", "/", "+"},
		target: 6,
	},
	{
		arg1:   []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
		target: 22,
	},
}

func (s *p0150TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, evalRPN(v.arg1))
	}
}

func TestP0150TestSuite(t *testing.T) {
	s := &p0150TestSuite{}
	suite.Run(t, s)
}
