package su.svn.invalidator.config;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import com.fasterxml.jackson.datatype.jsr310.deser.LocalDateTimeDeserializer;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateTimeSerializer;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.converter.HttpMessageConverter;
import org.springframework.http.converter.json.MappingJackson2HttpMessageConverter;
import org.springframework.scheduling.annotation.EnableAsync;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

import java.time.LocalDateTime;
import java.time.OffsetDateTime;
import java.time.format.DateTimeFormatter;
import java.time.format.DateTimeFormatterBuilder;
import java.time.temporal.ChronoField;
import java.util.List;

import static java.time.format.DateTimeFormatter.ISO_LOCAL_DATE_TIME;
import static java.time.format.DateTimeFormatter.ofPattern;

@Configuration
@EnableWebMvc
@EnableAsync
public class WebMvcConfig implements WebMvcConfigurer {

    public static final DateTimeFormatter OFFSET_DATE_TIME_FORMATTER = new DateTimeFormatterBuilder()
            // date/time
            .appendPattern("yyyy-MM-dd'T'HH:mm:ss")
            // optional fraction of seconds (from 0 to 9 digits)
            .optionalStart().appendFraction(ChronoField.NANO_OF_SECOND, 0, 9, true)
            .optionalEnd()
            // offset
            .appendPattern("xxx")
            // create formatter
            .toFormatter();
    public static final DateTimeFormatter FORMATTER = ofPattern("dd-MM-yyyy'T'HH:mm:ss.SSS");

    @Override
    public void configureMessageConverters(List<HttpMessageConverter<?>> converters) {
        LocalDateTimeSerializer localDateTimeSerializer = new LocalDateTimeSerializer(ISO_LOCAL_DATE_TIME);
        LocalDateTimeDeserializer localDateTimeDeserializer = new LocalDateTimeDeserializer(ISO_LOCAL_DATE_TIME);

        JavaTimeModule module = new JavaTimeModule();
        module.addSerializer(LocalDateTime.class, localDateTimeSerializer);
        module.addDeserializer(LocalDateTime.class, localDateTimeDeserializer);
        module.addSerializer(OffsetDateTime.class, new CustomOffsetDateTimeSerializer(OFFSET_DATE_TIME_FORMATTER));
        module.addDeserializer(OffsetDateTime.class, new CustomOffsetDateTimeDeserializer(OFFSET_DATE_TIME_FORMATTER));

        ObjectMapper mapper = new ObjectMapper();
        mapper.registerModule(module);
        // add converter at the very front
        // if there are same type mappers in converters, setting in first mapper is used.
        converters.add(0, new MappingJackson2HttpMessageConverter(mapper));
    }
}
