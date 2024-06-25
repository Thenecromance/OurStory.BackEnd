# **this folder is whole project's documentation folder**

## Basic Information

- this project is a simple web application's backend
- also this is used to record the trivialities between couples

## Folder Structure

```
├── Document        all documents of the project (just because I might always forget the project's structure)
├── Interface       all interfaces of the project
├── application     the logic services implemented in here
├── constants       all constants value of the project stored in this folder
├── middleware      due to this project based on gin, it can support middlewares(also other web frameworks should be supported to)
├── response        structuring the response of the project
├── route           route management of the project
├── server          a monkey core in here
├── thirdParty      third party packages include my other projects or some services might use other third party APIs' wrapper will be in here
└── utility         utility functions or some packages I have no idea where to put in
```

### ./application

all based on MVC design pattern, so the folder structure is like below

```
├── ReadMe.md   this is the documentation of the application folder
├── controller  controllers implemented in here
├── models      models implemented in here
├── repository  repositories layer implemented in here
└── services    services layer implemented in here
```