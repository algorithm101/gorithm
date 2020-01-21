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

package p0016

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   []int{-1, 2, 1, -4},
		arg2:   1,
		target: 2,
	},
	{
		arg1:   []int{0, 0, 0},
		arg2:   1,
		target: 0,
	},
	{
		arg1:   []int{1, 1, 1, 0},
		arg2:   -100,
		target: 2,
	},
}

type p0016TestSuite struct {
	suite.Suite
}

func (s *p0016TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, threeSumClosest(v.arg1, v.arg2))
	}
}

func TestP0016TestSuite(t *testing.T) {
	s := &p0016TestSuite{}
	suite.Run(t, s)
}
