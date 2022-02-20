package de.koware.blumhuegel.transformers;

import org.apache.beam.sdk.coders.RowCoder;
import org.apache.beam.sdk.schemas.Schema;
import org.apache.beam.sdk.transforms.DoFn;
import org.apache.beam.sdk.values.Row;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;

/*
    Builds rows from single-type iterables
 */
public class SimpleRowBuilder extends DoFn<List<?>, Row> {

    private static final Logger LOGGER = LoggerFactory.getLogger(SimpleRowBuilder.class);

    private final Schema rowSchema;

    public SimpleRowBuilder(Schema rowSchema) {
        this.rowSchema = rowSchema;
    }

    @ProcessElement
    public void processElement(@Element List<String> record, OutputReceiver<Row> out) {
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
                        .addArray(
                        record
                ).build()
        );
    }


}
