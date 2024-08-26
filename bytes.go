package pflag

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

type bytesHexValue []byte

func (bytesHex bytesHexValue) String() string {
	return fmt.Sprintf("%X", []byte(bytesHex))
}

func (bytesHex *bytesHexValue) Set(value string) error {
	bin, err := hex.DecodeString(strings.TrimSpace(value))

	if err != nil {
		return err
	}

	*bytesHex = bin

	return nil
}

func (*bytesHexValue) Type() string {
	return "bytesHex"
}

func newBytesHexValue(val []byte, p *[]byte) *bytesHexValue {
	*p = val
	return (*bytesHexValue)(p)
}

func bytesHexConv(sval string) (interface{}, error) {
	bin, err := hex.DecodeString(sval)

	if err == nil {
		return bin, nil
	}

	return nil, fmt.Errorf("invalid string beging converted to Bytes: %s %s", sval, err)
}

func (f *FlagSet) GetBytesHex(name string) ([]byte, error) {
	val, err := f.getFlagType(name, "bytesHex", bytesHexConv)

	if err != nil {
		return []byte{}, err
	}

	return val.([]byte), nil
}

func (f *FlagSet) BytesHexVar(p *[]byte, name string, value []byte, usage string) {
	f.VarP(newBytesHexValue(value, p), name, "", usage)
}

func (f *FlagSet) BytesHexVarP(p *[]byte, name, shorthand string, value []byte, usage string) {
	f.VarP(newBytesHexValue(value, p), name, shorthand, usage)
}

func BytesHexVar(p *[]byte, name string, value []byte, usage string) {
	CommandLine.VarP(newBytesHexValue(value, p), name, "", usage)
}

func BytesHexVarP(p *[]byte, name, shorthand string, value []byte, usage string) {
	CommandLine.VarP(newBytesHexValue(value, p), name, shorthand, usage)
}

func (f *FlagSet) BytesHex(name string, value []byte, usage string) *[]byte {
	p := new([]byte)
	f.BytesHexVarP(p, name, "", value, usage)
	return p
}

func (f *FlagSet) BytesHexP(name, shorthand string, value []byte, usage string) *[]byte {
	p := new([]byte)
	f.BytesHexVarP(p, name, shorthand, value, usage)
	return p
}

func BytesHex(name string, value []byte, usage string) *[]byte {
	return CommandLine.BytesHexP(name, "", value, usage)
}

// BytesHexP is like BytesHex, but accepts a shorthand letter that can be used after a single dash.
func BytesHexP(name, shorthand string, value []byte, usage string) *[]byte {
	return CommandLine.BytesHexP(name, shorthand, value, usage)
}

type bytesBae64Value []byte 

func (bytesBase64 bytesBae64Value) String() string {
	return base64.StdEncoding.EncodeToString([]byte(bytesBase64))
}

func (bytesBse64 *bytesBae64Value) Set(value string) error {
	bin, err := base64.StdEncoding.DecodeString(strings.TrimSpace(value))

	if err != nil {
		return err
	}

	*bytesBse64 = bin

	return nil
}



