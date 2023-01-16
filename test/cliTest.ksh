#!  /usr/bin/ksh

#   build go-dmig cli
export  CURRENT_DIR="$( pwd )"

cd ..
go build -o bin/go-dmig main.go
cd ${CURRENT_DIR}

#   test scenatio #01
export SCENARIO="01"
export DESCRIPTION="fixed position input format"

echo "[scenario #${SCENARIO}] ${DESCRIPTION}"

cd "scenario${SCENARIO}"
../../bin/go-dmig config.yaml
cd ${CURRENT_DIR}
