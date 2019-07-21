# Api Hub Data Design

## rational table list

| name        | description             |
| ----------- | ----------------------- |
| ah_doc      | api document table      |
| ah_account  | api hub account table   |
| ah_category | document category table |
| ah_tag      | document tags           |
| ah_doc_tag  |                         |



### ah_doc

| field       | descritpion                   | type             | default | remark                                  |
| ----------- | ----------------------------- | ---------------- | ------- | --------------------------------------- |
| doc_id      | document auto-increasement id | UNSIGNED INT     |         | PK; Auto-Increasement                   |
| title       | doc title                     | VARCHAR(255)     |         | NOT NULL                                |
| spec        | api specification content     | TEXT             |         | NOT NULL                                |
| spec_url    | api specification url         | VARCHAR(500)     | ""      | NOT NULL                                |
| category_id | category id                   | UNSIGNED INT     | 1       | NOT NULL                                |
| author_id   | author id                     | UNSIGNED INT     |         | NOT NULL                                |
| post_status | post status                   | UNSIGNED TINYINT | 0       | 0:not published; 1: published;2:deleted |
| update_time | update time                   | UNSIGNED INT     |         | NOT NULL                                |
| create_time | create time                   | UNSIGNED INT     |         | NOT NULL                                |



### ah_account

| field       | description  | type         | default | remark                |
| ----------- | ------------ | ------------ | ------- | --------------------- |
| acc_id      | account id   | UNSIGNED INT |         | PK; Auto-Increasement |
| name        | account name | VARCHAR(32)  |         | NOT NULL              |
| update_time | update time  | UNSIGNED INT |         | NOT NULL              |
| create_time | create time  | UNSIGNED INT |         | NOT NULL              |



### ah_category

| field       | description     | type         | default | remark                |
| ----------- | --------------- | ------------ | ------- | --------------------- |
| cat_id      | category id     | UNSIGNED INT |         | PK; Auto-Increasement |
| name        | category name   | VARCHAR(32)  |         | NOT NULL              |
| doc_num     | document number | UNSIGNED INT |         | NOT NULL              |
| update_time | update time     | UNSIGNED INT |         | NOT NULL              |
| create_time | create_time     | UNSIGNED INT |         | NOT NULL              |



### ah_tag

| field       | description     | type         | default | remark                |
| ----------- | --------------- | ------------ | ------- | --------------------- |
| tag_id      | tag id          | UNSIGNED INT |         | PK; Auto-Increasement |
| name        | tag name        | VARCHAR(32)  |         | NOT NULL              |
| doc_num     | document number | UNSIGNED INT |         | NOT NULL              |
| update_time | update time     | UNSIGNED INT |         | NOT NULL              |
| create_time | create_time     | UNSIGNED INT |         | NOT NULL              |



### ah_doc_tag

| field       | description | type         | default | remark             |
| ----------- | ----------- | ------------ | ------- | ------------------ |
| tag_id      | tag id      | UNSIGNED INT |         | NOT NULL; Unoin PK |
| doc_id      | doc id      | UNSIGNED INT |         | NOT NULL; Unoin PK |
| create_time | create time | UNSIGNED INT |         | NOT NULL           |

