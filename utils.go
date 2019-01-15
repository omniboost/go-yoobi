package yoobi

import (
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	null "gopkg.in/guregu/null.v3"

	"github.com/gorilla/schema"
)

type SchemaMarshaler interface {
	MarshalSchema() string
}

type ToURLValues interface {
	ToURLValues() (url.Values, error)
}

func AddQueryParamsToRequest(requestParams interface{}, req *http.Request, skipEmpty bool) error {
	var err error
	params := url.Values{}

	to, ok := requestParams.(ToURLValues)
	if ok == true {
		params, err = to.ToURLValues()
		if err != nil {
			return err
		}
	} else {
		encoder := newSchemaEncoder()
		err := encoder.Encode(requestParams, params)
		if err != nil {
			return err
		}
	}

	query := req.URL.Query()
	for k, vals := range params {
		for _, v := range vals {
			if skipEmpty && v == "" {
				continue
			}

			if skipEmpty && v == "0" {
				continue
			}

			query.Add(k, v)
		}
	}

	req.URL.RawQuery = query.Encode()
	return nil
}

func newSchemaEncoder() *schema.Encoder {
	encoder := schema.NewEncoder()

	// // register custom encoders
	// encodeSchemaMarshaler := func(v reflect.Value) string {
	// 	marshaler, ok := v.Interface().(SchemaMarshaler)
	// 	if ok == true {
	// 		return marshaler.MarshalSchema()
	// 	}

	// 	stringer, ok := v.Interface().(fmt.Stringer)
	// 	if ok == true {
	// 		return stringer.String()
	// 	}

	// 	return ""
	// }

	encodeNullFloat := func(v reflect.Value) string {
		nullFloat, _ := v.Interface().(null.Float)
		if nullFloat.IsZero() {
			return ""
		}
		return strconv.FormatFloat(nullFloat.Float64, 'f', 6, 64)
	}

	encodeNullBool := func(v reflect.Value) string {
		nullBool, _ := v.Interface().(null.Bool)
		if nullBool.IsZero() {
			return ""
		}
		return strconv.FormatBool(nullBool.Bool)
	}

	// encoder.RegisterEncoder(Date{}, encodeSchemaMarshaler)
	encoder.RegisterEncoder(null.Float{}, encodeNullFloat)
	encoder.RegisterEncoder(null.Bool{}, encodeNullBool)
	return encoder
}
