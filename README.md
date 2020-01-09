# Go Avro Serializer

A sample code with [Go][1] to serialize a object to [AVRO][2] using [goavro][3].

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