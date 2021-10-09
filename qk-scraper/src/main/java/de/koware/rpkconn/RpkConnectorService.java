package de.koware.rpkconn;

import de.koware.finnhubClient.FinnhubService;
import org.eclipse.microprofile.rest.client.RestClientBuilder;

import java.net.MalformedURLException;
import java.net.URL;

/*
    Connector for Redpanda Kafka-compatible message broker.
 */
public class RpkConnectorService {

    private final FinnhubService finnhubService;

    public RpkConnectorService() throws MalformedURLException {
        this.finnhubService = RestClientBuilder.newBuilder()
                .baseUrl(new URL("https://finnhub.io/api"))
                .build(FinnhubService.class);
    }


}
