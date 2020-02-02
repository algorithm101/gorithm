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

package p0046

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0046TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	target [][]int
}

var values = []result{
	{
		arg1: []int{1, 2, 3},
		target: [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 2, 1},
			{3, 1, 2},
		},
	},
}

func (s *p0046TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, permute(v.arg1))
	}
}

func TestP0046TestSuite(t *testing.T) {
	s := &p0046TestSuite{}
	suite.Run(t, s)
}
