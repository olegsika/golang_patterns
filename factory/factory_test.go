package factory

import (
	"errors"
	"patterns/tools"
	"reflect"
	"testing"
)

func TestFactory(t *testing.T) {
	tCases := []struct {
		name           string
		company        string
		expectedResult IComputer
		expectedErr    error
	}{
		{
			name:    "Should select Apple Factory",
			company: AppleCompanyName,
			expectedResult: &Apple{
				Computer: Computer{
					memory:    512,
					processor: "Apple M1",
				},
			},
		},
		{
			name:    "Should select Samsung Factory",
			company: SamsungCompanyName,
			expectedResult: &Samsung{Computer: Computer{
				memory:    2048,
				processor: "Samsung A1",
			}},
		},
		{
			name:        "Should select Wrong Factory",
			company:     "random string",
			expectedErr: errors.New("wrong factory passed"),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			res, err := getComputer(tc.company)
			if !tools.IsEqual(err, tc.expectedErr) {
				tt.Errorf("expected err = %v but actual err = %v", tc.expectedErr, err)
				return
			}

			if !reflect.DeepEqual(res, tc.expectedResult) {
				tt.Errorf("expected result = %v but actual result = %v", tc.expectedResult, res)
				return
			}
		})
	}
}
