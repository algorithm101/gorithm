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

package p0035

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
		arg1:   []int{1, 3, 5, 6},
		arg2:   5,
		target: 2,
	},
	{
		arg1:   []int{1, 3, 5, 6},
		arg2:   2,
		target: 1,
	},
	{
		arg1:   []int{1, 3, 5, 6},
		arg2:   7,
		target: 4,
	},
	{
		arg1:   []int{1, 3, 5, 6},
		arg2:   0,
		target: 0,
	},
}

type p0035TestSuite struct {
	suite.Suite
}

func (s *p0035TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, searchInsert(v.arg1, v.arg2))
	}
}

func TestP0035TestSuite(t *testing.T) {
	s := &p0035TestSuite{}
	suite.Run(t, s)
}
