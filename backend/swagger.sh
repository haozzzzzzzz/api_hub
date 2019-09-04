#!/usr/bin/env bash
contact=$1
title=$2
description=$3
host=$4

if [[
  -z ${contact} ||
  -z ${title} ||
  -z ${description} ||
  -z ${host}
]]
then
  echo "lack of params: swagger.sh <contact> <title> <description> <host>"
  exit
fi

/data/luohao/go1.12.9/gopath/bin/api swagger -c "${contact}" -n "${title}" -d "${description}" -H "${host}"

# post swagger spec
body=$(python3 -c "import json;f=open('./swagger.json', 'r');content=f.read();reqBody = {'title': '${title}','category_id': 1,'spec_content': content};strBody=json.dumps(reqBody);print(strBody);f.close()")
curl -X POST \
  http://10.10.63.17:18000/api/api_hub/v1/doc/doc/check_post \
  -H 'Content-Type: application/json' \
  -d "${body}"
