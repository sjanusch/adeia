## Adeia

[![Go Report Card](https://goreportcard.com/badge/github.com/seibert-media/adeia)](https://goreportcard.com/report/github.com/seibert-media/adeia)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/513590eff4e54095a25b66bf65bd1323)](https://www.codacy.com/app/kwiesmueller/adeia?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=seibert-media/adeia&amp;utm_campaign=Badge_Grade)
[![Build Status](https://travis-ci.org/seibert-media/adeia.svg?branch=master)](https://travis-ci.org/seibert-media/adeia)
[![Docker Repository on Quay](https://quay.io/repository/seibertmedia/adeia/status "Docker Repository on Quay")](https://quay.io/repository/seibertmedia/adeia)

Adeia is an application to deploy ingresses to a kubernetes cluster depending on a list of domains provided by an http get call

## Name

Adeia is originating from the Greek word `άδεια` meaning permit in English.

## Requirements

To run the tool use our [Makefile](Makefile) 

## Dependencies
All dependencies inside this project are being managed by [dep](https://github.com/golang/dep) and are checked in.
After pulling the repository, it should not be required to do any further preparations aside from `make deps` to prepare the dev tools.

If new dependencies get added while coding, make sure to add them using `dep ensure --add "importpath"`.

## Testing
To run tests you can use:
```bash
make test
# or
go test
# or
ginkgo -r
```

## Contributing

This application is developed using behavior driven development. 
Please keep in mind to use this development method when contribution to the project.

If you are new to BDD have a look at this video which inspired this project to use BDD:
 
https://www.youtube.com/watch?v=uFXfTXSSt4I

Feedback and contributions are highly welcome. Feel free to file issues or pull requests.

## Run example

```bash
go get -u github.com/bborbe/server/bin/file_server

file_server \
-v=4 \
-logtostderr \
-root=files \
-port=1337

curl http://localhost:1337/example.json

adeia \
-v=4 \
-logtostderr \
-namespace=test \
-service-name=test \
-service-port=test \
-url=http://localhost:1337/example.json
```


## Attributions

* [Kolide for providing `kit`](https://github.com/kolide/kit)
