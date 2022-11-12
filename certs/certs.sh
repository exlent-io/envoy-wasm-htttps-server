openssl genrsa -passout pass:1111 -out envoy-root-ca.key 4096
openssl req -new -x509 -days 365 -key envoy-root-ca.key -subj "/C=TW/ST=Taiwan/L=Taipei/O=Example/OU=IT/CN=Example Root CA" -out envoy-root-ca.crt -passin pass:1111 -set_serial 0 -extensions v3_ca -config openssl.cnf

openssl req -nodes -new -keyout envoy-intermediate-ca.key -out envoy-intermediate-ca.csr -subj "/C=TW/ST=Taiwan/L=Taipei/O=Example/OU=IT/CN=Example Intermediate CA" -config openssl.cnf -passout pass:1111
openssl x509 -days 365 -req -in envoy-intermediate-ca.csr -CAcreateserial -CA envoy-root-ca.crt -CAkey envoy-root-ca.key -out envoy-intermediate-ca.crt -extfile openssl.cnf -extensions v3_intermediate_ca -passin pass:1111

openssl req -nodes -new -keyout envoy-proxy-server.key -out envoy-proxy-server.csr -config openssl.cnf -passout pass:1111
openssl x509 -days 365 -req -in envoy-proxy-server.csr -CAcreateserial -CA envoy-intermediate-ca.crt -CAkey envoy-intermediate-ca.key -out envoy-proxy-server.crt -extfile openssl.cnf -extensions server_cert -passin pass:1111

# openssl req -nodes -new -keyout envoy-proxy-client.key -out envoy-proxy-client.csr -config openssl.cnf -passout pass:1111
# openssl x509 -days 365 -req -in envoy-proxy-client.csr -CAcreateserial -CA envoy-intermediate-ca.crt -CAkey envoy-intermediate-ca.key -out envoy-proxy-client.crt -extfile openssl.cnf -extensions usr_cert -passin pass:1111

# cat envoy-root-ca.crt envoy-intermediate-ca.crt > envoy-intermediate-and-envoy-root-ca-chain.crt

