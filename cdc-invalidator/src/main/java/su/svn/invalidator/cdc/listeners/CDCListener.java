package su.svn.invalidator.cdc.listeners;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

import io.debezium.config.Configuration;
import io.debezium.embedded.EmbeddedEngine;
import org.apache.kafka.connect.data.Struct;
import org.apache.kafka.connect.source.SourceRecord;

import javax.annotation.PostConstruct;
import javax.annotation.PreDestroy;
import java.nio.ByteBuffer;
import java.util.concurrent.Executor;
import java.util.concurrent.Executors;

@Slf4j
@Component
public class CDCListener {

    public static final int CRC32_CAPACITY = 8;
    /**
     * Single thread pool which will run the Debezium engine asynchronously.
     */
    private final Executor executor = Executors.newSingleThreadExecutor();

    /**
     * The Debezium engine which needs to be loaded with the configurations, Started and Stopped - for the
     * CDC to work.
     */
    private final EmbeddedEngine engine;

    private final KafkaTemplate<String, ByteBuffer> kafkaTemplate;

    private final String topic;

    /**
     * Constructor which loads the configurations and sets a callback method 'handleEvent', which is invoked when
     * a DataBase transactional operation is performed.
     * @param abdbConnector
     * @param kafkaTemplate
     * @param topic
     */
    private CDCListener(
            Configuration abdbConnector,
            KafkaTemplate<String, ByteBuffer> kafkaTemplate,
            @Value("${app.cap-topic}") String topic) {
        this.engine = EmbeddedEngine
                .create()
                .using(abdbConnector)
                .notifying(this::handleEvent).build();
        this.kafkaTemplate = kafkaTemplate;
        this.topic = topic;
    }

    /**
     * The method is called after the Debezium engine is initialized and started asynchronously using the Executor.
     */
    @PostConstruct
    private void start() {
        this.executor.execute(engine);
    }

    /**
     * This method is called when the container is being destroyed. This stops the debezium, merging the Executor.
     */
    @PreDestroy
    private void stop() {
        if (this.engine != null) {
            this.engine.stop();
        }
    }

    /**
     * This method is invoked when a transactional action is performed on any of the tables that were configured.
     *
     * @param sourceRecord
     */
    private void handleEvent(SourceRecord sourceRecord) {
        Struct sourceRecordValue = (Struct) sourceRecord.value();
        log.info("handleEvent: {}", sourceRecordValue);
    }
}
