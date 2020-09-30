# Webhook

## Main Endpoint

The webhook adapter exposes a default endpoint `POST /` that returns a JSON
payload.

```bash
rackup -s puma -p 4567
```

Sample call:

```bash
curl -X POST http://localhost:4567/ -d '{ "spam": "eggs" }'
```

Sample response:

```json
{
  "meta": {
    "timestamp": "2020-09-23T15:59:57+10:00",
    "request": { "spam": "eggs" }
  },
  "data": { "hello": "world" }
}
```

The webhook accepts `POST` requests on `/` and returns:

* `200`: OK
* `500`: Internal Server Error
* `422`: Unprocessable Entity

Return payload is a JSON object which contains `meta` and `data` as "roots".

## Healthcheck

The webhook adapter also comes with an healthcheck endpoint `GET /` that simply
return `200 - OK`.
