package abstractfactory

import (
	"errors"
	"patterns/tools"
	"testing"
)

func TestGetTechFactory(t *testing.T) {
	tCases := []struct {
		name           string
		company        string
		expectedResult ITechFactory
		expectedErr    error
	}{
		{
			name:           "Should select Apple Factory",
			company:        AppleCompanyName,
			expectedResult: &Apple{},
		},
		{
			name:           "Should select Samsung Factory",
			company:        SamsungCompanyName,
			expectedResult: &Samsung{},
		},
		{
			name:        "Should select Wrong Factory",
			company:     "random string",
			expectedErr: errors.New("wrong factory passed"),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			res, err := GetTechFactory(tc.company)

			if !tools.IsEqual(err, tc.expectedErr) {
				tt.Errorf("expected err = %v but actual err = %v", tc.expectedErr, err)
				return
			}

			if res != tc.expectedResult {
				tt.Errorf("expected result = %v but actual result = %v", tc.expectedResult, res)
				return
			}
		})
	}
}
