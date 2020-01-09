# Go Avro Serializer

A sample code in [Go][1] to serialize an object in [Avro][2] using [goavro][3].

## Kafka

If you want to store your Avro in Kafka, you should serialize your object using
[Single-object encoding][4] with a [canonical schema][5] (in file header).

Initially, I was using [gogen-avro][6] to serialize objects, but the serialization is not
an [Avro Object Container File (OCF)][8]. (Although its documentation has
information about [OCF][7] and an [opened issue][9] about this problem too).

At the end of the day, I decided to select [goavro][3] to serialize the object in Avro OCF.

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

