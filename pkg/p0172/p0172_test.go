// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0172

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
		arg1:   3,
		target: 0,
	},
	{
		arg1:   5,
		target: 1,
	},
}

type p0172TestSuite struct {
	suite.Suite
}

func (s *p0172TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, trailingZeroes(v.arg1))
	}
}

func TestP0172TestSuite(t *testing.T) {
	s := &p0172TestSuite{}
	suite.Run(t, s)
}
