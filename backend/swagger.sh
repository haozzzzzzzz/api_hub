#!/usr/bin/env bash
title="Api hub backend api"
api swagger -c "luohao" -d "Api hub backend api specication" -H "127.0.0.1:18000" -n ${title}

# post swagger spec
body=$(python3 -c "import json;f=open('./swagger.json', 'r');content=f.read();reqBody = {'title': '${title}','category_id': 1,'spec_content': content};strBody=json.dumps(reqBody);print(strBody);f.close()")
curl -X POST \
  http://127.0.0.1:18000/api/api_hub/v1/doc/doc/check_add \
  -H 'Content-Type: application/json' \
  -d "${body}"