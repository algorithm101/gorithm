// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0028

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   string
	target int
}

var values = []result{
	{
		arg1:   "aaaaa",
		arg2:   "bba",
		target: -1,
	},
	{
		arg1:   "abba",
		arg2:   "bba",
		target: 1,
	},
}

type p0028TestSuite struct {
	suite.Suite
}

func (s *p0028TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, strStr(v.arg1, v.arg2))
	}
}

func TestP0028TestSuite(t *testing.T) {
	s := &p0028TestSuite{}
	suite.Run(t, s)
}
