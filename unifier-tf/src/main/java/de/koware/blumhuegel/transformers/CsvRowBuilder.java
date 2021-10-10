package de.koware.blumhuegel.transformers;

import org.apache.beam.sdk.schemas.Schema;
import org.apache.beam.sdk.transforms.DoFn;
import org.apache.beam.sdk.values.Row;
import org.apache.commons.csv.CSVFormat;
import org.apache.commons.csv.CSVParser;
import org.apache.commons.csv.CSVRecord;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.io.StringReader;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/*
    Builds rows from single-type iterables
 */
public class CsvRowBuilder extends DoFn<String, Row> {

    private static final Logger LOGGER = LoggerFactory.getLogger(CsvRowBuilder.class);

    private final Schema rowSchema;

    public CsvRowBuilder(Schema rowSchema) {
        this.rowSchema = rowSchema;
    }

    @ProcessElement
    public void processElement(@Element String record, OutputReceiver<Row> out) {
        LOGGER.trace("creating row out of record: {}", record);

        CSVRecord csvRecord = separateCsvRecord(record);

        if (null == csvRecord) {
            out.output(null);
            return;
        }
        Object[] values = csvRecord.stream().sequential().toArray();

        Row.Builder builder = Row.withSchema(rowSchema);
        out.output(
                builder.withFieldValues(getSchemaValues(values)).build()
        );
    }


    private Map<String, Object> getSchemaValues(Object[] values) {
        HashMap<String, Object> schemaValues = new HashMap<>();

        for (int i = 0; i < rowSchema.getFieldCount(); i++) {
            schemaValues.put(
                    rowSchema.getField(i).getName(),
                    values[i]
            );
        }
        return schemaValues;
    }

    private CSVRecord separateCsvRecord(String record) {
        StringReader stringReader = new StringReader(record);
        CSVFormat format = CSVFormat.newFormat(',');

        try {

            CSVParser parser = new CSVParser(stringReader, format);
            return parser.getRecords().get(0);

        } catch (IOException ioe) {
            LOGGER.error("csv record parsing error");
            ioe.printStackTrace();
        }

        // TODO Check apache beam error handling philosophy (e.g. null tolerance)
        return null;
    }
}
