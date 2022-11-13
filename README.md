# Envoy Wasm Https Server
This project is based on reference [stevesloka/envoy-xds-server](https://github.com/stevesloka/envoy-xds-server), which is modified to support TLS and WASM functionality.

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
    $ server --port 18000 --configFileName config/config.yaml --watchVersionFileName version.txt
    ```

* Apply new config after edit config file

    > Note that using vim to edit `version.txt` files may cause issue. You can refer to see [this](https://unix.stackexchange.com/questions/36467/why-inode-value-changes-when-we-edit-in-vi-editor).
    ```bash
    $ date > version.txt
    ```
