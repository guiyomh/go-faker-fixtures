package resolver

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit"
	"github.com/guiyomh/charlatan/contract"
	"github.com/guiyomh/charlatan/generator"
)

var regexFaker, _ = regexp.Compile(`(?i)<(?P<method>[^\(]+)\((?P<args>[^\)]+)?\)>`)

var funcs = map[string]interface{}{
	"Address":              gofakeit.Address,
	"BeerAlcohol":          gofakeit.BeerAlcohol,
	"BeerBlg":              gofakeit.BeerBlg,
	"BeerHop":              gofakeit.BeerHop,
	"BeerIbu":              gofakeit.BeerIbu,
	"BeerMalt":             gofakeit.BeerMalt,
	"BeerName":             gofakeit.BeerName,
	"BeerStyle":            gofakeit.BeerStyle,
	"BeerYeast":            gofakeit.BeerYeast,
	"Bool":                 gofakeit.Bool,
	"BS":                   gofakeit.BS,
	"BuzzWord":             gofakeit.BuzzWord,
	"Categories":           gofakeit.Categories,
	"ChromeUserAgent":      gofakeit.ChromeUserAgent,
	"City":                 gofakeit.City,
	"Color":                gofakeit.Color,
	"Company":              gofakeit.Company,
	"CompanySuffix":        gofakeit.CompanySuffix,
	"Contact":              gofakeit.Contact,
	"Country":              gofakeit.Country,
	"CountryAbr":           gofakeit.CountryAbr,
	"CreditCard":           gofakeit.CreditCard,
	"CreditCardCvv":        gofakeit.CreditCardCvv,
	"CreditCardExp":        gofakeit.CreditCardExp,
	"CreditCardNumber":     gofakeit.CreditCardNumber,
	"CreditCardNumberLuhn": gofakeit.CreditCardNumberLuhn,
	"CreditCardType":       gofakeit.CreditCardType,
	"Currency":             gofakeit.Currency,
	"CurrencyLong":         gofakeit.CurrencyLong,
	"CurrencyShort":        gofakeit.CurrencyShort,
	"Date":                 gofakeit.Date,
	"DateRange":            gofakeit.DateRange,
	"Day":                  gofakeit.Day,
	"DomainName":           gofakeit.DomainName,
	"DomainSuffix":         gofakeit.DomainSuffix,
	"Email":                gofakeit.Email,
	"Extension":            gofakeit.Extension,
	"FirefoxUserAgent":     gofakeit.FirefoxUserAgent,
	"FirstName":            gofakeit.FirstName,
	"Float32":              gofakeit.Float32,
	"Float32Range":         gofakeit.Float32Range,
	"Float64":              gofakeit.Float64,
	"Float64Range":         gofakeit.Float64Range,
	"Gender":               gofakeit.Gender,
	"Generate":             gofakeit.Generate,
	"HackerAbbreviation":   gofakeit.HackerAbbreviation,
	"HackerAdjective":      gofakeit.HackerAdjective,
	"HackerIngverb":        gofakeit.HackerIngverb,
	"HackerNoun":           gofakeit.HackerNoun,
	"HackerPhrase":         gofakeit.HackerPhrase,
	"HackerVerb":           gofakeit.HackerVerb,
	"HexColor":             gofakeit.HexColor,
	"HipsterParagraph":     gofakeit.HipsterParagraph,
	"HipsterSentence":      gofakeit.HipsterSentence,
	"HipsterWord":          gofakeit.HipsterWord,
	"Hour":                 gofakeit.Hour,
	"HTTPMethod":           gofakeit.HTTPMethod,
	"ImageURL":             gofakeit.ImageURL,
	"Int16":                gofakeit.Int16,
	"Int32":                gofakeit.Int32,
	"Int64":                gofakeit.Int64,
	"Int8":                 gofakeit.Int8,
	"IPv4Address":          gofakeit.IPv4Address,
	"IPv6Address":          gofakeit.IPv6Address,
	"Job":                  gofakeit.Job,
	"JobDescriptor":        gofakeit.JobDescriptor,
	"JobLevel":             gofakeit.JobLevel,
	"JobTitle":             gofakeit.JobTitle,
	"LastName":             gofakeit.LastName,
	"Latitude":             gofakeit.Latitude,
	"LatitudeInRange":      gofakeit.LatitudeInRange,
	"Letter":               gofakeit.Letter,
	"Lexify":               gofakeit.Lexify,
	"LogLevel":             gofakeit.LogLevel,
	"Longitude":            gofakeit.Longitude,
	"LongitudeInRange":     gofakeit.LongitudeInRange,
	"MimeType":             gofakeit.MimeType,
	"Minute":               gofakeit.Minute,
	"Month":                gofakeit.Month,
	"Name":                 gofakeit.Name,
	"NamePrefix":           gofakeit.NamePrefix,
	"NameSuffix":           gofakeit.NameSuffix,
	"NanoSecond":           gofakeit.NanoSecond,
	"Number":               gofakeit.Number,
	"Numerify":             gofakeit.Numerify,
	"OperaUserAgent":       gofakeit.OperaUserAgent,
	"Paragraph":            gofakeit.Paragraph,
	"Password":             gofakeit.Password,
	"Person":               gofakeit.Person,
	"Phone":                gofakeit.Phone,
	"PhoneFormatted":       gofakeit.PhoneFormatted,
	"Price":                gofakeit.Price,
	"RandString":           gofakeit.RandString,
	"RGBColor":             gofakeit.RGBColor,
	"SafariUserAgent":      gofakeit.SafariUserAgent,
	"SafeColor":            gofakeit.SafeColor,
	"Second":               gofakeit.Second,
	"Seed":                 gofakeit.Seed,
	"Sentence":             gofakeit.Sentence,
	"ShuffleInts":          gofakeit.ShuffleInts,
	"ShuffleStrings":       gofakeit.ShuffleStrings,
	"SimpleStatusCode":     gofakeit.SimpleStatusCode,
	"SSN":                  gofakeit.SSN,
	"State":                gofakeit.State,
	"StateAbr":             gofakeit.StateAbr,
	"StatusCode":           gofakeit.StatusCode,
	"Street":               gofakeit.Street,
	"StreetName":           gofakeit.StreetName,
	"StreetNumber":         gofakeit.StreetNumber,
	"StreetPrefix":         gofakeit.StreetPrefix,
	"StreetSuffix":         gofakeit.StreetSuffix,
	"Struct":               gofakeit.Struct,
	"TimeZone":             gofakeit.TimeZone,
	"TimeZoneAbv":          gofakeit.TimeZoneAbv,
	"TimeZoneFull":         gofakeit.TimeZoneFull,
	"TimeZoneOffset":       gofakeit.TimeZoneOffset,
	"Uint16":               gofakeit.Uint16,
	"Uint32":               gofakeit.Uint32,
	"Uint64":               gofakeit.Uint64,
	"Uint8":                gofakeit.Uint8,
	"URL":                  gofakeit.URL,
	"UserAgent":            gofakeit.UserAgent,
	"Username":             gofakeit.Username,
	"UUID":                 gofakeit.UUID,
	"WeekDay":              gofakeit.WeekDay,
	"Word":                 gofakeit.Word,
	"Year":                 gofakeit.Year,
	"Zip":                  gofakeit.Zip,
}

func call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("the number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

type FakerValue struct {
}

func (fv *FakerValue) Resolve(
	value contract.Value,
	fixture contract.Fixture,
	fixtureSet generator.ResolvedFixtureSet) interface{} {
	fakesData := regexFaker.FindStringSubmatch(value.String())
	if len(fakesData) <= 0 {
		return value
	}
	method := fakesData[1]
	args := make([]string, 0)
	if fakesData[2] != "" {
		args = strings.Split(fakesData[2], ",")
	}

	result, err := fv.invoke(method, args)
	if err != nil {
		panic(err)
	}
	return fv.convert(result)
}

func (fv *FakerValue) convert(val reflect.Value) interface{} {
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

func (fv *FakerValue) invoke(name string, args []string) (reflect.Value, error) {

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
