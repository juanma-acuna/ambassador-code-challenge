# CODE CHALLENGE

This challenge has three main folders:

- **backend**
- **frontend**
- **kubernetes** (kubernetes configuration)

## Backend

The backend is written in Go (golang) and its task is to read a json file and obtain the information of the Services, to later present this services by endpoints that can be consumed by third parties.

## Frontend

The frontend is written in React, keeping the responsibilities and functions of each part separate. The entry point shows an index with each of the services displayed by the backend, allowing them to be accessed independently. Functions and hooks are in separate folders.

## Kubernetes
The Kubernetes configurations are in a folder with the same name. In this folder, you will find all the files that are necessary to deploy the services.

## Deployment

In some computers we have detected that it is necessary to execute the following command to be able to access the resource.

`minikube tunnel`

To run the project from minikube, open a terminal in the root folder of the project and run the following command:

`kubectl apply -f kubernetes`

This could take a few minutes, depending of your internet connection, because it is necessary to download the images of the services.

Once this is done, you will be able to access the project by http://localhost:80

To terminate the project, you will need to run the following command:

`kubectl delete -f kubernetes`

### Updating the code

If you need to modify the code, is needed to rebuild and push the image to the hub. You must run the following command inside the folder you modified, p.e., if you modified part of the frontend, you must run the command inside the frontend folder:

- Create a docker image of the backend and send to hub.
  - `docker build -t gps1mx/go-challenge:1 .`
  - `docker push gps1mx/go-challenge:1`
- Create a docker image of the frontend and send to hub.
  - `docker build -t gps1mx/react-challenge:1 .`
  - `docker push gps1mx/react-challenge:1`

Once the push has finished, you must run the following command from the root of the project to update the deployment:

`kubectl delete -f kubernetes && kubectl apply -f kubernetes`
