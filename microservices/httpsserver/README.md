Zertifikate erstellen
=====================

- Root Zertifikat erstellen:
```
openssl genrsa -out myRoot.key 2048
openssl req -x509 -new -nodes -key myRoot.key -sha256 -days 3650 -out myRoot.crt
```

- Zertifikat für Server erstellen:
```
#CSR erstellen
openssl req -new -sha256 -nodes -out test.example.csr -newkey rsa:2048 -keyout test.example.key -config server-cert.request.conf

openssl x509 -req -in test.example.csr -CA myRoot.crt -CAkey myRoot.key -CAcreateserial -extfile test.example.ext.cnf -out test.example.crt -days 3650 -sha256
```