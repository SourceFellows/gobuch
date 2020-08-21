Starten des docker Containers
-----------------------------

```
docker run -v `pwd`:/tmp -it node:latest /bin/bash
```

Danach im Container
```
npm install -g tsc
cd /tmp
time tsc hello.ts
```
