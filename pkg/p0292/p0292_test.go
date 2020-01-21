// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0292

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
		arg1:   4,
		target: false,
	},
}

type p0292TestSuite struct {
	suite.Suite
}

func (s *p0292TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, canWinNim(v.arg1))
	}
}

func TestP0292TestSuite(t *testing.T) {
	s := &p0292TestSuite{}
	suite.Run(t, s)
}
