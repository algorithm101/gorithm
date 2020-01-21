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

package p0300

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0300TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	target int
}

var values = []result{
	{
		arg1:   []int{10, 9, 2, 5, 3, 7, 101, 18},
		target: 4,
	},
	{
		arg1:   []int{2, 2},
		target: 1,
	},
}

func (s *p0300TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, lengthOfLIS(v.arg1))
	}
}

func TestP0300TestSuite(t *testing.T) {
	s := &p0300TestSuite{}
	suite.Run(t, s)
}
