import io.quarkus.test.junit.QuarkusTest;
import org.junit.jupiter.api.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@QuarkusTest
public class DummyTest {

    private static final Logger LOGGER = LoggerFactory.getLogger(DummyTest.class);

    @Test
    public void testLog() {
        LOGGER.info("Test123 ");
        assert true;
    }

}
