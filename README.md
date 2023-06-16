
# simply distributed in-mem cache 🚀

### simple KV cache for simple usage

## The server will start listening for incoming connections on the specified address. Clients can connect to the server and send commands to interact with the cache. The supported commands are:
* SET <key> <value> [ttl]: Sets a key-value pair in the cache with an optional time-to-live (TTL) duration.
* GET <key>: Retrieves the value associated with the given key from the cache.
* JOIN: Not implemented yet.
#### NOTE:The server will handle these commands and respond accordingly. The cache operations will be performed based on the received commands.

