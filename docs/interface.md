# interface

## user (mel)
* create
* input: username
* output: uuid
* view
* input: uuid
* output: username
* search
* input: query
* output: uuids

### stub
* uuid, username

## auth (mel)
* register
* input: username, password
* output: uuid
* login
* input: username, password
* output: uuid
* verify
* input: cookie
* output: uuid
* logout
* input: uuid
* output:

### stub
* uuid, hash(password), status

## follow (matheus)
* follow
* input: uuid1, uuid2
* output:
* unfollow
* input: uuid1, uuid2
* output:
* view
* input: uuid
* output: uuids

### stub
* uuid, uuids

## post (matheus)
* create
* input: uuid, data
* output: puid
* view
* input: puid
* output: data
* list
* input: uuid
* output: puids

### stub
* uuid, puid, data

## feed (miguel)
* view
* input: uuid
* output: puids

### stub
* uuid, puids
