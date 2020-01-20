// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0819

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   []string
	target string
}

var values = []result{
	{
		arg1:   "Bob hit a ball, the hit BALL flew far after it was hit.",
		arg2:   []string{"hit"},
		target: "ball",
	},
}

type p0819TestSuite struct {
	suite.Suite
}

func (s *p0819TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, mostCommonWord(v.arg1, v.arg2))
	}
}

func TestP0819TestSuite(t *testing.T) {
	s := &p0819TestSuite{}
	suite.Run(t, s)
}
