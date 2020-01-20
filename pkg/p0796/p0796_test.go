// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0796

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   string
	target bool
}

var values = []result{
	{
		arg1:   "abcde",
		arg2:   "cdeab",
		target: true,
	},
	{
		arg1:   "abcde",
		arg2:   "abced",
		target: false,
	},
	{
		arg1:   "aa",
		arg2:   "a",
		target: false,
	},
}

type p0796TestSuite struct {
	suite.Suite
}

func (s *p0796TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, rotateString(v.arg1, v.arg2))
	}
}

func TestP0796TestSuite(t *testing.T) {
	s := &p0796TestSuite{}
	suite.Run(t, s)
}
