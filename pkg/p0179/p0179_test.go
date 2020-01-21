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

package p0179

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0179TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	target string
}

var values = []result{
	{
		arg1:   []int{10, 2},
		target: "210",
	},
	{
		arg1:   []int{3, 30, 34, 5, 9},
		target: "9534330",
	},
	{
		arg1:   []int{0, 0},
		target: "0",
	},
}

func (s *p0179TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, largestNumber(v.arg1))
	}
}

func TestP0179TestSuite(t *testing.T) {
	s := &p0179TestSuite{}
	suite.Run(t, s)
}
