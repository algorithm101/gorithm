// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0283

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{0, 1, 0, 3, 12},
		target: []int{1, 3, 12, 0, 0},
	},
}

type p0283TestSuite struct {
	suite.Suite
}

func (s *p0283TestSuite) Test() {
	for _, v := range values {
		moveZeroes(v.arg1)
		s.Equal(v.target, v.arg1)
	}
}

func TestP0283TestSuite(t *testing.T) {
	s := &p0283TestSuite{}
	suite.Run(t, s)
}
