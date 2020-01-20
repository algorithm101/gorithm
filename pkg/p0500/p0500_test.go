// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0500

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []string
	target []string
}

var values = []result{
	{
		arg1:   []string{"Hello", "Alaska", "Dad", "Peace"},
		target: []string{"Alaska", "Dad"},
	},
}

type p0500TestSuite struct {
	suite.Suite
}

func (s *p0500TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findWords(v.arg1))
	}
}

func TestP0500TestSuite(t *testing.T) {
	s := &p0500TestSuite{}
	suite.Run(t, s)
}
