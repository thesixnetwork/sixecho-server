APINAME=$1
NETWORK=$2
ENV=$3

if [ "$APINAME" == "" ]; then
    echo "ERROR: Missing API name"
    exit 1
fi

if [ "$NETWORK" == "" ]; then
    echo "ERROR: Missing Network name"
    exit 1
fi
if [ "$ENV" == "" ]; then
    echo "ERROR: Missing Env name"
    exit 1
fi

if [ -f .$ENV.env ]; then
    . .$ENV.env
else
    echo "ERROR: ENV file not exists"
fi

_CNT_BUILD=`find . -name "${APINAME}.json" | wc -l`
if [ ${_CNT_BUILD} -eq 0 ]; then
    echo "ERROR: Could not find built contract, please compile first"
    exit 1
fi

echo "Running Migration for $APINAME on $NETWORK ($ENV) at Storage(${STORAGE_ADDR})"
STORAGE_ADDR=`echo ${STORAGE_ADDR} | awk -F "0x" '{print $2}'`
truffle migrate --network ${NETWORK} -s new_api -r ${STORAGE_ADDR} -a ${APINAME}
# echo $CMD
# APIADDR=`$CMD`
# echo "APIADDR = ${APIADDR}"
echo -e "Please enter New API address:\c"
read APIADDR
# APIADDR=`echo $APIADDR | awk '{print $1}'`

# if [ $? -ne 0 ]; then
#     echo "ERROR: Fail to migrate"
#     exit 1
# fi

echo "Deployed API address ${APIADDR}"

echo "Add API to EchoApp"

if [ "${WALLET_SECRET}" == "" ]; then 
    read WALLET_SECRET
fi

WALLET_SECRET=`echo ${WALLET_SECRET} | awk -F "0x" '{print $2}'`
APIADDR=`echo ${APIADDR} | awk -F "0x" '{print $2}'`
ECHO_APP_ADDR=`echo ${ECHO_APP_ADDR} | awk -F "0x" '{print $2}'`
CALLER=`echo ${CALLER} | awk -F "0x" '{print $2}'`


CMD=$(echo "node add_new_api.js -n ${NETWORK_URL} -w ${WALLET_SECRET} -p ${APINAME} -a ${APIADDR} -c ${CALLER} -e ${ECHO_APP_ADDR}")
echo $CMD
$CMD
