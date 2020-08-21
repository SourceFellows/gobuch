Starten des docker Containers
-----------------------------

```
docker run -v `pwd`:/tmp -it openjdk:latest /bin/bash
```

Danach im Container
```
cd /tmp
time javac Main.java
```