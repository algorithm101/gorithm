// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0070

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target int
}

var values = []result{
	{
		arg1:   2,
		target: 2,
	},
	{
		arg1:   3,
		target: 3,
	},
	{
		arg1:   4,
		target: 5,
	},
}

type p0070TestSuite struct {
	suite.Suite
}

func (s *p0070TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, climbStairs(v.arg1))
	}
}

func TestP0070TestSuite(t *testing.T) {
	s := &p0070TestSuite{}
	suite.Run(t, s)
}
