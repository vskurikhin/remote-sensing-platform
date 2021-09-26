package su.svn.invalidator.config;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;

import java.io.IOException;
import java.time.OffsetDateTime;
import java.time.format.DateTimeFormatter;

import static su.svn.invalidator.config.WebMvcConfig.OFFSET_DATE_TIME_FORMATTER;

public class CustomOffsetDateTimeDeserializer extends JsonDeserializer<OffsetDateTime> {
    private DateTimeFormatter formatter;

    public CustomOffsetDateTimeDeserializer() {
        this(OFFSET_DATE_TIME_FORMATTER);
    }

    public CustomOffsetDateTimeDeserializer(DateTimeFormatter formatter) {
        this.formatter = formatter;
    }

    @Override
    public OffsetDateTime deserialize(JsonParser parser, DeserializationContext context) throws IOException {
        return OffsetDateTime.parse(parser.getText(), this.formatter);
    }
}
