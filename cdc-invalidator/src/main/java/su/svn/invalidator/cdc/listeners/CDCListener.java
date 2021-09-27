package su.svn.invalidator.cdc.listeners;

import io.lettuce.core.RedisClient;
import io.lettuce.core.api.StatefulRedisConnection;
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
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.Executor;
import java.util.concurrent.Executors;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.function.BiConsumer;

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

    private final RedisClient redisClient;
    private final StatefulRedisConnection<String, String> connection;

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
        this.redisClient = RedisClient.create("redis://oprosso@localhost/0");
        this.connection = redisClient.connect();
        log.info("Connected to Redis");
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
            connection.close();
            redisClient.shutdown();
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
        Struct source = (Struct) sourceRecordValue.get("source");
        String table = source.get("table").toString();
        switch (table) {
            case "screen_main":
            case "group":
            case "question":
            case "question_audio":
            case "question_card_sorting":
            case "question_closed":
            case "question_comparison":
            case "question_csi":
            case "question_first_click":
            case "question_matrix":
            case "question_media":
            case "question_nps":
            case "question_opened":
            case "question_opinion":
            case "question_password":
            case "question_ranging":
            case "question_rating":
            case "question_semantic_differential":
            case "question_site_test":
            case "question_slideshow":
            case "question_test_text":
            case "question_tree_testing":
                extracted(sourceRecordValue);
        }
        log.info("handleEvent: {}", sourceRecordValue);
    }

    private void extracted(Struct sourceRecordValue) {
        try {
            String op = sourceRecordValue.get("op").toString();
            Struct after = (Struct) sourceRecordValue.get("after");
            int id = Integer.parseInt(after.get("id").toString());
            String key = String.format("InvalidateArrayOfFScreenMain-pollItemId-%d", id);
            String value = connection.sync().get(key);
            connection.async().del(value);
            connection.async().del(key);
            log.info("value: {}", value);
        } catch (Exception e) {
            log.error("handleEvent: ", e);
        }
    }
}
