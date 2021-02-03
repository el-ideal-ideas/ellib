package elencode

import (
	"encoding/base32"
	"encoding/base64"
)

// alias.
var Base32EncodeToString = base32.StdEncoding.EncodeToString
var Base32DecodeString = base32.StdEncoding.DecodeString
var Base32Encode = base32.StdEncoding.Encode
var Base32Decode = base32.StdEncoding.Decode
var Base64EncodeToString = base64.StdEncoding.EncodeToString
var Base64DecodeString = base64.StdEncoding.DecodeString
var Base64Encode = base64.StdEncoding.Encode
var Base64Decode = base64.StdEncoding.Decode
