#!/usr/bin/env sh

BUILD_MODE=${1}

TEMP_DIR="./genxoft-site-${BUILD_MODE}"

ARTIFACT_NAME="genxoft-site_${2}-${1}.tar.gz"

if [ -f "$TEMP_DIR" ] ; then
    rm "$TEMP_DIR"
fi

mkdir ${TEMP_DIR}
cp ./genxoft-server ${TEMP_DIR}/genxoft-server

mkdir ${TEMP_DIR}/web
cp -R ./web/* ${TEMP_DIR}/web

mkdir ${TEMP_DIR}/migrations
cp -R ./migrations/* ${TEMP_DIR}/migrations


mkdir ${TEMP_DIR}/data
mkdir ${TEMP_DIR}/data/files

rm ${ARTIFACT_NAME}

tar -czvf ${ARTIFACT_NAME} ${TEMP_DIR}/*

rm -rf ${TEMP_DIR}

echo "Created artifact ${VERSION}-${BUILD_MODE} build: ${RELEASE_ID}"
echo "${ARTIFACT_NAME}"