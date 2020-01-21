// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0033

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
		arg1:   []int{4, 5, 6, 7, 0, 1, 2},
		arg2:   0,
		target: 4,
	},
	{
		arg1:   []int{4, 5, 6, 7, 0, 1, 2},
		arg2:   3,
		target: -1,
	},
	{
		arg1:   []int{3, 1},
		arg2:   1,
		target: 1,
	},
}

type p0033TestSuite struct {
	suite.Suite
}

func (s *p0033TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, search(v.arg1, v.arg2))
	}
}

func TestP0033TestSuite(t *testing.T) {
	s := &p0033TestSuite{}
	suite.Run(t, s)
}
