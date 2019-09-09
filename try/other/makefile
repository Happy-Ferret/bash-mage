GO_PATH=$(shell go env GOPATH)
DEP_PATH=$(GO_PATH)/bin/dep
MAGE=go run ci/mage.go

default:
ifeq ("$(wildcard $(DEP_PATH))", "")
	go get -u github.com/golang/dep/cmd/dep
endif
	${DEP_PATH} ensure
	${MAGE} -l

% :
ifeq ("$(wildcard $(DEP_PATH))", "")
	go get -u github.com/golang/dep/cmd/dep
endif
	${DEP_PATH} ensure
	${MAGE} $(MAKECMDGOALS)