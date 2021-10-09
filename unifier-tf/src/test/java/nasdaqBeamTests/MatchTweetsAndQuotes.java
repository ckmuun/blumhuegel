package nasdaqBeamTests;

import de.koware.blumhuegel.transformers.RowBuilder;
import de.koware.blumhuegel.transformers.SimpleRowBuilder;
import io.quarkus.test.junit.QuarkusTest;
import org.apache.beam.sdk.Pipeline;
import org.apache.beam.sdk.PipelineResult;
import org.apache.beam.sdk.coders.RowCoder;
import org.apache.beam.sdk.io.Compression;
import org.apache.beam.sdk.io.FileIO;
import org.apache.beam.sdk.io.TextIO;
import org.apache.beam.sdk.io.fs.EmptyMatchTreatment;
import org.apache.beam.sdk.options.PipelineOptions;
import org.apache.beam.sdk.options.PipelineOptionsFactory;
import org.apache.beam.sdk.schemas.AutoValueSchema;
import org.apache.beam.sdk.schemas.Schema;
import org.apache.beam.sdk.state.State;
import org.apache.beam.sdk.transforms.Create;
import org.apache.beam.sdk.transforms.JsonToRow;
import org.apache.beam.sdk.transforms.MapElements;
import org.apache.beam.sdk.transforms.ParDo;
import org.apache.beam.sdk.values.PCollection;
import org.apache.beam.sdk.values.Row;
import org.apache.beam.sdk.values.TypeDescriptor;
import org.apache.beam.sdk.values.TypeDescriptors;
import org.junit.jupiter.api.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import utils.SimpleHelperFns;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.regex.MatchResult;

@QuarkusTest
public class MatchTweetsAndQuotes {

    private static final Logger LOGGER = LoggerFactory.getLogger(MatchTweetsAndQuotes.class);

    // yeah, amount is a numeric value not a literal, but this will do for this simple test
    private Schema fundamentalsSchema = Schema.of(
            Schema.Field.of("period", Schema.FieldType.STRING),
            Schema.Field.of("company", Schema.FieldType.STRING),
            Schema.Field.of("ticker", Schema.FieldType.STRING),
            Schema.Field.of("indicator", Schema.FieldType.STRING),
            Schema.Field.of("amount", Schema.FieldType.STRING)
    );

    @Test
    public void readPrintCsvFile() {
        LOGGER.info("testing reading of stock exchange csv files");
        PipelineOptions pOptions = PipelineOptionsFactory.create();
        Pipeline readCsvPipe = Pipeline.create(pOptions);

        PCollection<String> fileContent = readCsvPipe.apply(
                "Read CSV", TextIO.read().from(
                        "src/test/resources/nasdaq-fundamentals/fundamentals_dataset.csv"
                )
        ).apply(
                ParDo.of(
                        new SimpleHelperFns.StringPrintFn()
                )
        );
        // just assert there is a result and things didn't blow up
        PipelineResult pipelineResult = readCsvPipe.run();
        PipelineResult.State pipelineState = pipelineResult.waitUntilFinish();
        assert null != pipelineState;
    }

    @Test
    public void convertCsvToRecord() {
        PipelineOptions pOptions = PipelineOptionsFactory.create();
        Pipeline pipeline = Pipeline.create(pOptions);


        PCollection<Row> fileContent = pipeline
                .apply(
                        "Read CSV", TextIO.read().from("src/test/resources/nasdaq-fundamentals/fundamentals_dataset.csv")
                ).apply("split record", MapElements
                        .into(TypeDescriptors.lists(TypeDescriptors.strings()))
                        .via(
                                (String record) -> Arrays.asList(record.split(","))
                        )
                ).apply(
                        //     ParDo.of(new SimpleRowBuilder(fundamentalsSchema))
                        "create rows", MapElements
                                .into(TypeDescriptor.of(Row.class))
                                .via(
                                        (List<String> record) -> Row.withSchema(fundamentalsSchema)

                                                .withFieldValue("period", "bob")
                                                .withFieldValue("company", "bob one")
                                                .withFieldValue("ticker", "bob two")
                                                .withFieldValue("indicator", "bob three")
                                                .withFieldValue("amount", "bob five")
                                                .build()

                                )
                )
                .setCoder(RowCoder.of(fundamentalsSchema));


        PCollection<Object> printCol = fileContent.apply(
                ParDo.of(new SimpleHelperFns.PrintFn())
        );

        pipeline.run()
                .waitUntilFinish();
    }

    @Test
    public void testRowCreation() {
        PipelineOptions pOptions = PipelineOptionsFactory.create();
        Pipeline pipeline = Pipeline.create(pOptions);

        PCollection<Row> rows = pipeline.apply(
                "Read CSV", TextIO.read().from("src/test/resources/nasdaq-fundamentals/fundamentals_dataset.csv")
        ).apply(
                ParDo.of(new RowBuilder(fundamentalsSchema))

        )
                .setCoder(RowCoder.of(fundamentalsSchema));

        PCollection<Object> printCol = rows.apply(
                ParDo.of(new SimpleHelperFns.PrintFn())
        );

        pipeline.run()
                .waitUntilFinish();

    }

}
