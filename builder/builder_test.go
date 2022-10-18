package builder

import (
	"errors"
	"patterns/tools"
	"reflect"
	"testing"
)

func TestBuilder(t *testing.T) {
	tCases := []struct {
		name           string
		builderName    string
		expectedResult Computer
		expectedErr    error
	}{
		{
			name:        "Should select Apple builder",
			builderName: AppleBuilderName,
			expectedResult: Computer{
				RAM: 2048,
				GPU: "nvidia GeForce",
				SSD: 2048,
				CPU: 2.2,
			},
		},
		{
			name:        "Should select Samsung builder",
			builderName: SamsungBuilderName,
			expectedResult: Computer{
				RAM: 1024,
				GPU: "AMD radeon",
				SSD: 1024,
				CPU: 1.1,
			},
		},
		{
			name:        "Should select Random builder",
			builderName: "test test test",
			expectedErr: errors.New("wrong builder passed"),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			builder, err := getBuilder(tc.builderName)
			if !tools.IsEqual(err, tc.expectedErr) {
				tt.Errorf("expected err = %v but actual err = %v", tc.expectedErr, err)
				return
			}

			if err != nil {
				return
			}

			director := newDirector(builder)
			res := director.buildComputer()

			if !reflect.DeepEqual(tc.expectedResult, res) {
				tt.Errorf("expected result = %v but actual result = %v", tc.expectedResult, res)
				return
			}
		})
	}
}
