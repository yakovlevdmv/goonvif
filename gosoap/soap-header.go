package gosoap

type headerHandler func(value string) interface{}

//Xlmns XML Scheam
var supportedHeaders = map[string]headerHandler{
	"to":     toHandler,
	"action": actionHadler,
}

func getHeaderSectionfrom(k, v string) interface{} {

	if f, found := supportedHeaders[k]; found {
		return f(v)
	}
	return nil
}

func toHandler(value string) interface{} {
	return NewTo(value)
}

func actionHadler(value string) interface{} {
	return NewAction(value)
}
