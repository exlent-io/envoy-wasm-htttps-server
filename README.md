# Envoy Wasm Https Server

## Sample

* Create certs
```bash
$ cd certs
$ certs.sh
```

* Start envoy
```bash
$ envoy -c config/bootstrap.yaml #-l debug
```

* Build and run management server
```bash
$ go build ./cmd/server
$ server --port 18000 --watchDirectoryFileName config/config.yaml
```
