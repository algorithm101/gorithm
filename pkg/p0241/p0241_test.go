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

package p0241

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0241TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   string
	target []int
}

var values = []result{
	{
		arg1:   "2-1-1",
		target: []int{2, 0},
	},
	{
		arg1:   "2*3-4*5",
		target: []int{-34, -10, -14, -10, 10},
	},
}

func (s *p0241TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, diffWaysToCompute(v.arg1))
	}
}

func TestP0241TestSuite(t *testing.T) {
	s := &p0241TestSuite{}
	suite.Run(t, s)
}
