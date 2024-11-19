# Application Idea

A simple todo application with a scalable backend  and multiple components capable of scaling individually.

# Source code repository
All the source code can be found on https://github.com/kdkanishka/PA2577

infrastructure/application - contains the kubernetes defenitions for the resources

# Components of the architecture
architrue diagram is attached as an image. (simple shopping list architecture.png)

There are three main components of the app which is deployed using Kubernetes 

 #### Mongodb

 - **Role**  : Primary data store
 - **Responsibility** :  Store and manage the data for the application. It acts as the primary database for storing shopping list items
    
#### Shopping list service
- **Role**  : Business logic provider
- **Responsibility**  Handle the core business logic of the application. This includes processing requests from the Shopping List API, interacting with the MongoDB database, and performing CRUD operations on shopping lists & it's items.

#### Shopping list api
- **Role**  : API gateway
- **Responsibility** : Act as the entry point for client requests. It handles authentication, and routes requests to the backend service (Shopping List Service). Also this service supposed to service static content such as asserts and the SPA (single page application)

#### Architecture Principles
In this section I am trying to explain what are the main principals that I have followed in my design.

- **Scalability**

The architecture is designed to be scalable by leveraging Kubernetes. Each component can be scaled independently based on demand, ensuring that the system can handle varying loads efficiently. For example the backend service is decoupled from the API service and exposed via a kubenetes service. So it can be easily scaled in case of a request hike.

- **Microservices**

The application is divided into microservices (Shopping List Service and Shopping List API), each responsible for a specific function. This separation of concerns makes the system more modular and easier to manage. When requirements grow it is possible to define more modules in the same mannaar and follow the same pattern to seamlessly incooperate with other modules ( microservices )

**Containerization**

Docker is used to containerize the application components as well as application builds (go compilation happens in a different container and copiled to a lightweight container), As it ensures consistency across different environments and simplifying deployment.

- **Service Discovery**

Kubernetes services are used for service discovery, allowing components to locate and communicate with each other seamlessly.

- **Security**

Authentication is handled by the Shopping List API, ensuring that only authorized users can access the application. Future enhancements will include proper TLS setup for secure communication.


 ### Challanges
 - **Service discovery**  
 
 since my architrecutre contains 2 different scalable components, it should be able to discover the service instances from the calling service.

 For example, api service need to forward the authenticatd requests for the backend service. 

 I have solved this by using kubenetes services. Backend services are exposed as a ClusterIP service resource type. It will not be available to the outside network.

 When it comes to the Api service, it should be available to the outside world, So I have used NodePort for the sake of simplysity here. But ideally it should be an ingress with proper TLS setup.

## Technologies used

Both service and api projects are built using golang echosystem. Shell scripts are added to build and publish docker images for the respective project. 

I have used Dockerhub as the docker repository.
Images can be found here https://hub.docker.com/repositories/kdkanishka

I have used minikube in my linux development environment.

For the frontend I have used vanilla javascript, css and html to develop the single page webpage for the app.

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
Integrate the app with a authentication service. Until that I have hardcoded basic authentication credentials in the middleware for the demostration purposes.

