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

package p0153

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0153TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	target int
}

var values = []result{
	{
		arg1:   []int{3, 4, 5, 1, 2},
		target: 1,
	},
	{
		arg1:   []int{4, 5, 6, 7, 0, 1, 2},
		target: 0,
	},
}

func (s *p0153TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findMin(v.arg1))
	}
}

func TestP0153TestSuite(t *testing.T) {
	s := &p0153TestSuite{}
	suite.Run(t, s)
}
