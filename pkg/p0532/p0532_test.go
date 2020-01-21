// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0532

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   []int{3, 1, 4, 1, 5},
		arg2:   2,
		target: 2,
	},
	{
		arg1:   []int{1, 2, 3, 4, 5},
		arg2:   1,
		target: 4,
	},
	{
		arg1:   []int{1, 3, 1, 5, 4},
		arg2:   0,
		target: 1,
	},
	{
		arg1:   []int{1, 2, 3, 4, 5},
		arg2:   -1,
		target: 0,
	},
	{
		arg1:   []int{1, 1, 1, 2, 1},
		arg2:   1,
		target: 1,
	},
}

type p0532TestSuite struct {
	suite.Suite
}

func (s *p0532TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findPairs(v.arg1, v.arg2))
	}
}

func TestP0532TestSuite(t *testing.T) {
	s := &p0532TestSuite{}
	suite.Run(t, s)
}
