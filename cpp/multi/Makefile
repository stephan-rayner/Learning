.DEFAULT_GOAL := main

main: build run
build:
	g++ -v -g -Wall -std=c++11 ./src/main.cpp ./src/libs/*.cpp -o ./bin/app
run:
	clear
	time ./bin/app
clean:
	rm ./bin/app | true
	rm -r ./bin/app.dSYM | true