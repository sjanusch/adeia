## Adeia

[![Go Report Card](https://goreportcard.com/badge/github.com/seibert-media/adeia)](https://goreportcard.com/report/github.com/seibert-media/adeia)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/513590eff4e54095a25b66bf65bd1323)](https://www.codacy.com/app/kwiesmueller/adeia?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=seibert-media/adeia&amp;utm_campaign=Badge_Grade)
[![Build Status](https://travis-ci.org/seibert-media/adeia.svg?branch=master)](https://travis-ci.org/seibert-media/adeia)
[![Docker Repository on Quay](https://quay.io/repository/seibertmedia/adeia/status "Docker Repository on Quay")](https://quay.io/repository/seibertmedia/adeia)

Adeia is an application to deploy ingress rules to a Kubernetes cluster depending on a list of domains provided by an http get call

## Name

The word adeia originates from the Greek word άδεια which means 'to permit' in English.

## Installing

Simply run `make all`. There should now be a command `adeia` in your `$GOPATH`.  

## First Run

Get a basic file server and let it run in the background on port 1337: 
```bash
go get -u github.com/bborbe/server/bin/file_server

file_server \
-v=4 \
-logtostderr \
-root=files \
-port=1337
```

Put a file called `example.json` into the `files/` directory. It should contain the list of domains you want adeia to provide ingresses for. Check it's working by running:

```bash
curl http://localhost:1337/example.json
```

You should see the list again.

You can now run:

```bash
adeia \
-v=4 \
-logtostderr \
-namespace=test \
-ingress-name=test \
-service-name=test \
-service-port=test \
-url=http://localhost:1337/example.json \
-dry-run
```

to see what changes adeia would make to your Kubernetes cluster. 

If you're happy with what you see, actually apply the changes by leaving out the `-dry-run` flag: 

```bash
adeia \
-v=4 \
-logtostderr \
-namespace=test \
-ingress-name=test \
-service-name=test \
-service-port=test \
-url=http://localhost:1337/example.json
```

## Usage

Adeia provides a `-help` flag which prints usage information and exits. 

## Dependencies

All dependencies inside this project are being managed by [dep](https://github.com/golang/dep) and are checked in.
After pulling the repository, it should not be required to do any further preparations aside from `make deps` to prepare the dev tools.

When adding new dependencies while contributing, make sure to add them using `dep ensure --add "importpath"`.

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

## Attributions

* [Kolide for providing `kit`](https://github.com/kolide/kit)
