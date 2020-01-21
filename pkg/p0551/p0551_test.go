// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0551

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
		arg1:   "PPALLP",
		target: true,
	},
	{
		arg1:   "PPALLL",
		target: false,
	},
	{
		arg1:   "LLPPPLL",
		target: true,
	},
}

type p0551TestSuite struct {
	suite.Suite
}

func (s *p0551TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, checkRecord(v.arg1))
	}
}

func TestP0551TestSuite(t *testing.T) {
	s := &p0551TestSuite{}
	suite.Run(t, s)
}
