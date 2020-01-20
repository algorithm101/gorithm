// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0401

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target []string
}

var values = []result{
	{
		arg1: 1,
		target: []string{
			"0:01",
			"0:02",
			"0:04",
			"0:08",
			"0:16",
			"0:32",
			"1:00",
			"2:00",
			"4:00",
			"8:00",
		},
	},
}

type p0401TestSuite struct {
	suite.Suite
}

func (s *p0401TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, readBinaryWatch(v.arg1))
	}
}

func TestP0401TestSuite(t *testing.T) {
	s := &p0401TestSuite{}
	suite.Run(t, s)
}
