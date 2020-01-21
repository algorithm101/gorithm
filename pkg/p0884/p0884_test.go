// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0884

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   string
	target []string
}

var values = []result{
	{
		arg1:   "this apple is sweet",
		arg2:   "this apple is sour",
		target: []string{"sweet", "sour"},
	},
	{
		arg1:   "apple apple",
		arg2:   "banana",
		target: []string{"banana"},
	},
}

type p0884TestSuite struct {
	suite.Suite
}

func (s *p0884TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, uncommonFromSentences(v.arg1, v.arg2))
	}
}

func TestP0884TestSuite(t *testing.T) {
	s := &p0884TestSuite{}
	suite.Run(t, s)
}
