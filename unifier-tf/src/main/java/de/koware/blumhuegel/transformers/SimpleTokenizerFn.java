package de.koware.blumhuegel.transformers;

import org.apache.beam.sdk.metrics.Counter;
import org.apache.beam.sdk.metrics.Distribution;
import org.apache.beam.sdk.metrics.Metrics;
import org.apache.beam.sdk.transforms.DoFn;


/*
    More or less taken from Apache Beam Examples (github.com/apache/beam)
 */
public class SimpleTokenizerFn extends DoFn<String, String> {

    public static final String TOKENIZER_PATTERN = "[^\\p{L}]+";


    private final Counter emptyLines = Metrics.counter(SimpleTokenizerFn.class, "emptyLines");
    private final Distribution lineLenDist =
            Metrics.distribution(SimpleTokenizerFn.class, "lineLenDistro");

    @ProcessElement
    public void processElement(@Element String element, OutputReceiver<String> receiver) {
        lineLenDist.update(element.length());
        if (element.trim().isEmpty()) {
            emptyLines.inc();
        }

        // Split the line into words.
        String[] words = element.split(TOKENIZER_PATTERN, -1);

        // Output each word encountered into the output PCollection.
        for (String word : words) {
            if (!word.isEmpty()) {
                receiver.output(word);
            }
        }
    }
}
