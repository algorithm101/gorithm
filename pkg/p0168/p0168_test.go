// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0168

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
		target: "A",
	},
	{
		arg1:   28,
		target: "AB",
	},
	{
		arg1:   701,
		target: "ZY",
	},
	{
		arg1:   26,
		target: "Z",
	},
}

type p0168TestSuite struct {
	suite.Suite
}

func (s *p0168TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, convertToTitle(v.arg1))
	}
}

func TestP0168TestSuite(t *testing.T) {
	s := &p0168TestSuite{}
	suite.Run(t, s)
}
