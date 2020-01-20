// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0067

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   string
	target string
}

var values = []result{
	{
		arg1:   "11",
		arg2:   "01",
		target: "100",
	},
	{
		arg1:   "1010",
		arg2:   "1011",
		target: "10101",
	},
}

type p0067TestSuite struct {
	suite.Suite
}

func (s *p0067TestSuite) Test() {
	for _, v := range values {
		s.Equal([]byte(v.target), []byte(addBinary(v.arg1, v.arg2)))
	}
}

func TestP0067TestSuite(t *testing.T) {
	s := &p0067TestSuite{}
	suite.Run(t, s)
}
