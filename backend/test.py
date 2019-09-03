import json;f=open('./swagger.json', 'r');content=f.read();reqBody = {"title": "xxx","category_id": 1,"spec_content": content};strBody=json.dumps(reqBody);print(strBody);f.close()
