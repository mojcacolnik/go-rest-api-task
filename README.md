# go-rest-api
Basic go api with users chi router and gorm ORM and sqlite DB.

## Quick Run Project
First clone the repo then go to go-rest-api folder.

```
git clone https://github.com/mojcacolnik/go-rest-api-task.git

cd go-rest-api-task
```
### To test the endpoints

Build and run the app
```
go build

./go-rest-api-task
```
Then call the endpoints

```
curl http://localhost:9000/api/users

curl http://localhost:9000/api/users/1

curl -d '{"firstname":"Post", "lastname":"User", "email":"user@email.io", "is_active":true}' -H "Content-Type: application/json" -X POST http://localhost:9000/api/users

curl -X DELETE http://localhost:9000/api/users/1
```
