// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0434

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target int
}

var values = []result{
	{
		arg1:   "Hello, my name is John",
		target: 5,
	},
	{
		arg1:   ",,,,",
		target: 1,
	},
}

type p0434TestSuite struct {
	suite.Suite
}

func (s *p0434TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, countSegments(v.arg1))
	}
}

func TestP0434TestSuite(t *testing.T) {
	s := &p0434TestSuite{}
	suite.Run(t, s)
}
