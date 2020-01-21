// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0561

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target int
}

var values = []result{
	{
		arg1:   []int{1, 4, 3, 2},
		target: 4,
	},
}

type p0561TestSuite struct {
	suite.Suite
}

func (s *p0561TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, arrayPairSum(v.arg1))
	}
}

func TestP0561TestSuite(t *testing.T) {
	s := &p0561TestSuite{}
	suite.Run(t, s)
}
