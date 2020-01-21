// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0762

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   6,
		arg2:   10,
		target: 4,
	},
	{
		arg1:   10,
		arg2:   15,
		target: 5,
	},
	{
		arg1:   990,
		arg2:   1048,
		target: 28,
	},
}

type p0762TestSuite struct {
	suite.Suite
}

func (s *p0762TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, countPrimeSetBits(v.arg1, v.arg2))
	}
}

func TestP0762TestSuite(t *testing.T) {
	s := &p0762TestSuite{}
	suite.Run(t, s)
}
