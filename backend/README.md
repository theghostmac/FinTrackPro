# FinTrackPro Backend

## Usage
Use curl to test endpoints.

### Register user:
Request:
```shell
curl -X POST http://localhost:9020/register \
 -H 'Content-Type: application/json' \
 -d '{"userName": "GhostMac", "email": "theghostmac@dev.com", "password": "dev@ghost123"}'
```
Response:
```shell
{"message":"User registered successfully"}
```

### Login User:
Request:
```shell
curl -X POST http://localhost:9020/login \
 -H 'Content-Type: application/json' \
 -d '{"userName": "GhostMac", "email": "theghostmac@dev.com", "password": "dev@ghost123"}'
```
Response:
```shell
{"message":"User logged in successfully","user":{"UserID":"e1bed4ce-d156-4efc-a92e-ddaa35ddbb3b","UserName":"GhostMac","Email":"theghostmac@dev.com","PasswordHash":"$2a$10$tUmZr9/m55UELcmS5eKMxecbr1GB2FNbKyFKXRtu.inry8CYrF0V2","Transactions":null,"CreatedAt":"2024-03-10T08:20:04.236595+01:00"}}%
```