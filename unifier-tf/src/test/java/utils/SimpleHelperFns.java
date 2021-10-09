package utils;

import org.apache.beam.sdk.transforms.DoFn;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class SimpleHelperFns {

    private static final Logger LOGGER = LoggerFactory.getLogger(SimpleHelperFns.class);

    public static class PrintFn extends DoFn<Object, Object> {
        @ProcessElement
        public void processElement(@Element Object object, OutputReceiver<Object> out) {
            LOGGER.info("object is: {}", object);
            out.output(object);
        }
    }

    public static class StringPrintFn extends DoFn<String, String> {
        @ProcessElement
        public void processElement(@Element String word, OutputReceiver<String> out) {
            LOGGER.info("token is: {}", word);

            out.output(word);
        }
    }

}
