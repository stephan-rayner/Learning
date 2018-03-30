.DEFAULT_GOAL := main

main: build run
build:
	g++ -v -g -Wall -std=c++11 ./src/main.cpp ./src/libs/*.cpp -o app
run:
	clear
	./app
clean:
	rm ./app | true
	rm -r app.dSYM | true