## Docker commands:
For more info check [official documentation](https://docs.docker.com/) or  [GIT](https://github.com/docker)



* delete all images:
    ```
    $ docker images purge
    ```
* delete all containers

    ```
    $ docker rm $(docker ps -a -q)
    ```
* delete all images
    ```
    $ docker rmi $(docker images -q)
    ```

*  how to find proper running docker process for kafka
    ```
    $ docker ps | grep kafka
    ```
* login to bash inside running docker container
    ```
    $ docker exec -ti d1730c8eb42c bash
    ```
*  docker build image
    ```
    $ docker-compose -f docker-compose-dev.yml build
    ```
*  docker run builded image in background
    ```
    $ docker-compose -f docker-compose-dev.yml up -d
    ```
* run docker image under the proxy and with some parameters
    ```
    $ docker run -d --restart=unless-stopped \
    -p 80:80 -p 443:443 \
    -e HTTP_PROXY="http://10.144.1.10:8080/" \
    -e HTTPS_PROXY="http://10.144.1.10:8080/" \
    -e NO_PROXY="localhost,127.0.0.1,0.0.0.0,192.168.10.0/24,example.com," \
    rancher/rancher:latest
    ```