Starten des docker Containers
-----------------------------

```
docker run -v `pwd`:/tmp -it node:latest /bin/bash
```

Danach im Containers
```
cd /tmp
node index.js
```