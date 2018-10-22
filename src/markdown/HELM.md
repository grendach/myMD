* Deploying chart with service type `NodePort`
    ```sh
    $ helm install --name example ./mychart --set service.type=NodePort
    ```