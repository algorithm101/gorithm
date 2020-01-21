// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0443

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []byte
	target []byte
}

var values = []result{
	{
		arg1:   []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
		target: []byte{'a', '2', 'b', '2', 'c', '3'},
	},
	{
		arg1:   []byte("a"),
		target: []byte("a"),
	},
	{
		arg1:   []byte("abbbbbbbbbbbb"),
		target: []byte("ab12"),
	},
}

type p0443TestSuite struct {
	suite.Suite
}

func (s *p0443TestSuite) Test() {
	for _, v := range values {
		r := v.arg1[0:compress(v.arg1)]
		s.Equal(v.target, r)
	}
}

func TestP0443TestSuite(t *testing.T) {
	s := &p0443TestSuite{}
	suite.Run(t, s)
}
