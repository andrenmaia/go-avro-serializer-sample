# Go Avro Serializer

A sample code in [Go][1] to serialize an object in [Avro][2] using [goavro][3].

## Kafka

If you want to store your Avro in Kafka, you should serialize your object using
with schema registry header information.

Initially, I was using [gogen-avro][6] to serialize objects, but the
serialization do not include schema registry headers information.

At the end of the day, I decided to select [goavro][3] to serialize the object
and set header manually.

## Dependencies

```
go get "github.com/linkedin/goavro"
```

## Test

``` 
go build src/main.go;
```

```
# Generate file
./main

# Print json from avro file
java -jar /opt/avro-tools/avro-tools-1.9.1.jar tojson event-from-byte-array.avro| jq
```

>  Using Avro tools from https://avro.apache.org/releases.html#2+September+2019%3A+Avro+1.9.1+Released



[1]: https://golang.org/
[2]: https://avro.apache.org/
[3]: https://github.com/linkedin/goavro
[4]: https://avro.apache.org/docs/1.8.2/spec.html#single_object_encoding
[5]: https://avro.apache.org/docs/1.8.2/spec.html#Transforming+into+Parsing+Canonical+Form
[6]: https://github.com/actgardner/gogen-avro
[7]: https://github.com/actgardner/gogen-avro#working-with-object-container-files-ocf
[8]: https://avro.apache.org/docs/1.8.1/spec.html#Object+Container+Files
[9]: https://github.com/actgardner/gogen-avro/issues/111

