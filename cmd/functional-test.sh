#!  /usr/bin/ksh

export  CURRENT_DIR="$( pwd )"

#   test scenatio #01
export SCENARIO="01"
export DESCRIPTION="fixed position input format"

echo "[scenario #${SCENARIO}] ${DESCRIPTION}"

cd "test/scenario${SCENARIO}"
../../bin/go-dmig config.yaml
cd ${CURRENT_DIR}

#   test scenatio #02
export SCENARIO="02"
export DESCRIPTION="CSV input format"

echo
echo "[scenario #${SCENARIO}] ${DESCRIPTION}"

cd "test/scenario${SCENARIO}"
../../bin/go-dmig config.yaml
cd ${CURRENT_DIR}
