Nutzer der EB-Cli über Docker-Container:

Image erstellen und ausführen:
```
docker build -t elasticbeanstalk-cli .
docker run -v `pwd`:/app -it elasticbeanstalk-cli /bin/bash
```

Innerhalb des Containers müssen folgende Kommandos ausgeführt werden:
```
eb init -p go go-beanstalk-sample --region eu-central-1
eb create go-env
```