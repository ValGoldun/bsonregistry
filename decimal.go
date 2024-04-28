package bsonregistry

import (
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"reflect"
)

var (
	tDecimal = reflect.TypeOf(decimal.Decimal{})
)

func decimalEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tDecimal {
		return bsoncodec.ValueEncoderError{Name: "decimalEncodeValue", Types: []reflect.Type{tDecimal}, Received: val}
	}
	return vw.WriteString(val.Interface().(decimal.Decimal).String())
}

func decimalDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tDecimal {
		return bsoncodec.ValueDecoderError{Name: "decimalDecodeValue", Types: []reflect.Type{tDecimal}, Received: val}
	}

	data, err := vr.ReadString()
	if err != nil {
		return err
	}

	value, err := decimal.NewFromString(data)
	if err != nil {
		return err
	}

	val.Set(reflect.ValueOf(value))

	return nil
}
