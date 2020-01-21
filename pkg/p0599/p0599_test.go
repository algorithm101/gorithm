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

package p0599

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0599TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []string
	arg2   []string
	target []string
}

var values = []result{
	{
		arg1:   []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		arg2:   []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
		target: []string{"Shogun"},
	},
	{
		arg1:   []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		arg2:   []string{"KFC", "Shogun", "Burger King"},
		target: []string{"Shogun"},
	},
}

func (s *p0599TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findRestaurant(v.arg1, v.arg2))
	}
}

func TestP0599TestSuite(t *testing.T) {
	s := &p0599TestSuite{}
	suite.Run(t, s)
}
