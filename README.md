# AsiaYo Currency Service

In this service, it will provide APIs for the client to do the following things:

* Currency Convertor

<details>
    <summary>
        <code>GET</code> <code><b>/api/v1/currency?source={source}&target={target}&amount={amount}</b></code>
    </summary>

### Parameters

| name   | type     | data type | description                    |
|--------|----------|-----------|--------------------------------|
| source | required | string    | The specific source currency   |
| target | required | string    | The specific target currency   |
| amount | required | string    | The amount of want to exchange |

### Response

| http code | content-type              | response                               |
|-----------|---------------------------|----------------------------------------|
| `200`     | `application/json`        | `{msg: "success", "amount": {amount}}` |
| `400`     | `application/json`        | `{"msg":"failed", "amount": 0}`        |

### Example
```bash
curl -X GET -H "Content-Type: application/json" http://localhost:3030/api/v1/currency\?source\=USD\&target\=JPY\&amount\=\$1,525
```
</details>

## Setup local development

* Docker desktop
* Golang

## How to run

* Local
    * Run service

    ```bash
    make start
    ```

    * Run test
    ```bash
    make test
    ```

* Docker
    * Build image
  ```bash
  docker build --no-cache -t asia-yo-curr-svc .
  ```
    * Run container
  ```bash
  docker run -it --rm --name tmp -p 3030 \
  -e SERVER_HOST=0.0.0.0 \
  asia-yo-curr-svc
  ```
