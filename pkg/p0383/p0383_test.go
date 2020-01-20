// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0383

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
		arg1:   "a",
		arg2:   "b",
		target: false,
	},
	{
		arg1:   "aa",
		arg2:   "ab",
		target: false,
	},
	{
		arg1:   "aa",
		arg2:   "aab",
		target: true,
	},
}

type p0383TestSuite struct {
	suite.Suite
}

func (s *p0383TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, canConstruct(v.arg1, v.arg2))
	}
}

func TestP0383TestSuite(t *testing.T) {
	s := &p0383TestSuite{}
	suite.Run(t, s)
}
