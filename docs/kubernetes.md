Kubernetes
==========

Tests below were done with [Minikube](https://kubernetes.io/docs/setup/minikube/), locally, but works with any other Kubernetes cluster.

Imperatively
------------

```
kubectl run sitego --image=ovicrisan/sitego --port=8000
kubectl expose deployment sitego --type=NodePort
```

Then, you can see each separate resource:

```
kubectl get deployments
kubectl get rs
kubectl get pods
kubectl get services
```

Or see all of them with 

```
kubectl get all
```

You can start the browser pointing to your local cluster IP and proper port number with

```
minikube service sitego
```

Or instead, just try it from terminal with 

```
curl http://192.168.99.100:31475/
```

After testing the site in the browser you can delete everything (and make sure of this operation) with

```
kubectl delete deployment sitego
kubectl delete service sitego
kubectl get all
```

Declaratively
-------------

Download the YAML file [sitego_kube.yaml](sitego_kube.yaml) and run

```
kubectl create -f sitego_kube.yaml
minikube service sitego
```

If you want to scale it try:

```
kubectl scale --replicas=3 deploy/sitego
```
