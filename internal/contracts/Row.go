package contracts

// Row represente data that will be persist in database
type Row interface {
	Pk() interface{}
	SetPk(interface{})
	Name() string
	Schema() string
	TableName() string
	Fields() map[string]interface{}
	AddField(key string, value interface{})
	DependencyReference() map[string]Relation
	AddDependency(key string, value Relation)
	HasDependencies() bool
}
