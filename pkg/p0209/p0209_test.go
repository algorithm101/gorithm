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

package p0209

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0209TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   int
	arg2   []int
	target int
}

var values = []result{
	{
		arg1:   7,
		arg2:   []int{2, 3, 1, 2, 4, 3},
		target: 2,
	},
	{
		arg1:   100,
		arg2:   []int{},
		target: 0,
	},
	{
		arg1:   11,
		arg2:   []int{1, 2, 3, 4, 5},
		target: 3,
	},
}

func (s *p0209TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, minSubArrayLen(v.arg1, v.arg2))
	}
}

func TestP0209TestSuite(t *testing.T) {
	s := &p0209TestSuite{}
	suite.Run(t, s)
}
