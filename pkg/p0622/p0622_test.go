// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0622

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
}

var values = []result{}

type p0622TestSuite struct {
	suite.Suite
}

func (s *p0622TestSuite) Test() {
	for _, v := range values {
		_ = v
	}
}

func TestP0622TestSuite(t *testing.T) {
	s := &p0622TestSuite{}
	suite.Run(t, s)
}
