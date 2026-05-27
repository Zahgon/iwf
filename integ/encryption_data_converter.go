package integ

import (
	commonpb "go.temporal.io/api/common/v1"

	"go.temporal.io/sdk/converter"
)

// NOTE: the code is from https://github.com/temporalio/samples-go/blob/main/encryption/data_converter.go

const (
	// MetadataEncodingEncrypted is "binary/encrypted"
	MetadataEncodingEncrypted = "binary/encrypted"

	// MetadataEncryptionKeyID is "encryption-key-id"
	MetadataEncryptionKeyID = "encryption-key-id"
)

type DataConverter struct {
	// Until EncodingDataConverter supports workflow.ContextAware we'll store parent here.
	parent converter.DataConverter
	converter.DataConverter
	options DataConverterOptions
}

type DataConverterOptions struct {
	KeyID string
	// Enable ZLib compression before encryption.
	Compress bool
}

// Codec implements PayloadCodec using AES Crypt.
type Codec struct {
	KeyID string
}

func (e *Codec) getKey(keyID string) (key []byte) {
	_ = "STUB: not implemented"
	// Key must be fetched from secure storage in production (such as a KMS).
	// For testing here we just hard code a key.
	return nil
}

var encryptionDataConverter = NewEncryptionDataConverter(converter.GetDefaultDataConverter(), DataConverterOptions{})

// NewEncryptionDataConverter creates a new instance of EncryptionDataConverter wrapping a DataConverter
func NewEncryptionDataConverter(dataConverter converter.DataConverter, options DataConverterOptions) *DataConverter {
	_ = "STUB: not implemented"
	return nil
}

// Enable compression if requested.
// Note that this must be done before encryption to provide any value. Encrypted data should by design not compress very well.
// This means the compression codec must come after the encryption codec here as codecs are applied last -> first.

// Encode implements converter.PayloadCodec.Encode.
func (e *Codec) Encode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decode implements converter.PayloadCodec.Decode.
func (e *Codec) Decode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only if it's encrypted
