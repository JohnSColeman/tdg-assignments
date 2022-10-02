package com.qbyteconsulting.stringid;

import java.io.*;
import java.nio.file.*;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;

/**
 * Provides a utility function for providing unique string identification.
 * To guarantee uniqueness the identity values are sequentially generated and persisted to local file storage.
 */
public class StringIdGenerator {

    public static final int IDENTITY_SIZE = 10;
    public static final String IDENTITY_FORMAT = "%0" + IDENTITY_SIZE + "d";
    public static final double MAX_SEQUENCE = Math.pow(10, IDENTITY_SIZE) - 1;
    public static final int MAX_STRING_LENGTH = 255;
    public static final IllegalArgumentException ILLEGAL_ARGUMENT_LENGTH_EXCEPTION
            = new IllegalArgumentException("string value cannot be longer than 255 characters");
    public static final RuntimeException IDENTITY_PERSISTENCE_EXCEPTION
            = new RuntimeException("persistence mechanism failed");

    private static StringIdGenerator instance;

    private final Path logPath = Paths.get("StringIdGenerator.idx");
    private final ConcurrentHashMap<String, String> ids = new ConcurrentHashMap<>();

    private Long sequence;

    /**
     * Obtain a generator instance.
     *
     * @return an instance of the generator
     * @throws IOException thrown when the underlying index access fails
     */
    public static synchronized StringIdGenerator getInstance() throws IOException {
        if (instance == null) {
            instance = new StringIdGenerator();
        }
        return instance;
    }

    private StringIdGenerator() throws IOException {
        sequence = 0L;
        loadState();
    }

    /**
     * Returns a unique digital ID for the given input string value (case insensitive), throws IllegalArgumentException
     * if the given value is longer than 255 characters - (an emoji counts as 2 characters).
     *
     * @param value the string value to obtain a digital ID
     * @return the optional unique 10 digit ID of the string or empty if no more IDs available
     */
    public Optional<String> stringToId(String value) {
        if (value.length() > MAX_STRING_LENGTH) throw ILLEGAL_ARGUMENT_LENGTH_EXCEPTION;
        return Optional.ofNullable(ids.computeIfAbsent(value.toLowerCase(),
                key -> {
                    if (sequence > MAX_SEQUENCE) return null;
                    return persistedKeyIdentity(key, nextIdentity(++sequence));
                })
        );
    }

    private void loadState() throws IOException {
        try (DataInputStream stream = new DataInputStream(new BufferedInputStream(Files.newInputStream(logPath)))) {
            while (stream.available() > 0) {
                ids.put(stream.readUTF(), nextIdentity(++sequence));
            }
        } catch (NoSuchFileException ignored) {
            // expected
        }
    }

    private String nextIdentity(long sequence) {
        return String.format(IDENTITY_FORMAT, sequence);
    }

    private String persistedKeyIdentity(String key, String identity) {
        if (identity != null) {
            try (DataOutputStream stream = new DataOutputStream(new BufferedOutputStream(Files.newOutputStream(logPath,
                    StandardOpenOption.CREATE, StandardOpenOption.APPEND), MAX_STRING_LENGTH))) {
                stream.writeUTF(key);
            } catch (IOException ex) {
                throw IDENTITY_PERSISTENCE_EXCEPTION;
            }
        }
        return identity;
    }
}
