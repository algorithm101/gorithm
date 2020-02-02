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

package p0048

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0048TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]int
	target [][]int
}

var values = []result{
	{
		arg1: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		target: [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		},
	},
	{
		arg1: [][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		},
		target: [][]int{
			{15, 13, 2, 5},
			{14, 3, 4, 1},
			{12, 6, 8, 9},
			{16, 7, 10, 11},
		},
	},
}

func (s *p0048TestSuite) Test() {
	for _, v := range values {
		rotate(v.arg1)
		s.Equal(v.target, v.arg1)
	}
}

func TestP0048TestSuite(t *testing.T) {
	s := &p0048TestSuite{}
	suite.Run(t, s)
}
