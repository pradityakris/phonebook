# phonebook
This project is creating phonebook REST API with in memory database
Use the beego framework package and their plugins inside

# Instalation and Library used
1. Install Go language [Golang Documentation](https://golang.org/)
2. Add Beego Framework [Beego (Beego Framework)](https://github.com/astaxie/beego)
3. Add Beego Tools [Bee (Beego Tools)](https://github.com/beego/bee)

# Steps to Run The Project
1. Clone this project inside "**GOPATH/src**"
2. Run command `cd GOPATH/src`
3. Run command `cd bee run`
4. Wait until message `http server Running on http://:9090` show
5. Access http://localhost:9090/v1/phonebook
6. The dialog for username and password will show and can use username : `astaxie` password : `helloBeego`
7. The Json Result for get all phonebook show

# Features
1. Authorization basic authentication username : `astaxie` and password : `helloBeego`
2. [GET] [Get All Phonebook] http://localhost:9090/v1/phonebook result : show all phonebook
3. [GET] [Get Single Data] http://localhost:9090/v1/phonebook/hjkhsbnmn123 result : show phonebook with id `hjkhsbnmn123`
4. [DELETE] [Delete Data] http://localhost:9090/v1/phonebook/hjkhsbnmn123 result : delete phonebook with id `hjkhsbnmn123`
5. [POST] [Add New Data] http://localhost:9090/v1/phonebook with json body request `{"name" : "Cicil", "address" : "Plaza Kuningan", "phoneNumber" : "8912390"}` result : add new phonebook
6. [PUT] [Update New Data] http://localhost:9090/v1/phonebook/hjkhsbnmn123 with json body request `{"name" : "Cicil Update"}` result phonebook with id : `hjkhsbnmn123` updated with new name : `Cicil Update`
# Note 
Other explanations will be explained for each function
