// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0482

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   int
	target string
}

var values = []result{
	{
		arg1:   "5F3Z-2e-9-w",
		arg2:   4,
		target: "5F3Z-2E9W",
	},
	{
		arg1:   "2-5g-3-J",
		arg2:   2,
		target: "2-5G-3J",
	},
	{
		arg1:   "--a-a-a-a--",
		arg2:   2,
		target: "AA-AA",
	},
	{
		arg1:   "---",
		arg2:   3,
		target: "",
	},
}

type p0482TestSuite struct {
	suite.Suite
}

func (s *p0482TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, licenseKeyFormatting(v.arg1, v.arg2))
	}
}

func TestP0482TestSuite(t *testing.T) {
	s := &p0482TestSuite{}
	suite.Run(t, s)
}
