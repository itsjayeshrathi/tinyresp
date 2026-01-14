# RESP Parser (Minimal)

A minimal implementation of the Redis Serialization Protocol (RESP), built directly from the specification.

The goal is correctness at the protocol boundary:
byte-level framing, CRLF handling, and strict adherence to RESP types, without Redis-specific assumptions.

This is not a Redis clone and not a production-ready client — it is a learning project focused on understanding how real wire protocols work.

## Reference

Docs **[RESP](https://redis.io/docs/latest/develop/reference/protocol-spec/)**

Handwritten Parsers & Lexers in Go **[Article](https://blog.gopheracademy.com/advent-2014/parsers-lexers/)**
