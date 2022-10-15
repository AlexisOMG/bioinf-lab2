package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	cases := []struct {
		name                 string
		first                string
		second               string
		gap                  int
		expectedScore        int
		expectedFirstString  string
		expectedSecondString string
	}{
		{
			name:                 "lesson_example",
			first:                "AATCGA",
			second:               "AACGT",
			gap:                  -5,
			expectedScore:        11,
			expectedFirstString:  "AATCGA",
			expectedSecondString: "AA-CGT",
		},
		{
			name:                 "only_substitution",
			first:                "AATCGA",
			second:               "TTTTTT",
			gap:                  -5,
			expectedScore:        -15,
			expectedFirstString:  "AATCGA",
			expectedSecondString: "TTTTTT",
		},
		{
			name:                 "single_match_several_deleting",
			first:                "TTTTTT",
			second:               "T",
			gap:                  -5,
			expectedScore:        -20,
			expectedFirstString:  "TTTTTT",
			expectedSecondString: "-----T",
		},
		{
			name:                 "empty_second_string",
			first:                "TTTTTT",
			second:               "",
			gap:                  -5,
			expectedScore:        -30,
			expectedFirstString:  "TTTTTT",
			expectedSecondString: "------",
		},
		{
			name:                 "empty_first_string",
			first:                "",
			second:               "TTTTTT",
			gap:                  -5,
			expectedScore:        -30,
			expectedFirstString:  "------",
			expectedSecondString: "TTTTTT",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualScore := getScore(c.first, c.second, 5, -4, c.gap)
			assert.Equal(t, c.expectedScore, actualScore)
		})
	}
}
