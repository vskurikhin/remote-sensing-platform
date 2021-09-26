package su.svn.invalidator.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.format.datetime.standard.DateTimeFormatterRegistrar;
import org.springframework.format.support.DefaultFormattingConversionService;
import org.springframework.format.support.FormattingConversionService;

import static java.time.format.DateTimeFormatter.ofPattern;

@Configuration
public class DateTimeConfig {
    @Bean
    public FormattingConversionService conversionService() {
        DefaultFormattingConversionService conversionService = new DefaultFormattingConversionService(false);

        DateTimeFormatterRegistrar registrar = new DateTimeFormatterRegistrar();
        registrar.setDateFormatter(ofPattern("dd-MM-yyyy"));
        registrar.setDateTimeFormatter(ofPattern("dd-MM-yyyy'Z'HH:mm:ss.SSS"));
        registrar.registerFormatters(conversionService);

        return conversionService;
    }
}
