package types

import "strconv"

type PVEBool bool

func (b PVEBool) ToAPIRequestParam() string {
	if b {
		return "1"
	}
	return "0"
}

func (b *PVEBool) UnmarshalJSON(slb []byte) error {
	n, err := strconv.ParseBool(string(slb))
	if err != nil {
		return err
	}
	*b = PVEBool(n)
	return nil
}
