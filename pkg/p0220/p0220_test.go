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

package p0220

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0220TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	arg2   int
	arg3   int
	target bool
}

var values = []result{
	{
		arg1:   []int{1, 2, 3, 1},
		arg2:   3,
		arg3:   0,
		target: true,
	},
	{
		arg1:   []int{1, 0, 1, 1},
		arg2:   1,
		arg3:   2,
		target: true,
	},
	{
		arg1:   []int{1, 5, 9, 1, 5, 9},
		arg2:   2,
		arg3:   3,
		target: false,
	},
}

func (s *p0220TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, containsNearbyAlmostDuplicate(v.arg1, v.arg2, v.arg3))
	}
}

func TestP0220TestSuite(t *testing.T) {
	s := &p0220TestSuite{}
	suite.Run(t, s)
}
