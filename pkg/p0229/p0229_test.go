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

package p0229

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0229TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{3, 2, 3},
		target: []int{3},
	},
	{
		arg1:   []int{1, 1, 1, 3, 3, 2, 2, 2},
		target: []int{1, 2},
	},
	{
		arg1:   []int{},
		target: []int{},
	},
	{
		arg1:   []int{1},
		target: []int{1},
	},
	{
		arg1:   []int{2, 2},
		target: []int{2},
	},
	{
		arg1:   []int{3, 3, 4},
		target: []int{3},
	},
	{
		arg1:   []int{6, 6, 6, 7, 7},
		target: []int{6, 7},
	},
}

func (s *p0229TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, majorityElement(v.arg1))
	}
}

func TestP0229TestSuite(t *testing.T) {
	s := &p0229TestSuite{}
	suite.Run(t, s)
}
