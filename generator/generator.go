package generator

type Generater interface {
	Generate()
}

type Resolver interface {
	Resolve(value interface{})
}
