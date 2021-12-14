# reddit-assignment ♻☸🌐
## 2. Containerization

### Dockerfile
Images
- Pour l'images Go on a le choix entre les images `golang:1.17-buster` (Debian, 300 MB)
  et `golang:1.16-alpine` (Alpine, 100 MB), et pas d'image slim-buster. On choisit la 
  version Alpine.
- Pour l'image Redis, on choisit également la dernière version mineure basée sur Alpine.

Nos containers vont maintenant s'exécuter sur des hôtes différents, au lieu d'être sur
`localhost`. Il faut donc modifier le code Go pour qu'il puisse faire une résolution de
nom vers le container Redis.

Par défaut avec Docker, le hostname d'un container dans un network correspond au nom
attribué au container lors de sa création. Ex:
```powershell
docker network create net1
docker run --rm --net net1 --name mon_redis_1 -p 6379:6379 -d redis

# hostame → mon_redis_1:6379
```

On peut lancer et tester les container manuellement
```powershell
docker network create reas

docker run --rm --name redis --net reas -p 6379:6379 -d redis
docker build -t gforien/reas .
docker run --rm --name reas --net reas -p 5000:5000 -d gforien/reas
```

Après l'avoir lancé, on peut le tester ainsi
```
docker run --rm --name redis-test --net reas redis redis-cli -h redis ping
# → PONG
docker run --rm --name reas-test --net reas curlimages/curl curl -s reas:5000/ok
# → gin OK
docker run --rm --name reas-test --net reas curlimages/curl curl -s reas:5000/inc -XPOST
# → 1
docker run --rm --name reas-test --net reas curlimages/curl curl -s reas:5000/inc -XPOST
# → 2
```

### Docker-compose

On lance le cluster
```powershell
docker-compose build
docker-compose up
```

Après l'avoir lancé, on peut le tester ainsi
```
docker network ls
# → reddit-assignment_default
docker run --rm --net reddit-assignment_default curlimages/curl curl -s -XPOST reas:5000/inc
# → 1
docker run --rm --net reddit-assignment_default curlimages/curl curl -s -XPOST reas:5000/inc
# → 2
```

### Docker Swarm

On peut également déployer nos containers sur Docker Swarm, qui est un orchestrateur plus
complet que Docker-compose, mais plus simple à prendre en main que Kubernetes.

```powershell
docker swarm init
docker stack deploy -c docker-compose.yml reas
```

Après l'avoir lancé, on peut le tester ainsi
```
docker network ls
# → reas_default
docker service create `
    --name test-swarm `
    --network reas_default `
    -d curlimages/curl curl -s -XPOST reas:5000/inc

docker service logs -f test-swarm
# → 1
# → 2
# → 3
# → 4
```
*On lance ici un service, c'est-à-dire une tâche qui sera déployée sur tous les noeuds du
cluster Swarm. Avec Docker Swarm on ne lance pas de containers indépendants, on lance au
minimum un service qui sera déployé sur un seul noeud.*
