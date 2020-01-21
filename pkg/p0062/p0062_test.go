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

package p0062

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0062TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   int
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   3,
		arg2:   2,
		target: 3,
	},
	{
		arg1:   7,
		arg2:   3,
		target: 28,
	},
}

func (s *p0062TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, uniquePaths(v.arg1, v.arg2))
	}
}

func TestP0062TestSuite(t *testing.T) {
	s := &p0062TestSuite{}
	suite.Run(t, s)
}
