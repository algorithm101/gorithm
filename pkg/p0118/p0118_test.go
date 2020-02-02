// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0118

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target [][]int
}

var values = []result{
	{
		arg1: 5,
		target: [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		},
	},
}

type p0118TestSuite struct {
	suite.Suite
}

func (s *p0118TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, generate(v.arg1))
	}
}

func TestP0118TestSuite(t *testing.T) {
	s := &p0118TestSuite{}
	suite.Run(t, s)
}
