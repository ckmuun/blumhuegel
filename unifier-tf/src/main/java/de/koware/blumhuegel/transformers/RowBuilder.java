package de.koware.blumhuegel.transformers;

import org.apache.beam.sdk.schemas.Schema;
import org.apache.beam.sdk.transforms.DoFn;
import org.apache.beam.sdk.values.PCollection;
import org.apache.beam.sdk.values.Row;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;

/*
    Builds rows from single-type iterables
 */
public class RowBuilder extends DoFn<String, Row> {

    private static final Logger LOGGER = LoggerFactory.getLogger(RowBuilder.class);

    private final Schema rowSchema;

    public RowBuilder(Schema rowSchema) {
        this.rowSchema = rowSchema;
    }

    @ProcessElement
    public void processElement(@Element String record, OutputReceiver<Row> out) {
        LOGGER.trace("creating row out of record: {}", record);

        Schema fundamentalsSchema = Schema.of(
                Schema.Field.of("period", Schema.FieldType.STRING),
                Schema.Field.of("company", Schema.FieldType.STRING),
                Schema.Field.of("ticker", Schema.FieldType.STRING),
                Schema.Field.of("indicator", Schema.FieldType.STRING),
                Schema.Field.of("amount", Schema.FieldType.STRING)
        );


        out.output(

                Row.withSchema(fundamentalsSchema)
                        .withFieldValue("period", "bob")
                        .withFieldValue("company", "bob one")
                        .withFieldValue("ticker", "bob two")
                        .withFieldValue("indicator", "bob three")
                        .withFieldValue("amount", "bob five")
                        .build());
    }

}
