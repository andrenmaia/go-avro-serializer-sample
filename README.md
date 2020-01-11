# Go Avro Serializer

A sample code in [Go][1] to serialize an object in [Avro][2] using [goavro][3].

## Kafka

If you want to store your Avro in Kafka, you should serialize your object using
with schema registry header information.

Initially, I was using [gogen-avro][4] to serialize objects, but the
serialization do not include schema registry headers information.

At the end of the day, I decided to select [goavro][3] to serialize the object
and set header manually.

## Dependencies

```
go get "github.com/linkedin/goavro"
```

## Test

```
# Build go
cd go-serializer 
go build src/serializer.go;

# Generate file
./serializer
```

```
# Print json from avro file
cd ../java-serializer
./gradlew run
```

The output will be `{"code": "code code code 1", "amount": 1}`


[1]: https://golang.org/
[2]: https://avro.apache.org/
[3]: https://github.com/linkedin/goavro
[4]: https://github.com/actgardner/gogen-avro

