## Producer

```
                                            go run producer/main.go                  ✔  10455  13:49:56
2023/07/08 13:56:02 INF    1 (127.0.0.1:4150) connecting to nsqd
```

## Consumer

```
                                go run consumer/main.go                  ✔  10453  13:47:53
2023/07/08 13:48:08 INF    1 [NSQ_Topic/NSQ_Channel] (127.0.0.1:4150) connecting to nsqd
2023/07/08 13:48:08 Awaiting messages from NSQ topic "My NSQ Topic"...
2023/07/08 13:48:14 NSQ message received:
2023/07/08 13:48:14 Hello NSQ, Hello Hello Hello
2023/07/08 13:49:55 NSQ message received:
2023/07/08 13:49:55 Hello NSQ, Hello Hello Hello
2023/07/08 13:56:02 NSQ message received:
2023/07/08 13:56:02 Hello NSQ, Hello Hello Hello
2023/07/08 13:56:47 ERR    1 [NSQ_Topic/NSQ_Channel] (127.0.0.1:4150) IO error - EOF
2023/07/08 13:56:47 INF    1 [NSQ_Topic/NSQ_Channel] (127.0.0.1:4150) beginning close
2023/07/08 13:56:47 INF    1 [NSQ_Topic/NSQ_Channel] (127.0.0.1:4150) breaking out of writeLoop
2023/07/08 13:56:47 INF    1 [NSQ_Topic/NSQ_Channel] (127.0.0.1:4150) writeLoop exiting
2023/07/08 13:56:47 INF    1 [NSQ_Topic/NSQ_Channel] (127.0.0.1:4150) readLoop exiting
2023/07/08 13:56:47 INF    1 [NSQ_Topic/NSQ_Channel] (127.0.0.1:4150) clean close complete
2023/07/08 13:56:47 WRN    1 [NSQ_Topic/NSQ_Channel] there are 0 connections left alive
2023/07/08 13:56:47 INF    1 [NSQ_Topic/NSQ_Channel] (127.0.0.1:4150) re-connecting in 1m0s
2023/07/08 13:56:47 INF    1 [NSQ_Topic/NSQ_Channel] (127.0.0.1:4150) finished draining, cleanup exiting
```
