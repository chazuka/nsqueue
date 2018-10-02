## Prequirisite

NSQ installed, refer to [official documentation](https://nsq.io/overview/quick_start.html)

## Running Sample

1. clone repository
2. run `dep ensure -v` to install dependencies
3. run cli task

## Running NSQ

running nsqlookupd

```
nsqlookupd
```

running nsqd with nsqlookupd address
```
nsqd --lookupd-tcp-address=127.0.0.1:4160
```

running nsq dashboard
```
nsqadmin --lookupd-http-address=127.0.0.1:4161
```

## CLI

publishing sample messages
```
./nsqueue -topic=testing -channel=testing -task=pub -address=127.0.0.1:4150
```

consuming messages
```
./nsqueue -topic=testing -channel=testing -task=sub -address=127.0.0.1:4150
```

