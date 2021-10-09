package de.koware.finnhubClient;


import org.eclipse.microprofile.rest.client.inject.RegisterRestClient;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.QueryParam;

@Path("/v1")
public interface FinnhubService {


    @GET
    @Path("/stock/metric")
    public StockSymbolDto getBaseFinancials(@QueryParam("symbol") String symbol, @QueryParam("token") String token);
}
