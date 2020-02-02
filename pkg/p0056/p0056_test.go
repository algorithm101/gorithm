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

package p0056

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0056TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []Interval
	target []Interval
}

var values = []result{
	{
		arg1: []Interval{
			{Start: 1, End: 3},
			{Start: 2, End: 6},
			{Start: 8, End: 10},
			{Start: 15, End: 18},
		},
		target: []Interval{
			{Start: 1, End: 6},
			{Start: 8, End: 10},
			{Start: 15, End: 18},
		},
	},
	{
		arg1: []Interval{
			{Start: 1, End: 4},
			{Start: 4, End: 5},
		},
		target: []Interval{
			{Start: 1, End: 5},
		},
	},
	{
		arg1: []Interval{
			{Start: 1, End: 4},
			{Start: 0, End: 4},
		},
		target: []Interval{
			{Start: 0, End: 4},
		},
	},
	{
		arg1: []Interval{
			{Start: 1, End: 4},
			{Start: 2, End: 3},
		},
		target: []Interval{
			{Start: 1, End: 4},
		},
	},
}

func (s *p0056TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, merge(v.arg1))
	}
}

func TestP0056TestSuite(t *testing.T) {
	s := &p0056TestSuite{}
	suite.Run(t, s)
}
