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

package p0081

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0081TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	arg2   int
	target bool
}

var values = []result{
	{
		arg1:   []int{2, 5, 6, 0, 0, 1, 2},
		arg2:   0,
		target: true,
	},
	{
		arg1:   []int{2, 5, 6, 0, 0, 1, 2},
		arg2:   3,
		target: false,
	},
	{
		arg1:   []int{1, 3},
		arg2:   3,
		target: true,
	},
	{
		arg1:   []int{1, 1, 3},
		arg2:   3,
		target: true,
	},
}

func (s *p0081TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, search(v.arg1, v.arg2))
	}
}

func TestP0081TestSuite(t *testing.T) {
	s := &p0081TestSuite{}
	suite.Run(t, s)
}
