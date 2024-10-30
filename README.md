# Application Idea

A simple todo application with a scalable backend  and multiple components capable of scaling individually.

# Components
There are three main components of the app which is deployed using Kubernetes 

 - Mongodb : serves as the data store for the backend
 - Shopping list service : has the integration with th backend and contains the business logic
 - Shopping list api : Fullfils the authentication and provides the API functionality to the frontend

## Technologies used

Both service and api projects are built using golang echosystem. Shell scripts are added to build and publish docker images for the respective project. 

I have used Dockerhub as the docker repository.
Images can be found here https://hub.docker.com/repositories/kdkanishka

I have used minikube in my linux development environment.

For the frontend I have used Flutter, and compiling it as a web target.

## Deployment

under infrastructure/application there are several kubernetes defenition files to create required resources.
I have written a shell script `deploy.sh` to deploy all of these resources at once.

Once deployed application can be accessed by following below instructions.

    kubectl get services

It will show existing services

    NAME                       TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
    kubernetes                 ClusterIP   10.96.0.1       <none>        443/TCP          35h
    mongodb                    ClusterIP   10.101.33.181   <none>        27017/TCP        27h
    shoppinglist-api-service   NodePort    10.102.141.28   <none>        8080:30080/TCP   27h
    shoppinglist-service       ClusterIP   10.104.80.230   <none>        8080/TCP         27h

Then get the url  for the `shoppinglist-api-service` 

    minikube service shoppinglist-api-service --url
It will display the url as below

    http://192.168.49.2:30080

## ToDo
Complete the web application part. due to the deadline of the submittion I am submitting with the partially completed web app. But the backend is fully completed.


