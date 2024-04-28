package bsonregistry

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"reflect"
)

var (
	tUUID = reflect.TypeOf(uuid.UUID{})
)

func uuidEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tUUID {
		return bsoncodec.ValueEncoderError{Name: "uuidEncodeValue", Types: []reflect.Type{tUUID}, Received: val}
	}
	return vw.WriteString(val.Interface().(uuid.UUID).String())
}

func uuidDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tUUID {
		return bsoncodec.ValueDecoderError{Name: "uuidDecodeValue", Types: []reflect.Type{tUUID}, Received: val}
	}

	data, err := vr.ReadString()
	if err != nil {
		return err
	}

	value, err := uuid.Parse(data)
	if err != nil {
		return err
	}

	val.Set(reflect.ValueOf(value))

	return nil
}
