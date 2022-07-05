# Eirclode

A online version of the Eircode finder. Because for some reason the Irish
government cannot afford a public API.

All of the data in this repository is available on finder.eircode.ie, and is
public because this is just postcodes.

---

## Querying

Other than the data present in the `data/eircodes.cvs`, you can use the webserver
to query the Eircodes.

### Request

```text
GET /{eircode}
Content-Type: application/json
```

### Response

```text
Status 200

{
    "eircode": "{eircode}"
    "lattitude": 53.34022,
    "longitide": -6.25492,
    "address": "National Museum of Ireland, Kildare Street, Dublin 2"
}
```

---

## Running Locally

You can run your own instance of the service.

```sh
go run ./cmd/eirclode
```

This will expose the server on default port `8080`, configurable using the `PORT` env var.
