Starten des docker Containers
-----------------------------

```
docker run -v `pwd`:/tmp -it python:latest /bin/bash
```

Danach im Containers
```
cd /tmp
time python hello.py
```