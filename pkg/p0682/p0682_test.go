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

package p0682

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0682TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []string
	target int
}

var values = []result{
	{
		arg1:   []string{"5", "2", "C", "D", "+"},
		target: 30,
	},
	{
		arg1:   []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
		target: 27,
	},
}

func (s *p0682TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, calPoints(v.arg1))
	}
}

func TestP0682TestSuite(t *testing.T) {
	s := &p0682TestSuite{}
	suite.Run(t, s)
}
