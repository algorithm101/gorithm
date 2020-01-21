// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0860

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
		arg1:   []int{5, 5, 5, 10, 20},
		target: true,
	},
	{
		arg1:   []int{5, 5, 10},
		target: true,
	},
	{
		arg1:   []int{10, 10},
		target: false,
	},
	{
		arg1:   []int{5, 5, 10, 10, 20},
		target: false,
	},
}

type p0860TestSuite struct {
	suite.Suite
}

func (s *p0860TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, lemonadeChange(v.arg1))
	}
}

func TestP0860TestSuite(t *testing.T) {
	s := &p0860TestSuite{}
	suite.Run(t, s)
}
