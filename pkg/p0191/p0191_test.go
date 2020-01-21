// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0191

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   uint32
	target int
}

var values = []result{
	{
		arg1:   11,
		target: 3,
	},
	{
		arg1:   128,
		target: 1,
	},
}

type p0191TestSuite struct {
	suite.Suite
}

func (s *p0191TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, hammingWeight(v.arg1))
	}
}

func TestP0191TestSuite(t *testing.T) {
	s := &p0191TestSuite{}
	suite.Run(t, s)
}
