# I don't know if this works on anything other than linux
# this should be executed in root folder of the project
# a.k.a. /code-cadets-2021/lecture_3/03_project
# controller and calculator should already be running

# reinitializing databases
cd ./dbinitializer/bets
go run main.go
cd ../calcbets
go run main.go

# generating bets and event updates
cd ../../betgenerator
go run main.go
cd ../eventsettler
go run main.go