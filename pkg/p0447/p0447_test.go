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

package p0447

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0447TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]int
	target int
}

var values = []result{
	{
		arg1: [][]int{
			{0, 0},
			{1, 0},
			{2, 0},
		},
		target: 2,
	},
}

func (s *p0447TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, numberOfBoomerangs(v.arg1))
	}
}

func TestP0447TestSuite(t *testing.T) {
	s := &p0447TestSuite{}
	suite.Run(t, s)
}
