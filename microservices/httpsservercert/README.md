Zertifikate erstellen
=====================

AUSFÜHRE IM cert VERZEICHNIS

- Root Zertifikat erstellen:
```
openssl genrsa -out myRoot.key 2048
openssl req -x509 -new -nodes -key myRoot.key -sha256 -days 3650 -out myRoot.crt
```

- Zertifikat für Server erstellen:
```
#CSR erstellen
openssl req -new -sha256 -nodes -out test.example.csr -newkey rsa:2048 -keyout test.example.key -config server-cert.request.conf

openssl x509 -req -in test.example.csr -CA myRoot.crt -CAkey myRoot.key -CAcreateserial -out test.example.crt -days 3650 -sha256
```

-CA anlegen
```
mkdir -p ca/certsdb
touch ca/index.txt
touch ca/index.txt.attr
echo '01' > ca/serial
```

- Client-Zertifikat erstellen

```
openssl genrsa -des3 -out user.key # key is 1234
openssl req -new -key user.key -out user.req -config user-cert.request.conf
openssl ca -cert test.example.crt -keyfile test.example.key -out user.crt -in user.req -config user-cert.conf
openssl pkcs12 -export -in user.crt -inkey user.key -out user-client.p12
```

- Chain

```
cat
```

- Curl
```
curl --cacert ./certs/myRoot.crt --cert ./certs/user.crt --key ./certs/user.key -v https://test.example.private:8443
```