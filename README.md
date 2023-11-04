# Gocron Graceful Exit

This is an example how to use goroutine and channel to wait for SIGINT or SIGTERM signal.

In this case, gocron already waits for the running tasks to complete before stopping. But if you're not using gocron, you might want to use waitgroup like on the [first version here](https://github.com/jeremyleonardo/gocron-graceful-exit/blob/0a9e280577e89e26741dd842fd90f35aa86d6510/main.go).
