package su.svn.invalidator.config;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;
import java.time.OffsetDateTime;
import java.time.format.DateTimeFormatter;

import static su.svn.invalidator.config.WebMvcConfig.OFFSET_DATE_TIME_FORMATTER;

public class CustomOffsetDateTimeSerializer extends JsonSerializer<OffsetDateTime> {
    private DateTimeFormatter formatter;

    public CustomOffsetDateTimeSerializer() {
        this(OFFSET_DATE_TIME_FORMATTER);
    }

    public CustomOffsetDateTimeSerializer(DateTimeFormatter formatter) {
        this.formatter = formatter;
    }

    @Override
    public void serialize(OffsetDateTime value, JsonGenerator gen, SerializerProvider provider) throws IOException {
        gen.writeString(value.format(this.formatter));
    }
}
