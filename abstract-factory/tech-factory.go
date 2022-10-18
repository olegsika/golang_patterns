package abstractfactory

import "fmt"

func GetTechFactory(company string) (ITechFactory, error) {
	switch company {
	case AppleCompanyName:
		return &Apple{}, nil
	case SamsungCompanyName:
		return &Samsung{}, nil
	}

	return nil, fmt.Errorf("wrong factory passed")
}
