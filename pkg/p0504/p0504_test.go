// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0504

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
		arg1:   100,
		target: "202",
	},
	{
		arg1:   -7,
		target: "-10",
	},
	{
		arg1:   0,
		target: "0",
	},
}

type p0504TestSuite struct {
	suite.Suite
}

func (s *p0504TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, convertToBase7(v.arg1))
	}
}

func TestP0504TestSuite(t *testing.T) {
	s := &p0504TestSuite{}
	suite.Run(t, s)
}
