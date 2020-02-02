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

package p0401

import (
	"fmt"
)

func _readBinaryWatch01(num int) []string { //nolint
	hours := map[int][]string{
		0: []string{"0"},
		1: []string{"1", "2", "4", "8"},
		2: []string{"3", "5", "6", "9", "10"},
		3: []string{"7", "11"},
	}
	minutes := map[int][]string{
		0: []string{"00"},
		1: []string{"01", "02", "04", "08", "16", "32"},
		2: []string{
			"03", "05", "09", "17", "33",
			"06", "10", "18", "34",
			"12", "20", "36",
			"24", "40",
			"48",
		},
		3: []string{
			"07", "11", "19", "35",
			"13", "21", "37",
			"25", "41",
			"49",
			"14", "22", "38",
			"26", "42",
			"50",
			"28", "44",
			"52",
			"56",
		},
		4: []string{
			"15", "23", "27", "29", "30", "39", "43", "45", "46", "51", "53", "54", "57", "58",
		},
		5: []string{
			"31",
			"47",
			"55",
			"59",
		},
	}

	ret := make([]string, 0)

	for hoursUp := 0; hoursUp < 4 && hoursUp <= num; hoursUp++ {
		minutesUp := num - hoursUp
		if minutesUp <= 5 {
			minutesString := minutes[minutesUp]
			hoursString := hours[hoursUp]
			for _, hourString := range hoursString {
				for _, minuteString := range minutesString {
					ret = append(ret, hourString+":"+minuteString)
				}
			}
		}
	}

	return ret
}

func _readBinaryWatch02(num int) []string {
	countBinary1 := map[int]int{
		0:  0,
		1:  1,
		2:  1,
		3:  2,
		4:  1,
		5:  2,
		6:  2,
		7:  3,
		8:  1,
		9:  2,
		10: 2,
		11: 3,
		12: 2,
		13: 3,
		14: 3,
		15: 4,
		16: 1,
		17: 2,
		18: 2,
		19: 3,
		20: 2,
		21: 3,
		22: 3,
		23: 4,
		24: 2,
		25: 3,
		26: 3,
		27: 4,
		28: 3,
		29: 4,
		30: 4,
		31: 5,
		32: 1,
		33: 2,
		34: 2,
		35: 3,
		36: 2,
		37: 3,
		38: 3,
		39: 4,
		40: 2,
		41: 3,
		42: 3,
		43: 4,
		44: 3,
		45: 4,
		46: 4,
		47: 5,
		48: 2,
		49: 3,
		50: 3,
		51: 4,
		52: 3,
		53: 4,
		54: 4,
		55: 5,
		56: 3,
		57: 4,
		58: 4,
		59: 5,
	}

	ret := make([]string, 0)
	for hour := 0; hour < 12; hour++ {
		for minute := 0; minute < 60; minute++ {
			if num == countBinary1[hour]+countBinary1[minute] {
				if minute < 10 {
					ret = append(ret, fmt.Sprintf("%d:0%d", hour, minute))
				} else {
					ret = append(ret, fmt.Sprintf("%d:%d", hour, minute))
				}
			}
		}
	}
	return ret
}

func readBinaryWatch(num int) []string {
	return _readBinaryWatch02(num)
}
