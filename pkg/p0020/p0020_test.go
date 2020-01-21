// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0020

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
		arg1:   "()",
		target: true,
	},
	{
		arg1:   "()[]{}",
		target: true,
	},
	{
		arg1:   "(]",
		target: false,
	},
	{
		arg1:   "([)]",
		target: false,
	},
	{
		arg1:   "{[]}",
		target: true,
	},
}

type p0020TestSuite struct {
	suite.Suite
}

func (s *p0020TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isValid(v.arg1))
	}
}

func TestP0020TestSuite(t *testing.T) {
	s := &p0020TestSuite{}
	suite.Run(t, s)
}
