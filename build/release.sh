#!/usr/bin/env sh
if [ -z "$1" ]
    BUILD_MODE=beta
  then
    BUILD_MODE=$1
fi

OUTPUT=`make mode="${BUILD_MODE}" | tail -1`

echo "Uploading... ${OUTPUT}"
scp ./${OUTPUT} ${GENXOFT_HOST}:${GENXOFT_PATH}
echo "Uploading done"
ssh ${GENXOFT_HOST} "cd ${GENXOFT_PATH}; tar -xvf ${OUTPUT}; supervisorctl restart genxoft-site-${BUILD_MODE}:"