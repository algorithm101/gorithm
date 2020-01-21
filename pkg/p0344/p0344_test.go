// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0344

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target string
}

var values = []result{
	{
		arg1:   "hello",
		target: "olleh",
	},
}

type p0344TestSuite struct {
	suite.Suite
}

func (s *p0344TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, reverseString(v.arg1))
	}
}

func TestP0344TestSuite(t *testing.T) {
	s := &p0344TestSuite{}
	suite.Run(t, s)
}
