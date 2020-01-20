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

package p0215

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0215TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   []int{3, 2, 1, 5, 6, 4},
		arg2:   2,
		target: 5,
	},
	{
		arg1:   []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
		arg2:   4,
		target: 4,
	},
	{
		arg1:   []int{7, 6, 5, 4, 3, 2, 1},
		arg2:   5,
		target: 3,
	},
}

func (s *p0215TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findKthLargest(v.arg1, v.arg2))
	}
}

func TestP0215TestSuite(t *testing.T) {
	s := &p0215TestSuite{}
	suite.Run(t, s)
}
