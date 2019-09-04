#!/usr/bin/env bash
title="ApiHub后端服务接口"
api swagger -c "luohao" -d "ApiHub系统后端服务接口" -H "10.10.63.17:18000" -n ${title}

# post swagger spec
body=$(python3 -c "import json;f=open('./swagger.json', 'r');content=f.read();reqBody = {'title': '${title}','category_id': 1,'spec_content': content};strBody=json.dumps(reqBody);print(strBody);f.close()")
curl -X POST \
  http://10.10.63.17:18000/api/api_hub/v1/doc/doc/check_post \
  -H 'Content-Type: application/json' \
  -d "${body}"
