package decorator

import "testing"

func TestDecorator(t *testing.T) {
	tCases := []struct {
		name          string
		addGPU        bool
		addSSD        bool
		expectedPower int
		expectedPrice int
	}{
		{
			name:          "computer without gpu, ssd",
			addGPU:        false,
			addSSD:        false,
			expectedPower: 10,
			expectedPrice: 10,
		},
		{
			name:          "computer with gpu",
			addGPU:        true,
			addSSD:        false,
			expectedPower: 30,
			expectedPrice: 20,
		},
		{
			name:          "computer with ssd",
			addGPU:        false,
			addSSD:        true,
			expectedPower: 30,
			expectedPrice: 30,
		},
		{
			name:          "computer with gpu and ssd",
			addGPU:        true,
			addSSD:        true,
			expectedPower: 50,
			expectedPrice: 40,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			computer := &Computer{}

			if tc.addGPU && !tc.addSSD {
				computerWithGPU := &AddGPU{computer: computer}

				validPrice(computerWithGPU.getPrice(), tc.expectedPrice, tt)
				validPower(computerWithGPU.getPower(), tc.expectedPower, tt)

				return
			}

			if tc.addSSD && !tc.addGPU {
				computerWithSSD := &AddSSD{computer: computer}

				validPrice(computerWithSSD.getPrice(), tc.expectedPrice, tt)
				validPower(computerWithSSD.getPower(), tc.expectedPower, tt)

				return
			}

			if tc.addSSD && tc.addGPU {
				computerWithSSD := &AddSSD{computer: computer}

				computerWithGPU := &AddGPU{computer: computerWithSSD}

				validPrice(computerWithGPU.getPrice(), tc.expectedPrice, tt)
				validPower(computerWithGPU.getPower(), tc.expectedPower, tt)

				return
			}
		})
	}
}

func validPrice(price, expectedPrice int, t *testing.T) {
	if price != expectedPrice {
		t.Errorf("expected price %d but actual %d", expectedPrice, price)
	}
}

func validPower(power, expectedPower int, t *testing.T) {
	if power != expectedPower {
		t.Errorf("expected power %d but actual %d", expectedPower, power)
	}
}
