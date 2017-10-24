# deploy-example

This repository contains a simple [Ruby](https://www.ruby-lang.org) application. All configuration to run the app is included.

The following tasks will guage an applicants ability to understand the application and how to run and deploy it in containerized environments.

## How to Use This Repository

- Fork this repository into your Github profile
- Make changes, add files, and push to your fork to complete the following tasks
- When completed, create a Pull Request with all updates and include any additional comments

## Task 1

Run the [Ruby](https://www.ruby-lang.org) app locally using [Sinatra](http://www.sinatrarb.com) and [Rack](http://rack.github.io) and demonstrate how to use the web application.

## Task 2

Run the app locally within a [Docker](https://docs.docker.com/engine/) container

## Task 3

Run the app using [docker-compose](https://docs.docker.com/compose/) with [nginx](https://docs.docker.com/compose/) as a reverse proxy to the application.

## Task 4

Setup [TravisCI](https://travis-ci.org), [CircleCI](https://circleci.com), or other automated continous integration to build a container and test on push to Github repository.

## Task 5

Add deployment files for [Rancher](http://rancher.com), [Kubernetes](https://kubernetes.io) or other container [orchestration and management framework](https://github.com/cncf/landscape)

## Bonus

Additional tasks to further show mastery of application deployments and understanding.

- Add an additional route to the application to return full list of gifs
- Minimize container image
- Run application on platform other than x86
- Rewrite application in another language
