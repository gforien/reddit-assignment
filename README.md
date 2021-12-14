# reddit-assignment ♻☸🌐

### Coding assignment trouvé sur [r/devops](https://www.reddit.com/r/devops):<br>Écrire et déployer un microservice avec Go, Helm, Terraform, AWS, Github Actions

### Table des matières
1. [Microservice en Go](PART-1.md)
    1. [Tests](PART-1.md#Tests)
1. [Containerization](PART-2.md)
    1. [Écrire le Dockerfile](PART-2.md#dockerfile)
    1. [Lancer avec Docker-compose](PART-2.md#docker-compose)
    1. [Lancer avec Docker Swarm](PART-2.md#docker-swarm)

<br>

### Sujet (légèrement modifié)
```
1. Using Go, write a simple counter application (http server) which serves 3 endpoints
    - /inc which increment the hits and save in redis
    - /dec which decrement the hits and save in redis
    - /count will present the number of hits.
    Containerize all code including the redis, and save it as docker-compose build.
    Remove the source code from the container.
    Save it on github repository

2. Build a helm chart of the httpserver you wrote.
    Explain which kind of k8s resource you pick and why?
    Add liveness and readiness
    Expose the service port to 3030
    Install kind to test you application and install the helm chart
    Bonus**: Serve a route with the commit SHA of the deployed app.

3. Build github action file that trigger build on specific wildcard PR (*build*) and
   on merge to master and create tags, which build your container, run unittest, save
   to the registry with the correct version.
   Bonus: Write a terraform which spins a virtual machine and sets relevant firewall
   rules and load balancer for our application. Pick (GCP or AWS).

4. Script:
    Write a script (shell) which acts like installer (curl …..| bash -c… look a like)
    Its will download and the httpserver as binary (you can include the httpserver above
    as binary in the shell script).	Add check that its install on linux os and ubuntu 20.04
    Bonus: add systemctl service start for the application.

Note: Please document all command in README.md format and save on the github repository
```

#### Gabriel Forien <br> 5TC INSA Lyon
![](https://upload.wikimedia.org/wikipedia/commons/b/b9/Logo_INSA_Lyon_%282014%29.svg)
