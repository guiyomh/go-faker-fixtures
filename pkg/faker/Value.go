package faker

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/guiyomh/charlatan/pkg/faker/contracts"
)

type value struct {
}

var regexFaker, _ = regexp.Compile(`(?i)<(?P<method>[^\(]+)\((?P<args>[^\)]+)?\)>`)

// NewValue is a factory to create a generator of value
func NewValue() contracts.Faker {
	return &value{}
}

// Generate is a method that use faker to generate ramdom data
func (g *value) Generate(data string) interface{} {
	fakesData := regexFaker.FindStringSubmatch(data)
	if len(fakesData) <= 0 {
		return data
	}
	method := fakesData[1]
	args := make([]string, 0)
	if fakesData[2] != "" {
		args = strings.Split(fakesData[2], ",")
	}
	result, err := g.invoke(method, args)
	if err != nil {
		panic(err)
	}
	return g.convert(result)
}

func (g *value) convert(val reflect.Value) interface{} {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.String:
		return val.String()
	case reflect.Bool:
		return val.Bool()
	case reflect.Struct:
		return val.Interface()
	default:
		return "nothing"
	}
}

func (g *value) invoke(name string, args []string) (reflect.Value, error) {

	// method := reflect.ValueOf(gofakeit).MethodByName(name)
	method := reflect.ValueOf(funcs[name])
	methodType := method.Type()
	numIn := methodType.NumIn()
	if numIn > len(args) {
		return reflect.ValueOf(nil), fmt.Errorf("Method %s must have minimum %d params. Have %d", name, numIn, len(args))
	}
	if numIn != len(args) && !methodType.IsVariadic() {
		return reflect.ValueOf(nil), fmt.Errorf("Method %s must have %d params. Have %d", name, numIn, len(args))
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		var inType reflect.Type
		if methodType.IsVariadic() && i >= numIn-1 {
			inType = methodType.In(numIn - 1).Elem()
		} else {
			inType = methodType.In(i)
		}
		argValue := reflect.ValueOf(argConverter(args[i], inType.String()))
		if !argValue.IsValid() {
			return reflect.ValueOf(nil), fmt.Errorf("0.Method %s. Param[%d] must be %s. Have %s", name, i, inType, argValue.String())
		}
		argType := argValue.Type()
		if argType.ConvertibleTo(inType) {
			in[i] = argValue.Convert(inType)
		} else {
			return reflect.ValueOf(nil), fmt.Errorf("1.Method %s. Param[%d] must be %s. Have %s", name, i, inType, argType)
		}
	}
	result := method.Call(in)
	return result[0], nil
}

//https://gist.github.com/tkrajina/880eb4b9a10aee28707e2aa764257503
func argConverter(value string, typeof string) interface{} {
	switch typeof {
	case "bool":
		castValue, _ := strconv.ParseBool(value)
		return castValue
	case "int":
		castValue, _ := strconv.Atoi(value)
		return castValue
	default:
		return value
	}
}
