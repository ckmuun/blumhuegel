package beam;

import de.koware.blumhuegel.transformers.SimpleTokenizerFn;
import io.quarkus.test.junit.QuarkusTest;
import org.apache.beam.runners.direct.DirectRunner;
import org.apache.beam.sdk.Pipeline;
import org.apache.beam.sdk.coders.StringUtf8Coder;
import org.apache.beam.sdk.io.TextIO;
import org.apache.beam.sdk.options.PipelineOptions;
import org.apache.beam.sdk.options.PipelineOptionsFactory;
import org.apache.beam.sdk.transforms.Create;
import org.apache.beam.sdk.transforms.DoFn;
import org.apache.beam.sdk.transforms.ParDo;
import org.apache.beam.sdk.values.KV;
import org.apache.beam.sdk.values.PCollection;
import org.junit.jupiter.api.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Arrays;
import java.util.List;

@QuarkusTest
public class BeamTest {

    private static final Logger LOGGER = LoggerFactory.getLogger(BeamTest.class);

    @Test
    public void testLog() {
        LOGGER.info("test123");
    }


    @Test
    public void testBeamWordCount() {

        PipelineOptions pOptions = PipelineOptionsFactory.create();

        Pipeline pipeline = Pipeline.create(pOptions);

        LOGGER.info("Creating PCollection");
        final List<String> LINES = Arrays.asList(
                "To be, or not to be: that is the question: ",
                "Whether 'tis nobler in the mind to suffer ",
                "The slings and arrows of outrageous fortune, ",
                "Or to take arms against a sea of troubles, ");

        pipeline.apply(Create.of(LINES)).setCoder(StringUtf8Coder.of())
                .apply(ParDo.of(new SimpleTokenizerFn()))
                .apply(ParDo.of(new StringPrintFn()))
                .apply(ParDo.of(new WordsWithLengthFn()))
                .apply(ParDo.of(new PrintFn()));

        pipeline.run().waitUntilFinish();
    }


    @Test
    public void testPipeFromFile() {

        PipelineOptions pOptions = PipelineOptionsFactory.create();
        pOptions.setRunner(DirectRunner.class);

        Pipeline pipeline = Pipeline.create(pOptions);

        PCollection<String> words = pipeline
                .apply(
                        "ReadMyFile", TextIO.read().from("src/test/resources/someWords.txt"))
                .apply(ParDo.of(
                        new SimpleTokenizerFn()
                ))
                .apply(ParDo.of(
                        new StringPrintFn()
                ));

        PCollection<Integer> wordLengths = words.apply(
                ParDo
                        .of(new ComputeWordLengthFn()));


        wordLengths.setName("the collection");

        LOGGER.info("running pipeline");
        pipeline.run();
    }

    public static class WordsWithLengthFn extends DoFn<String, KV<String, Integer>> {
        @ProcessElement
        public void processElement(@Element String word, OutputReceiver<KV<String, Integer>> out) {

            out.output(KV.of(
                    word,
                    word.length()
            ));
        }
    }

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

    static class ComputeWordLengthFn extends DoFn<String, Integer> {
        @ProcessElement
        public void processElement(@Element String word, OutputReceiver<Integer> out) {
            // Use OutputReceiver.output to emit the output element.
            out.output(word.length());
        }
    }
}
