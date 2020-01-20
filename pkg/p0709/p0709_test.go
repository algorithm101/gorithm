// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0709

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
		arg1:   "Hello",
		target: "hello",
	},
	{
		arg1:   "here",
		target: "here",
	},
	{
		arg1:   "LOVELY",
		target: "lovely",
	},
}

type p0709TestSuite struct {
	suite.Suite
}

func (s *p0709TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, toLowerCase(v.arg1))
	}
}

func TestP0709TestSuite(t *testing.T) {
	s := &p0709TestSuite{}
	suite.Run(t, s)
}
