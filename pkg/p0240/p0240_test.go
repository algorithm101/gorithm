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

package p0240

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0240TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]int
	arg2   int
	target bool
}

var values = []result{
	{
		arg1: [][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		},
		arg2:   5,
		target: true,
	},
	{
		arg1: [][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		},
		arg2:   20,
		target: false,
	},
	{
		arg1: [][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		},
		arg2:   24,
		target: true,
	},
	{
		arg1: [][]int{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		},
		arg2:   5,
		target: true,
	},
	{
		arg1: [][]int{
			{1, 3, 5},
		},
		arg2:   1,
		target: true,
	},
	{
		arg1: [][]int{
			{3, 6, 9, 12, 17, 22},
			{5, 11, 11, 16, 22, 26},
			{10, 11, 14, 16, 24, 31},
			{10, 15, 17, 17, 29, 31},
			{14, 17, 20, 23, 34, 37},
			{19, 21, 22, 28, 37, 40},
			{22, 22, 24, 32, 37, 43},
		},
		arg2:   20,
		target: true,
	},
}

func (s *p0240TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, searchMatrix(v.arg1, v.arg2))
	}
}

func TestP0240TestSuite(t *testing.T) {
	s := &p0240TestSuite{}
	suite.Run(t, s)
}
