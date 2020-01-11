package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"

	"github.com/linkedin/goavro"
)

// Avro Schema to serialize
const schemaSpecification = `
{
	"namespace":"main", 
	"type":"record", 
	"name":"ReceivableSettled", 
	"fields":[ 
		{ 
			"name":"code", 
			"type":"string"
		},
		{
			"name":"amount",
			"type":"int"
		}
	]
}`

// MagicByte number for schema registry avro header
const MagicByte = 0x0

// Header of avro header
type Header struct {
	MagicByte byte
	ID        int32
}

// Receivable object to serialize in avro
type Receivable struct {
	Code   string
	Amount int
}

// ToMap a Receivable
func (r *Receivable) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":   r.Code,
		"amount": r.Amount,
	}
}

func main() {
	codec, _ := goavro.NewCodec(schemaSpecification)
	data := Receivable{
		Code:   "code code code " + strconv.Itoa(1),
		Amount: 1,
	}

	binBuffer := serialize(codec, data)

	saveToFile(binBuffer.Bytes())
}

func printBanner(msg string) {
	fmt.Println("===========================================")
	fmt.Println(msg)
	fmt.Println("===========================================")
}

func serialize(codec *goavro.Codec, data Receivable) bytes.Buffer {
	printBanner("Serialize object")
	binaryFromNative, err := codec.BinaryFromNative(nil, data.ToMap())
	if err != nil {
		panic(err)
	}

	var binBuffer bytes.Buffer

	header := Header{
		MagicByte: 0x0,
		ID:        931,
	}

	_ = binary.Write(&binBuffer, binary.BigEndian, header)
	_ = binary.Write(&binBuffer, binary.BigEndian, binaryFromNative)

	fmt.Println("Byte array print:")
	fmt.Println(binBuffer.Bytes())

	return binBuffer
}

func saveToFile(binaryData []byte) {
	printBanner("Save byte array in file")

	path := "event-from-byte-array.avro"
	f, _ := os.Create(path)
	f.Write(binaryData)

	fmt.Print("File \"event-from-byte-array.avro\" created\n\n")
}
