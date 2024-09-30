package prefix_sum

import (
	"log"
	"testing"
)

func TestPrefixSum(t *testing.T) {

	inputs := []struct {
		Description    string
		Values         []int
		Start          int
		End            int
		ExpectedOutput int
	}{
		{
			Description:    "happy path",
			Values:         []int{1, 2, 3, 4, 5, 6},
			Start:          1,
			End:            3,
			ExpectedOutput: 9,
		},
		{
			Description:    "empty input",
			Values:         []int{},
			Start:          1,
			End:            3,
			ExpectedOutput: 0,
		},
		{
			Description:    "end out of range",
			Values:         []int{1, 2, 3, 4, 5, 6},
			Start:          1,
			End:            8,
			ExpectedOutput: 0,
		},
		{
			Description:    "start greater than end",
			Values:         []int{1, 2, 3, 4, 5, 6},
			Start:          6,
			End:            3,
			ExpectedOutput: 0,
		},
		{
			Description:    "start equal to end",
			Values:         []int{1, 2, 3, 4, 5, 6},
			Start:          3,
			End:            3,
			ExpectedOutput: 4,
		},
		{
			Description:    "test case 5",
			Values:         []int{-2, 0, 3, -5, 2, -1},
			Start:          0,
			End:            2,
			ExpectedOutput: 1,
		},
		{
			Description:    "test case 6",
			Values:         []int{-2, 0, 3, -5, 2, -1},
			Start:          2,
			End:            5,
			ExpectedOutput: -1,
		},
		{
			Description:    "test case 7",
			Values:         []int{-2, 0, 3, -5, 2, -1},
			Start:          0,
			End:            5,
			ExpectedOutput: -3,
		},
	}

	for _, input := range inputs {
		log.Println(input.Description)

		output := prefixSum(input.Values, input.Start, input.End)

		if input.ExpectedOutput != output {
			t.Errorf("expected output %d does not match output: %d", input.ExpectedOutput, output)
		}
	}
}

func TestSumRange(t *testing.T) {

	inputs := []struct {
		Description    string
		Values         []int
		Start          int
		End            int
		ExpectedOutput int
	}{
		{
			Description:    "happy path",
			Values:         []int{1, 2, 3, 4, 5, 6},
			Start:          1,
			End:            3,
			ExpectedOutput: 9,
		},
		{
			Description:    "test case 1",
			Values:         []int{-2, 0, 3, -5, 2, -1},
			Start:          0,
			End:            2,
			ExpectedOutput: 1,
		},
		{
			Description:    "test case 2",
			Values:         []int{-2, 0, 3, -5, 2, -1},
			Start:          2,
			End:            5,
			ExpectedOutput: -1,
		},
		{
			Description:    "test case 3",
			Values:         []int{-2, 0, 3, -5, 2, -1},
			Start:          0,
			End:            5,
			ExpectedOutput: -3,
		},
	}

	for _, input := range inputs {
		log.Println(input.Description)

		prefixSumTheirs := Constructor(input.Values)
		output := prefixSumTheirs.SumRange(input.Start, input.End)

		if input.ExpectedOutput != output {
			t.Errorf("expected output %d does not match output: %d", input.ExpectedOutput, output)
		}
	}
}
