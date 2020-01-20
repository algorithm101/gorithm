// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0657

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target bool
}

var values = []result{
	{
		arg1:   "UD",
		target: true,
	},
	{
		arg1:   "LL",
		target: false,
	},
}

type p0657TestSuite struct {
	suite.Suite
}

func (s *p0657TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, judgeCircle(v.arg1))
	}
}

func TestP0657TestSuite(t *testing.T) {
	s := &p0657TestSuite{}
	suite.Run(t, s)
}
