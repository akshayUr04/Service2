# Service2
## Follow these steps to Run and Test API Gateway Service locally

#### Clone The Repository of API Gateway

git clone https://github.com/akshayUr04/Service2.git

#### Change Directory To API Gateway

cd ./microservice-2


#### Install tThe dependencies

go mod tidy 


#### Setup The Env
.env
PORT="port that you wan't to run this Service"
USER_SERVICE_URL="Port that the User Service running"


#### Run The Application

go run ./cmd/main.go