# Graceful

Graceful demonstrates graceful shutdown of a service using sync.WaitGroup.

It adds each handler executed to sync.WaitGroup which can be waited on before exiting.

