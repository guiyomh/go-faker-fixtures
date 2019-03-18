package middleware

import (
	"fmt"

	"github.com/guiyomh/charlatan/pkg/faker/contracts"
)

func FakerMiddleware(faker contracts.Faker) func(interface{}) interface{} {
	return func(value interface{}) interface{} {
		typeof := fmt.Sprintf("%T", value)
		if typeof == "string" {
			value = faker.Generate(value.(string))
		}
		return value
	}
}
