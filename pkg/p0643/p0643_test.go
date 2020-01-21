// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0643

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   int
	target float64
}

var values = []result{
	{
		arg1:   []int{1, 12, -5, -6, 50, 3},
		arg2:   4,
		target: 12.75,
	},
}

type p0643TestSuite struct {
	suite.Suite
}

func (s *p0643TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findMaxAverage(v.arg1, v.arg2))
	}
}

func TestP0643TestSuite(t *testing.T) {
	s := &p0643TestSuite{}
	suite.Run(t, s)
}
