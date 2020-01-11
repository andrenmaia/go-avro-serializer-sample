package app;

import io.confluent.kafka.schemaregistry.client.MockSchemaRegistryClient;
import io.confluent.kafka.schemaregistry.client.SchemaRegistryClient;
import io.confluent.kafka.schemaregistry.client.rest.exceptions.RestClientException;
import io.confluent.kafka.serializers.KafkaAvroDeserializer;
import org.apache.avro.Schema;
import org.apache.commons.compress.utils.IOUtils;

import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;

public class Deserializer {

    public static final String SCHEMA = "{\"namespace\":\"main\", \"type\":\"record\", \"name\":\"ReceivableSettled\", \"fields\":[ { \"name\":\"code\", \"type\":\"string\"},{\"name\":\"amount\",\"type\":\"int\"}]}";
    public static final String SUBJECT = "event.schedule_receivable.receivable_settled.created-value";

    public static void main(String[] args) throws IOException, RestClientException {

        deserialize();
    }

    private static void deserialize() throws IOException, RestClientException {
        final SchemaRegistryClient client = new MockSchemaRegistryClient();
        final Schema schema = new Schema.Parser().parse(SCHEMA);
        client.register(SUBJECT, schema, 10, 931);

        final String path = "../go-serializer/event-from-byte-array.avro";

        try (FileInputStream file = new FileInputStream(new File(path))) {
            final byte[] bytes = IOUtils.toByteArray(file);
            final KafkaAvroDeserializer kafkaAvroDeserializer = new KafkaAvroDeserializer(client);
            final Object result = kafkaAvroDeserializer.deserialize(SUBJECT, bytes, schema);
            System.out.println(result.toString());
        }
    }


}
