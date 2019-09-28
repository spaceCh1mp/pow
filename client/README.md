# Introduction

This is the view (frontend) for the web app pow, for a better description of the application visit the main page. This is built using ReactJs and handles UI/UX functionality.

## Table of content

1. [Structure.](#structure)
2. [Components.](#components)
3. [Usage.](#usage)
4. [Testing.](#testing)

## Structure

    Client/
        build/
        node-modules/ 
        public/
            favicon.png
            index.html
            manifest.json
        src/
            components/
                component-name/
                    component-name.js
                    component-name.css
                    component-name.test.js
            css/
            pages/
            index.js
            routes.js
            serviceWorker.js

- build/ is the location of the production-ready build. This directory won’t exist until you run npm build or yarn build.

- node_modules/ is houses packages installed by NPM or Yarn.

- public/ is where your static files reside.

- src/ is houses the main application code and components.

## Components

    coming soon

## Usage

There are two parts to the client application, the server already takes care of ensuring both parts are up and running in production.

For developers trying to run the client as a stand-alone, run the command

    :~client$ npm install.

from the client directory to ensure that the dependencies are installed.

To start the React application, run:

    :~client$ npm start

### Envoy

Now it's time to run the envoy image. Ensure that you have docker installed on your system, for a step by step guide to install docker on your local machine visit [docker's website](https://docs.docker.com/install/).

Run:

    :~client$ docker build ./src -t pow/envoy:latest
    :~client$ docker run -d -p 8081:8081 --net=host pow/envoy

The first command creates an image from the dockerfile in client/src/

The second command runs the docker image and binds the hosts port 8081 to the container.

To see containers that are running use:

    :~$ docker container list

### Protocol buffers

For developers, the proto file is located in the src/helpers/api directory.

Ensure that protoc-gen-grpc-web package is installed (running npm install should take care of any dependencies) and that the path is set, to set the path run:

    :~client$ PATH=$PATH:node_modules/.bin

To create a new service client stub using the api.proto file run:

    :~client$ protoc -I=. src/helpers/api/api.proto --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.

the above command generates two files in the same src/helpers/api directory

>api_pb.js
>
>api_grpc_web_pb.js

## Testing

    coming soon

## Contributors

    coming soon
