// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0520

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
		arg1:   "USA",
		target: true,
	},
	{
		arg1:   "leetcode",
		target: true,
	},
	{
		arg1:   "Google",
		target: true,
	},
	{
		arg1:   "FlaG",
		target: false,
	},
}

type p0520TestSuite struct {
	suite.Suite
}

func (s *p0520TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, detectCapitalUse(v.arg1))
	}
}

func TestP0520TestSuite(t *testing.T) {
	s := &p0520TestSuite{}
	suite.Run(t, s)
}
