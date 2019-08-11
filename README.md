# Card

### 如何生成pem和pk  
`openssl genrsa -out ./testdata/server.key 2048`  
`openssl req -new -x509 -key ./testdata/server.key -out ./testdata/server.pem -days 365`