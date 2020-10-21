Starten des docker Containers
-----------------------------

```
docker run -v `pwd`:/tmp -it gcc:latest /bin/bash
```

Danach im Container
```
cd /tmp
time g++ hello.cpp
```
