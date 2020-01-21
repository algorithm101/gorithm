// Copyright (c) 2018 soren yang
//
// Licensed under the MIT License
// you may not use this file except in complicance with the License.
// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0056

import (
	"sort"
)

// Interval is a struct with start and end point
type Interval struct {
	Start int
	End   int
}

// Intervals for sort
type Intervals []Interval

// Len for sort.Interface
func (p Intervals) Len() int {
	return len(p)
}

// Swap for sort.Interface
func (p Intervals) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less for sort.Interface
func (p Intervals) Less(i, j int) bool {
	return p[i].Start < p[j].Start
}

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func merge(intervals []Interval) []Interval {
	// 假设 inervals 是有序的, 这个假设条件不成立...
	ret, w := make([]Interval, len(intervals)), 0

	sort.Sort(Intervals(intervals))

	if 0 == len(intervals) {
		return ret
	}

	last := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= last.End {
			if intervals[i].End > last.End {
				last.End = intervals[i].End
			}
		} else {
			ret[w] = last
			w++
			last = intervals[i]
		}
	}
	ret[w] = last
	w++

	return ret[0:w]
}
