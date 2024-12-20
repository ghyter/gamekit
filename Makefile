COMMANDFOLDER=./cmd
BINFOLDER=./bin
CMD=${BINFOLDER}/gamekit
MAIN=${COMMANDFOLDER}/gamekit.go

run: build
	${CMD}

build: 
	go build -o ${CMD} ${MAIN}

clean:
	rm -rf ${BINFOLDER}