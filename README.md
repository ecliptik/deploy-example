# deploy-example

This repository contains a simple [Ruby](https://www.ruby-lang.org) application. All configuration to run the app is included.

The following tasks will gauge an applicants ability to understand the application and how to run and deploy it in containerized environments. It is understood that there are multiple ways to achieve these tasks and there is no one-size-fits all approach.

## How to Use This Repository

- Clone this repository and add it to your Github profile (do not fork)
- Make changes, add files to complete the following tasks
  - Include command examples in-line in this README under each task.
  - commit and push changes
- When completed, send cloned repository link

### Task 1

Run the [Ruby](https://www.ruby-lang.org) app locally using [Sinatra](http://www.sinatrarb.com), [Rack](http://rack.github.io) or [Rails](http://rubyonrails.org) and demonstrate how to use the web application.

### Task 2

Run the app locally within a [Docker](https://docs.docker.com/engine/) container.

### Task 3

Run the app using [docker-compose](https://docs.docker.com/compose/) with [nginx](https://docs.docker.com/compose/) as a reverse proxy to the application.

### Task 4

Setup [TravisCI](https://travis-ci.org), [CircleCI](https://circleci.com), or other automated continous integration to build a container and test on push to Github repository.

### Task 5

Add deployment files for [Rancher](http://rancher.com), [Kubernetes](https://kubernetes.io) or other container [orchestration and management framework](https://github.com/cncf/landscape).

### Bonus

Additional tasks to further show mastery of application deployments and understanding.

- Add an additional route to the application to return full list of gifs
- Minimize container image
- Run application on platform other than x86
- Rewrite application in another language
