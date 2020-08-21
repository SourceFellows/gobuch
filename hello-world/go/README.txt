Starten des docker Containers
-----------------------------

```
docker run -v `pwd`:/tmp -it golang:latest /bin/bash
```

Danach im Container
```
cd /tmp
time go build main.go
```
