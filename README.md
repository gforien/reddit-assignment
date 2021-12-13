# reddit-assignment
Go + Helm + Terraform + AWS assignment found on r/devops


## 1. Microservice en Go

On utilise [gin](https://github.com/gin-gonic/gin) pour créer un serveur http simplement
et rapidement.

On utilise [go-redis](https://github.com/go-redis/redis) pour s'interfacer avec Redis.

Notre serveur a 4 routes:
- `GET /ok` → renvoie `'gin OK'`, et le code `HTTP 200`
- `GET /count` → renvoie le résultat de la commande `GET count`, et le code `HTTP 200`
- `POST /inc`  → renvoie le résultat de la commande `INCR count`, et le code `HTTP 200`
- `POST /dec`  → renvoie le résultat de la commande `DECR count`, et le code `HTTP 200`

*Notre API web expose des commandes Redis.
`$count` n'est donc pas une variable en mémoire, mais une clé dans le cache Redis.*

### Tests
On utilise [pester](https://github.com/pester/Pester) pour écrire rapidement des tests
d'intégration en Powershell. Docker est également nécessaire pour lancer le container Redis.

Lancer les tests
```powershell
Invoke-Pester -Output Detailed
```

**(!) Pour l'instant le container Redis et le microservice Go ne sont pas dans le même network
(pour reprendre la terminologie Docker). Nos tests fonctionnent uniquement car le port
local 6379 est exposé, ce qui ne se ferait pas forcément en production.**


## 2. Containerization

### Dockerfile
- On a le choix entre les images `golang:1.17-buster` (Debian, 300 MB) et `golang:1.16-alpine`
  (Alpine, 100 MB), et pas d'image slim-buster. On choisit la version Alpine.

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

# tests
docker run --rm --name redis-test --net reas redis redis-cli -h redis ping
# → PONG
docker run --rm --name reas-test --net reas curlimages/curl curl -s reas:5000/ok
# → gin OK
docker run --rm --name reas-test --net reas curlimages/curl curl -s reas:5000/inc -XPOST
# → 1
docker run --rm --name reas-test --net reas curlimages/curl curl -s reas:5000/inc -XPOST
# → 2
```