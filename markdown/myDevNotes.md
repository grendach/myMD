# Docker commands:

1. delete all images:
    ```
    $ docker images purge
    ```
2. delete all containers

    ```
    $ docker rm $(docker ps -a -q)
    ```
3. delete all images
    ```
    $ docker rmi $(docker images -q)
    ```

# Vault commands:

 
 1. Disable proxy for localhost:
    ```
    $ unset http_proxy
    ```
 2. Start Vault server on "dev mode" with "root" token on local machine:
    ```
    $ vault server -dev -dev-root-token-id="root"
    ```

 3. Check health of Vault
    ```
    $ curl http://127.0.0.1:8200/v1/sys/health
    ```
 4. Add some secret to Vault server using json file
    ```
    $ curl -H "X-Vault-Token: root" -X POST --data @/tmp/my_secret.json 0:8200/v1/secret/my-secret | jq
    ```
    or using '--data' key in your curl command:
    ```
    $ curl -H "X-Vault-Token: root" -X POST --data '{"she": "Yana"}'  0:8200/v1/secret/my-secret | jq
    ```
 5. Add your plugin to Vault server:
    ```
    $ curl -X PUT 0:8200/v1/sys/plugins/catalog/vabar -d '{"sha_256": "88d4ed3ad15cde18d649d9324434ab978a0d0c434dd5523738d252899ad312e9", "command": "vabar"}' -H "X-Vault-Token:c865-56d5-62cd-4220-15fda0ae8665"
    ```
 6. Mount your plugin to Vault server:
    ```
    $ curl -H "X-Vault-Token: root" -X GET  http://127.0.0.1:8200/v1/sys/mounts | jq
    ```
 7. Check list of Vault plugins, your plugin shour be in a list:
    ```
    $ curl -H "X-Vault-Token: root" -X LIST 0:8200/v1/sys/plugins/catalog | jq
    ```
 8. Run you plugin and check if it's work:
    ```
    $ curl -H "X-Vault-Token: root" -X POST 0:8200/v1/vabar/backup |
    ```
# Linux commands:

1. How to use SCP from local to remote server using private *.pem key
```
$ scp -r user@your.server.example.com:/path/to/foo /home/user/Desktop/
```



# how to find proper running docker process for kafka
docker ps | grep kafka

#login to bash inside running docker container
docker exec -ti d1730c8eb42c bash

#docker build image
docker-compose -f docker-compose-dev.yml build

#docker run builded image in background
docker-compose -f docker-compose-dev.yml up -d 

#find files which containes specific string
grep -rnw '/path/to/somewhere/' -e 'pattern'


#create new branch in git
git branch <branch-name>

#create and switch to new git branch
git checkout -b AVADEVOPS-626

#check git log with graphical mode
git log --graph --decorate --pretty=oneline --abbrev-commit

#force git push 
git push --force origin AVADEVOPS-626

#branch squash to 1 commit.

git rebase -i HEAD~[NUMBER OF COMMITS]

#OR
git rebase -i [SHA]

#If you have previously pushed your code to a remote branch, you will need to force push.

git push origin branchName --force

#show git log in one line


https://gitlabe1.ext.net.nokia.com/ava/ava-core/blob/master/group_vars/all/resources.yml

#using ipdb module for debug python program

import ipdb
ipdb.set_trace()


declare -x http_proxy="http://10.144.1.10:8080"
declare -x https_proxy="http://10.144.1.10:8080"



##install Python3.6 and SetUp it on linux:
#if you're using Ubuntu 16.04 LTS, you'll need to use a PPA:

sudo add-apt-repository ppa:jonathonf/python-3.6 # (only for 16.04 LTS)
#Then, run the following (this works out-of-the-box on 16.10 and 17.04):

sudo apt update
sudo apt install python3.6
sudo apt install python3.6-dev
sudo apt install python3.6-venv
wget https://bootstrap.pypa.io/get-pip.py
sudo python3.6 get-pip.py
sudo ln -s /usr/bin/python3.6 /usr/local/bin/python3
sudo ln -s /usr/local/bin/pip /usr/local/bin/pip3

# Do this only if you want python3 to be the default Python
# instead of python2 (may be dangerous, esp. before 2020):
# sudo ln -s /usr/bin/python3.6 /usr/local/bin/python

# launch venv for python3.6-de
virtualenv --python=python3.6 venv




#get all data from CouchDB

curl -X GET http://couchdb.caas.marathon.mesos:5984/_all_docs

#creating ssh tunnel
ssh wrodev3-openstack-slave-13 -L 5984:couchdb.caas.marathon.mesos:5984 -N

#show size of folder
du -h --max-depth=1 ~/MY/

#write iso image into pendrive
dd if=/home/grendach/Downloads/sw/ubuntu-18.04-desktop-amd64.iso of=/dev/sdb

#show filesystem
lsblk

docker run -e BUCKET_NAME=couch-backup-bucket -e COUCHDB_URL=http://couchdb.caas.marathon.mesos:5984/ -e LOG_RETENTION=2 -e S3_ACCESS_KEY=9f4ba036e47647caa6022aeb125ffc28 -e S3_SECRET_KEY=e9f841d0769a40fe99aa1a73fbca9b25 -e S3_URL=http://es-si-s3-z4.eecloud.nsn-net.net:10152 ava-docker-local.esisoj70.emea.nsn-net.net/ava/couchbackup:1.1.0

#run minikube behind the proxy:
no_proxy=$no_proxy,192.168.1.0/16 minikube start --vm-driver kvm2 --kvm-network minikube-net --docker-env http_proxy=$http_proxy --docker-env https_proxy=$https_proxy --docker-env no_proxy=$no_proxy,192.168.1.0/16

https_proxy=http://10.144.1.10:8080 minikube start --vm-driver=virtualbox --docker-env http_proxy=$http_proxy --docker-env https_proxy=$https_proxy --docker-env no_proxy=$no_proxy,192.168.1.0/16

#run minikube with specific version:
minikube start --kubernetes-version="v1.10.0" --vm-driver=kvm2

#delete minikube:
$minikube stop
$minikube delete && rm -rf ~/.minikube && rm -rf ~/.kube


