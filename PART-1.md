# reddit-assignment ♻☸🌐
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
