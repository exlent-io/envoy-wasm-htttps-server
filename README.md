# Envoy Wasm Https Server
This project is based on reference [stevesloka/envoy-xds-server](https://github.com/stevesloka/envoy-xds-server), which is modified to support TLS and WASM functionality.


```
$ envoy --version 
envoy  version: 15baf56003f33a07e0ab44f82f75a660040db438/1.24.0/Distribution/RELEASE/BoringSSL

$ go version
go version go1.19 darwin/arm64

$ rustup --version
rustup 1.25.1 (bb60b1e89 2022-07-12)
info: This is the version for the rustup toolchain manager, not the rustc compiler.
info: The currently active `rustc` version is `rustc 1.65.0 (897e37553 2022-11-02)`
```

## Refs
https://content.red-badger.com/resources/extending-istio-with-rust-and-webassembly

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
    $ touch version.txt
    $ ./server --port 18000 --configFileName config/config.yaml --watchVersionFileName version.txt
    ```

* Apply new config after edit config file

    > Note that using vim to edit `version.txt` files may cause issue. You can refer to see [this](https://unix.stackexchange.com/questions/36467/why-inode-value-changes-when-we-edit-in-vi-editor).
    ```bash
    $ date > version.txt
    ```
