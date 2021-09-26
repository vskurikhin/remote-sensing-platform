package su.svn.invalidator.config;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class DebeziumConnectorConfiguration {
    /**
     * Database details.
     */
    @Value("${abdb.datasource.host}")
    private String dbHost;

    @Value("${abdb.datasource.databasename}")
    private String dbName;

    @Value("${abdb.datasource.port}")
    private String dbPort;

    @Value("${abdb.datasource.username}")
    private String dbUserName;

    @Value("${abdb.datasource.password}")
    private String dbPassword;

    @Value("${abdb.offset.storage.file.filename:#{null}}")
    private String offsetStorageFileFilename;

    @Value("${debezium.kafka.bootstrap-servers:#{null}}")
    private String bootstrapServers;

    @Value("${debezium.kafka.client-id:#{null}}")
    private String clientId;

    @Value("${debezium.kafka.security-protocol:#{null}}")
    private String securityProtocol;

    @Value("${debezium.kafka.security-inter-broker-protocol:#{null}}")
    private String securityInterBrokerProtocol;

    @Value("${debezium.kafka.ssl-endpoint-identification-algorithm:#{null}}")
    private String sslEndpointIdentificationAlgorithm;

    @Value("${debezium.kafka.key-store-location:#{null}}")
    private String keyStoreLocation;

    @Value("${debezium.kafka.key-store-password:#{null}}")
    private String keyStorePassword;

    @Value("${debezium.kafka.trust-store-location:#{null}}")
    private String trustStoreLocation;

    @Value("${debezium.kafka.trust-store-password:#{null}}")
    private String trustStorePassword;

    @Value("${debezium.kafka.config.storage.topic:#{null}}")
    private String configStorageTopic;

    @Value("${debezium.kafka.config.storage.partitions:1}")
    private int configStoragePartitions;

    @Value("${debezium.kafka.offset.flush-interval-ms:1000}")
    private int offsetFlushIntervalMs;

    @Value("${debezium.kafka.config.storage.replication-factor:1}")
    private int configStorageReplicationFactor;

    @Value("${debezium.kafka.offset.storage.topic:#{null}}")
    private String offsetStorageTopic;

    @Value("${debezium.kafka.offset.storage.partitions:1}")
    private int offsetStoragePartitions;

    @Value("${debezium.kafka.offset.storage.replication-factor:1}")
    private int offsetStorageReplicationFactor;

    @Value("${debezium.kafka.status.storage.topic:#{null}}")
    private String statusStorageTopic;

    @Value("${debezium.kafka.status.storage.partitions:1}")
    private int statusStoragePartitions;

    @Value("${debezium.kafka.status.storage.replication-factor:1}")
    private int statusStorageReplicationFactor;

    private final AbdbTableNamesConfiguration abdbTableNamesConfiguration;

    public DebeziumConnectorConfiguration(
            AbdbTableNamesConfiguration abdbTableNamesConfiguration) {
        this.abdbTableNamesConfiguration = abdbTableNamesConfiguration;
    }

    /**
     * database connector.
     *
     * @return Configuration.
     */
    @Bean
    public io.debezium.config.Configuration abdbConnector() {
        io.debezium.config.Configuration.Builder builder = io.debezium.config.Configuration.create()
                .with("connector.class", "io.debezium.connector.postgresql.PostgresConnector")
                .with("offset.flush.interval.ms", offsetFlushIntervalMs)
                .with("name", "student-postgres-connector")
                .with("database.server.name", dbHost + "_" + dbName)
                .with("database.hostname", dbHost)
                .with("database.port", dbPort)
                .with("database.user", dbUserName)
                .with("database.password", dbPassword)
                .with("database.dbname", dbName)
                .with("plugin.name", "pgoutput")
                .with("snapshot.mode", "exported")
                .with("table.include.list", tableNames())
                .with("slot.name", "oprosso");
        if (offsetStorageFileFilename != null) {
            builder.with("offset.storage", "org.apache.kafka.connect.storage.FileOffsetBackingStore")
                   .with("offset.storage.file.filename", offsetStorageFileFilename);
        } else if (isValidDebeziumKafkaConfig()) {
            builder.with("bootstrap.servers", bootstrapServers)
                    .with("ssl.keystore.location", keyStoreLocation)
                    .with("ssl.keystore.password", keyStorePassword)
                    .with("ssl.truststore.location", trustStoreLocation)
                    .with("ssl.truststore.password", trustStorePassword)
                    .with("config.storage", "org.apache.kafka.connect.storage.KafkaConfigBackingStore")
                    .with("config.storage.topic", configStorageTopic)
                    .with("config.storage.partitions", configStoragePartitions)
                    .with("config.storage.replication.factor", configStorageReplicationFactor)
                    .with("offset.storage", "org.apache.kafka.connect.storage.KafkaOffsetBackingStore")
                    .with("offset.storage.topic", offsetStorageTopic)
                    .with("offset.storage.partitions", offsetStoragePartitions)
                    .with("offset.storage.replication.factor", offsetStorageReplicationFactor)
                    .with("status.storage", "org.apache.kafka.connect.storage.KafkaStatusBackingStore")
                    .with("status.storage.topic", statusStorageTopic)
                    .with("status.storage.partitions", statusStoragePartitions)
                    .with("status.storage.replication.factor", statusStorageReplicationFactor);
        } else
            throw new IllegalArgumentException();
        if (clientId != null) {
            builder.with("client-id", clientId);
            builder.with("client.id", clientId);
            builder.with("consumer.client-id", clientId);
            builder.with("consumer.client.id", clientId);
            builder.with("producer.client-id", clientId);
            builder.with("producer.client.id", clientId);
        }
        if (securityInterBrokerProtocol != null)
            builder.with("security.inter.broker.protocol", securityInterBrokerProtocol);
        if (securityProtocol != null)
            builder.with("security.protocol", securityProtocol);
        if (sslEndpointIdentificationAlgorithm != null)
            builder.with("ssl.endpoint.identification.algorithm", sslEndpointIdentificationAlgorithm);

        return builder.build();
    }

    private String tableNames() {
        return String.join(", ", abdbTableNamesConfiguration.getNames());
    }

    boolean isValidDebeziumKafkaConfig() {
        return bootstrapServers != null
                && configStorageTopic != null
                && offsetStorageTopic != null
                && statusStorageTopic != null;
    }
}
