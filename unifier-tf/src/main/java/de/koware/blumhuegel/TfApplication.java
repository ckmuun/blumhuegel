package de.koware.blumhuegel;

import io.quarkus.runtime.Quarkus;
import io.quarkus.runtime.annotations.QuarkusMain;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@QuarkusMain
public class TfApplication {

    private static final Logger LOGGER = LoggerFactory.getLogger(TfApplication.class);

    public static void main(String[] args) {

        LOGGER.info("starting quarkus main");
        Quarkus.run(args);
        LOGGER.info("quarkus main started");
    }
}
