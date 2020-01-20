// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0717

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target bool
}

var values = []result{
	{
		arg1:   []int{1, 0, 0},
		target: true,
	},
	{
		arg1:   []int{1, 1, 1, 0},
		target: false,
	},
}

type p0717TestSuite struct {
	suite.Suite
}

func (s *p0717TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isOneBitCharacter(v.arg1))
	}
}

func TestP0717TestSuite(t *testing.T) {
	s := &p0717TestSuite{}
	suite.Run(t, s)
}
