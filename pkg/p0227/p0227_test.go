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

package p0227

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0227TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   string
	target int
}

var values = []result{
	{
		arg1:   "3+2*2",
		target: 7,
	},
	{
		arg1:   " 3/2 ",
		target: 1,
	},
	{
		arg1:   " 3+5 / 2 ",
		target: 5,
	},
	{
		arg1:   "100000000/1/2/3/4/5/6/7/8/9/10",
		target: 27,
	},
}

func (s *p0227TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, calculate(v.arg1))
	}
}

func TestP0227TestSuite(t *testing.T) {
	s := &p0227TestSuite{}
	suite.Run(t, s)
}
