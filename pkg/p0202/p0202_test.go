// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0202

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
		arg1:   19,
		target: true,
	},
	{
		arg1:   2,
		target: false,
	},
}

type p0202TestSuite struct {
	suite.Suite
}

func (s *p0202TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isHappy(v.arg1))
	}
}

func TestP0202TestSuite(t *testing.T) {
	s := &p0202TestSuite{}
	suite.Run(t, s)
}
