// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0507

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
		arg1:   28,
		target: true,
	},
	{
		arg1:   6,
		target: true,
	},
	{
		arg1:   1,
		target: false,
	},
}

type p0507TestSuite struct {
	suite.Suite
}

func (s *p0507TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, checkPerfectNumber(v.arg1))
	}
}

func TestP0507TestSuite(t *testing.T) {
	s := &p0507TestSuite{}
	suite.Run(t, s)
}
