// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0521

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   string
	target int
}

var values = []result{
	{
		arg1:   "aba",
		arg2:   "cdc",
		target: 3,
	},
}

type p0521TestSuite struct {
	suite.Suite
}

func (s *p0521TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findLUSlength(v.arg1, v.arg2))
	}
}

func TestP0521TestSuite(t *testing.T) {
	s := &p0521TestSuite{}
	suite.Run(t, s)
}
