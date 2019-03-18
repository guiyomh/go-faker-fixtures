package contracts

type Faker interface {
	Generate(data string) interface{}
}
