package models

import (
	"fmt"

	"github.com/guiyomh/charlatan/contract"
)

type Data struct {
	value interface{}
}

func NewData(value interface{}) contract.Value {
	return &Data{
		value: value,
	}
}

func (d *Data) Value() interface{} {
	return d.value
}

func (d *Data) String() string {
	return fmt.Sprintf("%s", d)
}
