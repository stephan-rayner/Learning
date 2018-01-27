
OUTFILE:="./m"

clean:
	if [ -a ${OUTFILE} ]; then rm ${OUTFILE}; rm ./data/*; fi

build:
	go build -v -o ${OUTFILE} main.go

up: clean build
	${OUTFILE}
