**Upcoming Elections**

# Overview

Search for upcoming elections in your area by address.

Along with go standard library packages, Upcoming Elections uses the [gorilla](https://www.gorillatoolkit.org/) web toolkit for routing, csrf tokens, and converting form data to go structs.

## Running

Building this package requires the go toolchain. Follow these [instructions](https://golang.org/doc/install#install) to install go on your system.

### Build

```
git clone https://github.com/mlauter/upcoming-elections.git
cd upcoming-elections
make
```
This will fetch the necessary dependencies and compile the elections binary.

### Run

You will need an api key for google's [Civic Information API](https://developers.google.com/civic-information/docs/using_api)

```
export CIVIC_DATA_API_KEY="<your_api_key>"
ENV=development ./elections -host 127.0.0.1 -port 8080
```

And navigate to `localhost:8080` in your browser.

## Testing

To run the tests, simply run

```
make test
```
