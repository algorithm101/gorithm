// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0633

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target bool
}

var values = []result{
	{
		arg1:   5,
		target: true,
	},
	{
		arg1:   3,
		target: false,
	},
}

type p0633TestSuite struct {
	suite.Suite
}

func (s *p0633TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, judgeSquareSum(v.arg1))
	}
}

func TestP0633TestSuite(t *testing.T) {
	s := &p0633TestSuite{}
	suite.Run(t, s)
}
