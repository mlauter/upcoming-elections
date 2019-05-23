**Democracy Works Upcoming Elections Practical Exercise**


# Overview

This is a [go](https://golang.org/) implementation of the Democracy works upcoming elections practical excercise. Along with go standard library packages, it uses the [gorilla](https://www.gorillatoolkit.org/) web toolkit for routing, csrf tokens, and converting form data to go structs.

## Running

Building this package requires the go toolchain. Follow these [instructions](https://golang.org/doc/install#install) to install go on your system. (Just in case, I've also included the compiled elections binary for OSX, which should allow you to run the elections app on a mac without rebuilding.)

Go is a bit specific about directory structure.

```
mkdir -p ${GOPATH}/src/github.com/mlauter
cp -r /path/to/unzipped/dir ${GOPATH}/src/github.com/mlauter
``` 

### Build

From inside the unzipped directory, first run:

```
make
```
This will fetch the necessary dependencies and compile the elections binary. 

### Run

```
ENV=development ./elections -host 127.0.0.1 -port 8080
```

And navigate to `localhost:8080` in your browser.

## Testing

To run the tests, simply run

```
make test
```

## Future Improvements

* Better address validation (e.g. usps validation/standardization)
* Client side validation on the address form to match for better UX
* Better frontend in general, especially one that presents more of the upcoming election data in a way that is readable and easy to understand.
* Templates + clearer messaging for user facing errors
* Tune timeout settings (especially for turbovote api requests)
* Investigate if there are missing edge cases in the response format returned by the turbovote api that the current struct won't match.
