// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0038

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target string
}

var values = []result{
	{
		arg1:   1,
		target: "1",
	},
	{
		arg1:   4,
		target: "1211",
	},
}

type p0038TestSuite struct {
	suite.Suite
}

func (s *p0038TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, countAndSay(v.arg1))
	}
}

func TestP0038TestSuite(t *testing.T) {
	s := &p0038TestSuite{}
	suite.Run(t, s)
}
