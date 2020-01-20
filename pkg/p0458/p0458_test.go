// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0458

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	arg2   int
	arg3   int
	target int
}

var values = []result{
	{
		arg1:   1,
		arg2:   1,
		arg3:   1,
		target: 0,
	},
	{
		arg1:   1000,
		arg2:   15,
		arg3:   60,
		target: 5,
	},
}

type p0458TestSuite struct {
	suite.Suite
}

func (s *p0458TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, poorPigs(v.arg1, v.arg2, v.arg3))
	}
}

func TestP0458TestSuite(t *testing.T) {
	s := &p0458TestSuite{}
	suite.Run(t, s)
}
