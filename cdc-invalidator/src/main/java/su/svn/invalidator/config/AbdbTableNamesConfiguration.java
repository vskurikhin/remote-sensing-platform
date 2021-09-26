package su.svn.invalidator.config;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

import java.util.ArrayList;
import java.util.List;

@Configuration
@ConfigurationProperties("abdb.datasource.table")
public class AbdbTableNamesConfiguration {
    private List<String> names = new ArrayList<String>();

    public List<String> getNames() {
        return names;
    }
}
