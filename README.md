## Kubernetes Ingress
![Go Report Card](https://goreportcard.com/badge/github.com/seibert-media/k8s-ingress)](https://goreportcard.com/report/github.com/seibert-media/k8s-ingress)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/513590eff4e54095a25b66bf65bd1323)](https://www.codacy.com/app/kwiesmueller/k8s-ingress?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=seibert-media/k8s-ingress&amp;utm_campaign=Badge_Grade)
[![Build Status](https://travis-ci.org/seibert-media/k8s-ingress.svg?branch=master)](https://travis-ci.org/seibert-media/k8s-ingress)
[![Docker Repository on Quay](https://quay.io/repository/seibertmedia/k8s-ingress/status "Docker Repository on Quay")](https://quay.io/repository/seibertmedia/k8s-ingress)

Kubernetes Ingress is an application to deploy ingresse to k8s cluster depending on a list of domains provide by an http get call.

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
	#or
	go test
	#or
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
