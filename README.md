## Kubernetes Ingress

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

