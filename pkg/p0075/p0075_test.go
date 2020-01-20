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

package p0075

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0075TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{2, 0, 2, 1, 1, 0},
		target: []int{0, 0, 1, 1, 2, 2},
	},
	{
		arg1:   []int{2, 0, 1},
		target: []int{0, 1, 2},
	},
}

func (s *p0075TestSuite) Test() {
	for _, v := range values {
		sortColors(v.arg1)
		s.Equal(v.target, v.arg1)
	}
}

func TestP0075TestSuite(t *testing.T) {
	s := &p0075TestSuite{}
	suite.Run(t, s)
}
