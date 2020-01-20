// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0405

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target string
}

var values = []result{
	{
		arg1:   26,
		target: "1a",
	},
	{
		arg1:   -1,
		target: "ffffffff",
	},
	{
		arg1:   0,
		target: "0",
	},
}

type p0405TestSuite struct {
	suite.Suite
}

func (s *p0405TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, toHex(v.arg1))
	}
}

func TestP0405TestSuite(t *testing.T) {
	s := &p0405TestSuite{}
	suite.Run(t, s)
}
