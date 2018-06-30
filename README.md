# GO-TAG 
---
### Requirements 
Please make sure you have Go install and `$GOPATH` variable correctly set. For more information [here](https://golang.org/doc/install). 

### Usage 
After setting up your Go environment, run the following command:

```
go get github.com/jb-hirad/go-tag
cd $GOPATH/src/github.com/jb-hirad/go-tag
go install
```

This will download the code and install it. Please make sure you create the `.env` file in the root directory of the project. You can follow the `.env.example` provided.

You can generate a personal access token from github by following [these directions](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/).

After creating a tag for release in the desired repo, simply run `go-tag`. Go-tag will fetch the current and previous release info, look at the pull requests made, and update their corresponding JIT or CS Jira tickets. 

Please note that this has been tested on the Cloud Studio repo only.

---
### Future Work
- [ ] Support multiple repo at once 
- [ ] Take advantage of Go Concurrency  