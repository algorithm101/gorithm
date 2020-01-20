// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0492

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target []int
}

var values = []result{
	{
		arg1:   1,
		target: []int{1, 1},
	},
	{
		arg1:   4,
		target: []int{2, 2},
	},
}

type p0492TestSuite struct {
	suite.Suite
}

func (s *p0492TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, constructRectangle(v.arg1))
	}
}

func TestP0492TestSuite(t *testing.T) {
	s := &p0492TestSuite{}
	suite.Run(t, s)
}
