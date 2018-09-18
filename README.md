# GoLang Test Project

## Run the project

```bash
> git clone https://github.com/dantecoding/go-test-project.git
> cd go-test-project
> docker-compose -f docker-compose.yml up
```
Connect to Database with parameters
```
server=localhost
user=root
port=3306
password=password
```
Install database dump `go_test.sql`

Run at [localhost:8000](http://localhost:8000 "localhost:8000")

## API Event

#### Create Event

**POST** */event/*

**Request**
```json
{
	"name": "event name",
	"published": false
}
```
**Response**
```json
{
    "id": 1,
    "name": "event name",
    "published": false
}
```

#### Get Event

**GET** */event/{id}*

**Response**
```json
{
    "id": 1,
    "name": "event name",
    "published": false
}
```

#### Delete Event

**DELETE** */event/{id}*

**Response**
```json
{
    "result": "seccess"
}
```

#### Publish Event

**POST** */event/publish/{id}*

**Request**
```json
{
    "published": true
}
```

**Response**
```json
{
    "id": 1,
    "name": "event name",
    "published": true
}
```

#### Update Event

**PUT** */event/{id}*

**Request**
```json
{
    "id": 1,
    "name": "event rename",
    "published": true
}
```

**Response**
```json
{
    "id": 1,
    "name": "event rename",
    "published": true
}
```

## API Listener

#### Create Listener

**POST** */listener/*

**Request**
```json
{
    "event_id": 1
	"name": "listener name",
	"address": "http://server/hanlde"
}
```
**Response**
```json
{
    "id": 1,
    "event_id": 1
	"name": "listener name",
	"address": "http://server/hanlde"
}
```

#### Get Listener

**GET** */listener/{id}*

**Response**
```json
{
    "id": 1,
    "event_id": 1
    "name": "listener name",
    "address": "http://server/hanlde"
}
```

#### Delete Listener

**DELETE** */listener/{id}*

**Response**
```json
{
    "result": "seccess"
}
```

#### Update Listener

**PUT** */listener/{id}*

**Request**
```json
{
    "event_id": 1
    "name": "listener rename",
    "address": "http://server/hanlde"
}
```

**Response**
```json
{
    "id": 1,
    "event_id": 1
    "name": "listener rename",
    "address": "http://server/hanlde"
}
```