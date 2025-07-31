#!  /usr/bin/ksh

#   build go-dmig cli
export  CURRENT_DIR="$( pwd )"

go build -o bin/go-dmig main.go
cd ${CURRENT_DIR}
