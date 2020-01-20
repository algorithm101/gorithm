// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0263

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target bool
}

var values = []result{
	{
		arg1:   6,
		target: true,
	},
	{
		arg1:   8,
		target: true,
	},
	{
		arg1:   14,
		target: false,
	},
	{
		arg1:   0,
		target: false,
	},
}

type p0263TestSuite struct {
	suite.Suite
}

func (s *p0263TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isUgly(v.arg1))
	}
}

func TestP0263TestSuite(t *testing.T) {
	s := &p0263TestSuite{}
	suite.Run(t, s)
}
