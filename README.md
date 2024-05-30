# go-short-url

This small application helps us generate short URLs from long URLs.
Make sure you have go installed on your machine with version `go1.21.6`.

## Build

To build the application, run:

```bash
make build
```

## Run

To run the application locally, use:

```bash
make run-local
```

## Test

To run the tests, execute:

```bash
make test
```

### How to use?

To use this application, first start the service locally by running:

```bash
make run-local
```

Then, use the following endpoint:

#### Shorten the URL

```bash
Endpoint: http://localhost:8080/v1/urlshorter
```

Make a POST request to the above endpoint with the following payload:

```bash
{
    "url": "www.example.com"
}
```

The request will return a shortened URL like this:

```bash
http://localhost:8080/v1/urlshorter/22f930
```

### Retrive origanal URL

Paste the received shortened URL into any browser to be redirected to the original URL.

### Metrics

For getting the top 3 domain names:

```bash
http://localhost:8080/v1/urlshorter/metrics
```
