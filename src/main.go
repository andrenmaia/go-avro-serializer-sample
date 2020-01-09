package main

import (
	"github.com/linkedin/goavro"
	"strconv"
	"bytes"
	"fmt"
	"os"
)

func main() {
	codec, err := goavro.NewCodec(`
	{ 
	  "namespace":"main",
	  "type":"record",
	  "name":"ReceivableSettled",
	  "fields":[ 
		{ 
		  "name":"code",
		  "type":"string",
		  "doc":"Código do recebível"
		},
		{ 	
		  "name":"amount",
		  "type":"int",
		  "doc":"Valor total da liquidação. Padrão inteiro: 1234567 para R$ 123.145,67."
		}
	  ]
	}    
	`)
	if err != nil {
		panic(err)
	}

	receivables := buildObjects(3)

	fmt.Println("===========================================")
	fmt.Println("Create Byte Array")
	fmt.Println("===========================================")
	var buf bytes.Buffer
	ocfw, err := goavro.NewOCFWriter(goavro.OCFConfig{
		W:     &buf,
		Codec: codec,
	})
	if err != nil {
		panic(err)
	}
	if err = ocfw.Append(receivables); err != nil {
		panic(err)
	}

	bb := buf.Bytes()
	fmt.Println("Byte array print:")
	fmt.Println(bb)

	fmt.Println("\n\n===========================================")
	fmt.Println("Save byte array in file")
	fmt.Println("===========================================")
	path := "event-from-byte-array.avro"
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	f.Write(bb)

	fmt.Print("File \"event-from-byte-array.avro\" created\n\n")
}

type Receivable struct {
	Code string
	Amount int
}

func (r *Receivable) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code": r.Code, 
		"amount": r.Amount,
	}
}


func buildObjects(quantity int) []map[string]interface{} {
	var values []map[string]interface{}
	for i := 1; i < quantity+1; i++ {
		r := Receivable {
			Code: "code " + strconv.Itoa(i),
			Amount: i,
		}

		values = append(values, r.ToMap())
	}

	return values
}