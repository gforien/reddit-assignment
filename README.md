# reddit-assignment
Go + Helm + Terraform + AWS assignment found on r/devops


## 1. Microservice en Go

On utilise [gin](https://github.com/gin-gonic/gin) pour créer un serveur http simplement
et rapidement.

On lance un container Docker pour tester le client Redis.
```powershell
docker run --rm --name redis -p 6379:6379 -d redis
```

**(!) Pour l'instant le container Redis et le client Go ne sont pas dans le même network
(pour reprendre la terminologie Docker). Cela fonctionne uniquement car le port local 6379
est exposé, ce qui ne se ferait pas en production.**