// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0696

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
		arg1:   "00110011",
		target: 6,
	},
	{
		arg1:   "10101",
		target: 4,
	},
}

type p0696TestSuite struct {
	suite.Suite
}

func (s *p0696TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, countBinarySubstrings(v.arg1))
	}
}

func TestP0696TestSuite(t *testing.T) {
	s := &p0696TestSuite{}
	suite.Run(t, s)
}
