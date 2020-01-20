// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0746

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target int
}

var values = []result{
	{
		arg1:   []int{10, 15, 20},
		target: 15,
	},
	{
		arg1:   []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
		target: 6,
	},
}

type p0746TestSuite struct {
	suite.Suite
}

func (s *p0746TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, minCostClimbingStairs(v.arg1))
	}
}

func TestP0746TestSuite(t *testing.T) {
	s := &p0746TestSuite{}
	suite.Run(t, s)
}
