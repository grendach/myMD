
## Linux commands:

* create `archive.zip` file from `dir`
    ```
    $ zip -r 'archive.zip' /home/dir/
    ```

*  find files which containes specific string
    ```
    $ grep -rnw '/path/to/somewhere/' -e 'pattern'
    ```

* copy via SCP from local to remote server using private *.pem key
    ```
    $ scp -r user@your.server.example.com:/path/to/foo /home/user/Desktop/
    ```
* set up proxy server
    ```
    $ declare -x http_proxy="http://10.144.1.10:8080"
    $ declare -x https_proxy="http://10.144.1.10:8080"
    $ declare -x ftp_proxy="http://10.144.1.10:8080"
    ```
    or as `sudo user` add proxy servers to `/etc/environment` file

* create venv for python3.7
    ```
    $ virtualenv --python=python3.7 v
    ```

* creating ssh tunnel
    ```
    $ ssh wrodev3-openstack-slave-13 -L 5984:couchdb.caas.marathon.mesos:5984 -N
    ```

* show size of folder
    ```
    $ du -h --max-depth=1 <dir>
    ```

* show list block devices
    ```
    $ lsblk
    ```
* write `filename.iso` image to pendrive
    ```
    $ dd if=/home/grendach/Downloads/sw/ubuntu-18.04-desktop-amd64.iso of=/dev/sdb
    ```
* get all data from CouchDB
    ```
    $ curl -X GET <link_to_server>:<port>/_all_docs
    ```

* run minikube behind the proxy:
    ```
    $ no_proxy=$no_proxy,192.168.1.0/16 minikube start --vm-driver kvm2 --kvm-network minikube-net --docker-env http_proxy=$http_proxy --docker-env https_proxy=$https_proxy --docker-env no_proxy=$no_proxy,192.168.1.0/16
    ```

* run minikube with specific version:
    ```
    $ minikube start --kubernetes-version="v1.10.0" --vm-driver=kvm2
    ```
* delete minikube:
    ```
    $ minikube stop
    
    $ minikube delete && rm -rf ~/.minikube && rm -rf ~/.kube
    ```
* create virtualenv `v` with python3.7:
    ```
    $ virtualenv -p /usr/bin/python3.7 v
    ```
* run pytest and generate coverage report:
    ```
    $ PYTHONPATH=. pytest --cov --cov-report=html --cov-report=console
    ```