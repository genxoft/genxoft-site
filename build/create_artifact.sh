#!/usr/bin/env sh

temp_dir="./genxoft.dev"

if [ -f "$temp_dir" ] ; then
    rm "$temp_dir"
fi

mkdir ${temp_dir}
cp ./genxoft-server ${temp_dir}/genxoft-server

mkdir ${temp_dir}/web
cp -R ./web/* ${temp_dir}/web

mkdir ${temp_dir}/migrations
cp -R ./migrations/* ${temp_dir}/migrations


mkdir ${temp_dir}/data
mkdir ${temp_dir}/data/files

rm genxoft-site_${1}.tar.gz

tar -czvf genxoft-site_${1}.tar.gz ${temp_dir}/*

rm -rf ${temp_dir}