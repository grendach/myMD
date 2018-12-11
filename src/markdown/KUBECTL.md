Full list of commands U'll find in Kubect [cheatsheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)

## Creating Objects
* create resource(s)
    ```
    $ kubectl create -f ./my-manifest.yaml
    ```
* create from multiple files
    ```
    $ kubectl create -f ./my1.yaml -f ./my2.yaml
    ```
* create resource(s) in all manifest files in dir
    ```
    $ kubectl create -f ./dir
    ```
* create resource(s) from url
    ```
    $ kubectl create -f https://git.io/vPieo
    ```
* start a single instance of nginx
    ```
    $ kubectl run nginx --image=nginx
    ```
* get the documentation for pod and svc manifests
    ```
    $ kubectl explain pods,svc
    ```
## Get commands with basic output
* List all services in the namespace
    ```
    $ kubectl get services
    ```
* List all pods in all namespaces
    ```
    $ kubectl get pods --all-namespaces
    ```
* List all pods in the namespace, with more details
    ```
    $ kubectl get pods -o wide
    ```
* List a particular deployment
    ```
    $ kubectl get deployment my-dep
    ```
* List all pods in the namespace, including uninitialized ones
    ```
    $ kubectl get pods --include-uninitialized
    ```

## Describe commands with verbose output
* get node, pod description
    ```
    $ kubectl describe nodes my-node
    $ kubectl describe pods my-pod
    ```
* list Services Sorted by Name
    ```
    $ kubectl get services --sort-by=.metadata.name 
    ```

* list pods Sorted by Restart Count
    ```
    $ kubectl get pods --sort-by='.status.containerStatuses[0].restartCount'
    ```
* Get the version label of all pods with label app=cassandra
    ```
    kubectl get pods --selector=app=cassandra rc -o \
    jsonpath='{.items[*].metadata.labels.version}'
    ```

* Get all running pods in the namespace
    ```
    $ kubectl get pods --field-selector=status.phase=Running
    ```

* Get ExternalIPs of all nodes
    ```
    $ kubectl get nodes -o jsonpath='{.items[*].status.addresses[?(@.type=="ExternalIP")].address}'
    ```
* List Names of Pods that belong to Particular RC
 "jq" command useful for transformations that are too complex for jsonpath, it can be found at https://stedolan.github.io/jq/
    ```
    $ sel=${$(kubectl get rc my-rc --output=json | jq -j '.spec.selector | to_entries | .[] | "\(.key)=\(.value),"')%?}

    $ echo $(kubectl get pods --selector=$sel --output=jsonpath={.items..metadata.name})
    ```

* Check which nodes are ready
    ```
    JSONPATH='{range .items[*]}{@.metadata.name}:{range @.status.conditions[*]}{@.type}={@.status};{end}{end}' \
    && kubectl get nodes -o jsonpath="$JSONPATH" | grep "Ready=True"
    ```


* List all Secrets currently in use by a pod
    ```
    kubectl get pods -o json | jq '.items[].spec.containers[].env[]?.valueFrom.secretKeyRef.name' | grep -v null | sort | uniq
    ```
* List Events sorted by timestamp
    ```
    kubectl get events --sort-by=.metadata.creationTimestamp
    ```
## Updating Resources
* Rolling update pods of frontend-v1
    ```
    $ kubectl rolling-update frontend-v1 -f frontend-v2.json
    ```
* Change the name of the resource and update the image
    ```
    $ kubectl rolling-update frontend-v1 frontend-v2 --image=image:v2
    ```
* Update the pods image of frontend
    ```
    $ kubectl rolling-update frontend --image=image:v2
    ```
* Abort existing rollout in progress
    ```
    $ kubectl rolling-update frontend-v1 frontend-v2 --rollback
    ```
* Replace a pod based on the JSON passed into stdin
    ```
    $ cat pod.json | kubectl replace -f -
    ```


## Editing Resources

* Edit the service named docker-registry
    ```
    $ kubectl edit svc/docker-registry
    ```
* Use an alternative editor
    ```
    $ KUBE_EDITOR="nano" kubectl edit svc/docker-registry
    ```
## Scaling Resources
* Scale a replicaset named 'foo' to 3
    ```
    $ kubectl scale --replicas=3 rs/foo
    ```
* Scale a resource specified in "foo.yaml" to 3
    ```
    $ kubectl scale --replicas=3 -f foo.yaml
    ```
* If the deployment named mysql's current size is 2, scale mysql to 3
    ```
    $ kubectl scale --current-replicas=2 --replicas=3 deployment/mysql
    ```
* Scale multiple replication controllers
    ```
    $ kubectl scale --replicas=5 rc/foo rc/bar rc/baz
    ```
## Deleting Resources
* Delete a pod using the type and name specified in pod.json
    ```
    $ kubectl delete -f ./pod.json
    ```
* Delete pods and services with same names "baz" and "foo"
    ```
    $ kubectl delete pod,service baz foo
    ```
* Delete pods and services with label name=myLabel
    ```
    $ kubectl delete pods,services -l name=myLabel
    ```
* Delete pods and services, including uninitialized ones, with label name=myLabel
    ```
    $ kubectl delete pods,services -l name=myLabel --include-uninitialized
    ```
* Delete all pods and services, including uninitialized ones, in namespace my-ns,
    ```
    $ kubectl -n my-ns delete po,svc --all
    ```
## Interacting with running Pods
* dump pod logs (stdout)
    ```
    $ kubectl logs my-pod
    ```
* dump pod logs (stdout) for a previous instantiation of a container
    ```
    $ kubectl logs my-pod --previous
    ```
* dump pod container logs (stdout, multi-container case)
    ```
    $ kubectl logs my-pod -c my-container
    ```
* dump pod container logs (stdout, multi-container case) for a previous instantiation of a container
    ```
    $ kubectl logs my-pod -c my-container --previous
    ```
* stream pod logs (stdout)
    ```
    $ kubectl logs -f my-pod
    ```
* stream pod container logs (stdout, multi-container case)
    ```
    $ kubectl logs -f my-pod -c my-container
    ```
* Run pod as interactive shell
    ```
    $ kubectl run -i --tty busybox --image=busybox -- sh
    ```
* Attach to Running Container
    ```
    $ kubectl attach my-pod -i
    ```
* Listen on port 5000 on the local machine and forward to port 6000 on my-pod
    ```
    $ kubectl port-forward my-pod 5000:6000
    ```
* Run command in existing pod (1 container case)
    ```
    $ kubectl exec my-pod -- ls /
    ```
* Get a shell to the running Container in some namespace (1 container case)
    ```
    $ kubectl -n namespace exec -it testclient -- /bin/bash
    ```
* Run command in existing pod (multi-container case)
    ```
    $ kubectl exec my-pod -c my-container -- ls /
    ```
* Show metrics for a given pod and its containers
    ```
    $ kubectl top pod POD_NAME --containers
    ```
## Interacting with Nodes and Cluster

* Mark my-node as unschedulable
    ```
    kubectl cordon my-node
    ```
* Drain my-node in preparation for maintenance 
    ```
    kubectl drain my-node
    ```
* Mark my-node as schedulable
    ```
    kubectl uncordon my-node
    ```
* Show metrics for a given node
    ```
    $ kubectl top node my-node
    ```
* Display addresses of the master and services
    ```
    $ kubectl cluster-info
    ```
* Dump current cluster state to stdout
    ```
    $ kubectl cluster-info dump
    ```
* Dump current cluster state to /path/to/cluster-state
    ```
    $ kubectl cluster-info dump --output-directory=/path/to/cluster-state
    ```