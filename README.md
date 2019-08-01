# ginpractice

git clone https://github.com/gnavgire/ginpractice.git

#Download the vendor utility:

go get -u github.com/kardianos/govendor

cd govendor

go build

#Copy the govendor binary to the executable $PATH

cd ginpractice

govendor init

#Add git to the vendor list of your project

govendor fetch github.com/gin-gonic/gin@v1.3

#Build your webserver

gobuild

#Execute the webserver

./ginpractive

#Test the webserver

curl -X GET "localhost:8080/cars/sports?bhp=800"

curl -X POST "localhost:8080/status"
