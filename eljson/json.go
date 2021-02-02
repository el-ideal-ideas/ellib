// This package is a alias of json-iterator/go

package eljson

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var Get = json.Get
var Marshal = json.Marshal
var MarshalIndent = json.MarshalIndent
var BorrowIterator = json.BorrowIterator
var BorrowStream = json.BorrowStream
var DecoderOf = json.DecoderOf
var EncoderOf = json.EncoderOf
var MarshalToString = json.MarshalToString
var NewDecoder = json.NewDecoder
var NewEncoder = json.NewEncoder
var RegisterExtension = json.RegisterExtension
var Valid = json.Valid
var UnmarshalFromString = json.UnmarshalFromString
var ReturnStream = json.ReturnStream
var Unmarshal = json.Unmarshal
var ReturnIterator = json.ReturnIterator
