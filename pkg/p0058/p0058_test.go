// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0058

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
		arg1:   "Hello World",
		target: 5,
	},
	{
		arg1:   "a ",
		target: 1,
	},
}

type p0058TestSuite struct {
	suite.Suite
}

func (s *p0058TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, lengthOfLastWord(v.arg1))
	}
}

func TestP0058TestSuite(t *testing.T) {
	s := &p0058TestSuite{}
	suite.Run(t, s)
}
