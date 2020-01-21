// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0326

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
		arg1:   27,
		target: true,
	},
	{
		arg1:   0,
		target: false,
	},
	{
		arg1:   9,
		target: true,
	},
	{
		arg1:   45,
		target: false,
	},
}

type p0326TestSuite struct {
	suite.Suite
}

func (s *p0326TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isPowerOfThree(v.arg1))
	}
}

func TestP0326TestSuite(t *testing.T) {
	s := &p0326TestSuite{}
	suite.Run(t, s)
}
