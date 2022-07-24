package helpers

import (
	"encoding/json"
	"fmt"
)

type IntBool bool

func (i *IntBool) UnmarshalJSON(b []byte) error {
	*i = IntBool(string(b) == "1")
	return nil
}

func (i IntBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(IntBool(i))
}

func (i IntBool) Int() int {
	if i {
		return 1
	}
	return 0
}

func (i IntBool) IntAsString() string {
	return fmt.Sprintf("%d", i.Int())
}

func (i IntBool) Bool() bool {
	return bool(i)
}
