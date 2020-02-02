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

package p0077

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0077TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   int
	arg2   int
	target [][]int
}

var values = []result{
	{
		arg1: 4,
		arg2: 2,
		target: [][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4},
		},
	},
}

func (s *p0077TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, combine(v.arg1, v.arg2))
	}
}

func TestP0077TestSuite(t *testing.T) {
	s := &p0077TestSuite{}
	suite.Run(t, s)
}
