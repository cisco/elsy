package com.cisco;

import com.cisco.db.NotesDao;
import com.cisco.resources.NotesResource;
import io.dropwizard.Application;
import io.dropwizard.configuration.EnvironmentVariableSubstitutor;
import io.dropwizard.configuration.SubstitutingSourceProvider;
import io.dropwizard.db.DataSourceFactory;
import io.dropwizard.jdbi.DBIFactory;
import io.dropwizard.migrations.MigrationsBundle;
import io.dropwizard.setup.Bootstrap;
import io.dropwizard.setup.Environment;
import org.skife.jdbi.v2.DBI;

public class JavaNoteServiceApplication extends Application<JavaNoteServiceConfiguration> {

    public static void main(final String[] args) throws Exception {
        new JavaNoteServiceApplication().run(args);
    }

    @Override
    public String getName() {
        return "Java Note Service";
    }

    @Override
    public void initialize(final Bootstrap<JavaNoteServiceConfiguration> bootstrap) {
        // Enable variable substitution with environment variables
        bootstrap.setConfigurationSourceProvider(
                new SubstitutingSourceProvider(bootstrap.getConfigurationSourceProvider(),
                        new EnvironmentVariableSubstitutor(false)
                )
        );

        bootstrap.addBundle(new MigrationsBundle<JavaNoteServiceConfiguration>() {
            @Override
            public DataSourceFactory getDataSourceFactory(JavaNoteServiceConfiguration configuration) {
                return configuration.getDataSourceFactory();
            }
        });
    }

    @Override
    public void run(final JavaNoteServiceConfiguration config,
                    final Environment environment) {
        final DBIFactory factory = new DBIFactory();
        final DBI jdbi = factory.build(environment, config.getDataSourceFactory(), "mysql");
        final NotesDao dao = jdbi.onDemand(NotesDao.class);
        environment.jersey().register(new NotesResource(dao));
    }

}
