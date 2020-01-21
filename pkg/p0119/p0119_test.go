// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0119

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
		arg1:   3,
		target: []int{1, 3, 3, 1},
	},
	{
		arg1:   4,
		target: []int{1, 4, 6, 4, 1},
	},
}

type p0119TestSuite struct {
	suite.Suite
}

func (s *p0119TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, getRow(v.arg1))
	}
}

func TestP0119TestSuite(t *testing.T) {
	s := &p0119TestSuite{}
	suite.Run(t, s)
}
