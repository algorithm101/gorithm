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

package p0598

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0598TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   int
	arg2   int
	arg3   [][]int
	target int
}

var values = []result{
	{
		arg1: 3,
		arg2: 3,
		arg3: [][]int{
			{2, 2},
			{3, 3},
		},
		target: 4,
	},
}

func (s *p0598TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, maxCount(v.arg1, v.arg2, v.arg3))
	}
}

func TestP0598TestSuite(t *testing.T) {
	s := &p0598TestSuite{}
	suite.Run(t, s)
}
