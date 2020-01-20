// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0204

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target int
}

var values = []result{
	{
		arg1:   10,
		target: 4,
	},
}

type p0204TestSuite struct {
	suite.Suite
}

func (s *p0204TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, countPrimes(v.arg1))
	}
}

func TestP0204TestSuite(t *testing.T) {
	s := &p0204TestSuite{}
	suite.Run(t, s)
}
