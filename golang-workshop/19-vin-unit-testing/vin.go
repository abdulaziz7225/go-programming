package vin

type VIN struct {
	value  string
	region string
}

func NewVINFromString(vin string) (*VIN, error) {
	if len(vin) != 17 {
		return nil, ErrInvalidVINLength
	}

	return &VIN{
		value:  vin,
		region: string(vin[0]),
	}, nil
}
