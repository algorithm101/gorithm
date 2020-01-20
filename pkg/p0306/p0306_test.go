// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0306

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target bool
}

var values = []result{
	{
		arg1:   "112358",
		target: true,
	},
	{
		arg1:   "199100199",
		target: true,
	},
	{
		arg1:   "123",
		target: true,
	},
	{
		arg1:   "1023",
		target: false,
	},
	{
		arg1:   "199001200",
		target: false,
	},
	{
		arg1:   "101",
		target: true,
	},
}

type p0306TestSuite struct {
	suite.Suite
}

func (s *p0306TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isAdditiveNumber(v.arg1))
	}
}

func TestP0306TestSuite(t *testing.T) {
	s := &p0306TestSuite{}
	suite.Run(t, s)
}
