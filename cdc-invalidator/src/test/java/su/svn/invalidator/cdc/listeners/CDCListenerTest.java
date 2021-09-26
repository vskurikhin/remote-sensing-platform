package su.svn.invalidator.cdc.listeners;

class CDCListenerTest {
//
//    private static final Logger LOG = org.slf4j.LoggerFactory.getLogger(CDCListenerTest.class);
//
//    private static final SpecificData MODEL = new SpecificData();
//
//    Envelope envelope;
//
//    Envelope test;
//
//    BinaryMessageEncoder<Envelope> envelopeEncoder;
//
//    BinaryMessageDecoder<Envelope> envelopeDecoder;
//
//    long sign;
//
//    @BeforeEach
//    void setUp() {
//        sign = Long.MAX_VALUE;
//        ByteBuffer bb = ByteBuffer.allocate(CRC32_CAPACITY);
//        ByteBuffer buffer = bb.putLong(sign).rewind();
//        envelope = Envelope.newBuilder()
//                .setVersion(1)
//                .setSignAlg("CRC32")
//                .setSign(buffer)
//                .setMContainer(buffer)
//                .setTypeName("TYPE")
//                .build();
//        envelopeEncoder = new BinaryMessageEncoder<>(MODEL, Envelope.getClassSchema());
//        envelopeDecoder = new BinaryMessageDecoder<>(MODEL, Envelope.getClassSchema());
//    }
//
//    @AfterEach
//    void tearDown() {
//        envelopeEncoder = null;
//        envelopeDecoder = null;
//        envelope = null;
//        test = null;
//    }
//
//    @Test
//    void test() throws IOException {
//        ByteBuffer byteBuffer = envelopeEncoder.encode(envelope);
//        test = envelopeDecoder.decode(byteBuffer);
//        LOG.info("test = {}", test);
//        Assertions.assertEquals(envelope, test);
//        long signTest = test.getSign().getLong();
//        LOG.info("sign = {}", signTest);
//        Assertions.assertEquals(sign, signTest);
//    }
//
//    @Test
//    void test1() throws IOException {
//        BinaryMessageEncoder<AbtStatus> abtStatusEncoder = new BinaryMessageEncoder<>(MODEL, AbtStatus.getClassSchema());
//        ByteBuffer buffer = abtStatusEncoder.encode(ABT_STATUS_EXPECTED);
//        envelope.setMContainer(buffer);
//        sign = Crc32.crc32(buffer, 0, buffer.capacity());
//        ByteBuffer bb = ByteBuffer.allocate(CRC32_CAPACITY);
//        ByteBuffer lb = bb.putLong(sign).rewind();
//        envelope.setSign(lb);
//        test();
//        BinaryMessageDecoder<AbtStatus> abtStatusDecoder = new BinaryMessageDecoder<>(MODEL, AbtStatus.getClassSchema());
//        AbtStatus testAbtStatus = abtStatusDecoder.decode(test.getMContainer());
//        LOG.info("testAbtStatus = {}", testAbtStatus);
//        Assertions.assertEquals(ABT_STATUS_EXPECTED, testAbtStatus);
//    }
//
//    /*
//    @Test
//    void test2() throws IOException, URISyntaxException {
//        ByteBuffer buffer = null;
//        try {
//            RandomAccessFile aFile = new RandomAccessFile("src/test/resources/oneMessage.avro","r");
//
//            FileChannel inChannel = aFile.getChannel();
//            long fileSize = inChannel.size();
//
//            buffer = ByteBuffer.allocate((int) fileSize);
//            inChannel.read(buffer);
//            buffer.flip();
//
//            inChannel.close();
//            aFile.close();
//        }
//        catch (IOException e) {
//            LOG.error("test2 ", e);
//            throw e;
//        }
//        test = envelopeDecoder.decode(buffer);
//        LOG.info("test = {}", test);
//    }
//     */
}