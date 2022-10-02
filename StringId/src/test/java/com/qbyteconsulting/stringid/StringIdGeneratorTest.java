package com.qbyteconsulting.stringid;

import org.junit.jupiter.api.MethodOrderer;
import org.junit.jupiter.api.Order;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestMethodOrder;

import java.io.IOException;

import static org.junit.jupiter.api.Assertions.*;

@TestMethodOrder(MethodOrderer.OrderAnnotation.class)
class StringIdGeneratorTest {

    StringIdGenerator generator;

    StringIdGeneratorTest() throws IOException {
        generator = StringIdGenerator.getInstance();
    }

    @Test
    @Order(1)
    void test01_Identity_Is_Case_Insensitive() {
        assertEquals("0000000001", generator.stringToId("abc").get());
        assertEquals("0000000001", generator.stringToId("Abc").get());
        assertEquals("0000000001", generator.stringToId("ABC").get());
    }

    @Test
    @Order(2)
    void test02_Identity_For_Thai_Characters() {
        assertEquals("0000000002", generator.stringToId("กขค").get());
    }

    @Test
    @Order(3)
    void test03_Identity_For_Mix_With_Thai_Characters() {
        assertEquals("0000000003", generator.stringToId("ก_8").get());
    }

    @Test
    @Order(4)
    void test04_Identity_For_Emoji_Character() {
        assertEquals("0000000004", generator.stringToId("\uD83D\uDE00").get());
    }

    @Test
    @Order(5)
    void test05_Identity_For_255_Character_String() {
        assertEquals("0000000005",
                generator.stringToId("MfF78dIAiZen7RR7zaqSXxEpPGEVDy5wltAjLLzhL63y3km5rvRTcWSFbrbq5z7TzGNu7dyRj93EmA1DLlYlXCbDVZQk6eBFAOpNJGfVROAFbc1vViXkAwYHwUYJeTcFUBwE1M7fiGNKEzZZxb3CV2PYu59fQJvJIoA2dk8YkFr4MlSJ0CBzhgSJchLAlUUgqEfRFkdznEh21rItbsDctb3GdkMVxSWpfa9kJmbWxxIolcy0QXnc2j5Y4ywakSq")
                        .get()
        );
    }

    @Test
    @Order(6)
    void test06_Failure_For_256_Character_String() {
        assertThrows(IllegalArgumentException.class, () ->
                generator.stringToId("MfF78dIAiZen7RR7zaqSXxEpPGEVDy5wltAjLLzhL63y3km5rvRTcWSFbrbq5z7TzGNu7dyRj93EmA1DLlYlXCbDVZQk6eBFAOpNJGfVROAFbc1vViXkAwYHwUYJeTcFUBwE1M7fiGNKEzZZxb3CV2PYu59fQJvJIoA2dk8YkFr4MlSJ0CBzhgSJchLAlUUgqEfRFkdznEh21rItbsDctb3GdkMVxSWpfa9kJmbWxxIolcy0QXnc2j5Y4ywakSqI")
        );
    }

    @Test
    @Order(7)
    void test07_Test_Non_Alphanumeric_Printable_Characters() {
        assertEquals("0000000006", generator.stringToId(" ~`!@#$%&^*()_-+={}[]|\\:;<,>.?/").get());
    }

    @Test
    @Order(8)
    void test08_Capacity_Overload() {
        fail("test to overload huge capacity impractical");
    }
}