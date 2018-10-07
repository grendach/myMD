
# Vault commands:
For more info check [official documentation](https://www.vaultproject.io/docs/index.html) or  [GIT](https://github.com/hashicorp/vault)
 
* disable proxy for localhost:
    ```
    $ unset http_proxy
    ```
* start Vault server on "dev mode" with "root" token on local machine:
    ```
    $ vault server -dev -dev-root-token-id="root"
    ```

* check health of Vault
    ```
    $ curl http://127.0.0.1:8200/v1/sys/health
    ```
*  add some secret to Vault server using json file
    ```
    $ curl -H "X-Vault-Token: root" -X POST --data @/tmp/my_secret.json 0:8200/v1/secret/my-secret | jq
    ```
    or using `--data` key in your curl command:
    ```
    $ curl -H "X-Vault-Token: root" -X POST --data '{"capital": "Kyiv"}'  0:8200/v1/secret/my-secret | jq
    ```
*  add your plugin to Vault server:
    ```
    $ curl -X PUT 0:8200/v1/sys/plugins/catalog/vabar -d '{"sha_256":  d4ed3ad15cde18d649d9324434ab978a0d0c434dd5523738d252899ad312e9", "command": "vabar"}' -H     Vault-Token:c865-56d5-62cd-4220-15fda0ae8665"
    ```
*  mount your plugin to Vault server:
    ```
    $ curl -H "X-Vault-Token: root" -X GET  http://127.0.0.1:8200/v1/sys/mounts | jq
    ```
*  check list of Vault plugins, your plugin shour be in a list:
    ```
    $ curl -H "X-Vault-Token: root" -X LIST 0:8200/v1/sys/plugins/catalog | jq
    ```
*  run you plugin and check if it's work:
    ```
    $ curl -H "X-Vault-Token: root" -X POST 0:8200/v1/vabar/backup | jq
    ```