package bsonregistry

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
)

func Registry() *bsoncodec.Registry {
	registry := bson.NewRegistry()

	registry.RegisterTypeEncoder(tUUID, bsoncodec.ValueEncoderFunc(uuidEncodeValue))
	registry.RegisterTypeDecoder(tUUID, bsoncodec.ValueDecoderFunc(uuidDecodeValue))

	registry.RegisterTypeEncoder(tDecimal, bsoncodec.ValueEncoderFunc(decimalEncodeValue))
	registry.RegisterTypeDecoder(tDecimal, bsoncodec.ValueDecoderFunc(decimalDecodeValue))

	return registry
}
