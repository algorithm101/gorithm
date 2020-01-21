// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0069

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
		arg1:   4,
		target: 2,
	},
	{
		arg1:   8,
		target: 2,
	},
	{
		arg1:   1,
		target: 1,
	},
}

type p0069TestSuite struct {
	suite.Suite
}

func (s *p0069TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, mySqrt(v.arg1))
	}
}

func TestP0069TestSuite(t *testing.T) {
	s := &p0069TestSuite{}
	suite.Run(t, s)
}
