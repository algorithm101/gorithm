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

package p0223

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0223TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   int
	arg2   int
	arg3   int
	arg4   int
	arg5   int
	arg6   int
	arg7   int
	arg8   int
	target int
}

var values = []result{
	{
		arg1:   -3,
		arg2:   0,
		arg3:   3,
		arg4:   4,
		arg5:   0,
		arg6:   -1,
		arg7:   9,
		arg8:   2,
		target: 45,
	},
	{
		arg1:   -2,
		arg2:   -2,
		arg3:   2,
		arg4:   2,
		arg5:   -2,
		arg6:   -2,
		arg7:   2,
		arg8:   2,
		target: 16,
	},
}

func (s *p0223TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, computeArea(v.arg1, v.arg2, v.arg3, v.arg4, v.arg5, v.arg6, v.arg7, v.arg8))
	}
}

func TestP0223TestSuite(t *testing.T) {
	s := &p0223TestSuite{}
	suite.Run(t, s)
}
