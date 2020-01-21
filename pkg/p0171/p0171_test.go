// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0171

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target int
}

var values = []result{
	{
		arg1:   "A",
		target: 1,
	},
	{
		arg1:   "AB",
		target: 28,
	},
	{
		arg1:   "ZY",
		target: 701,
	},
	{
		arg1:   "Z",
		target: 26,
	},
}

type p0171TestSuite struct {
	suite.Suite
}

func (s *p0171TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, titleToNumber(v.arg1))
	}
}

func TestP0171TestSuite(t *testing.T) {
	s := &p0171TestSuite{}
	suite.Run(t, s)
}
