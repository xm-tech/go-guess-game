APP := guess-game

VERSION := $(shell git rev-list HEAD | head -1)
AUTHOR := $(shell git log --pretty=format:"%an"|head -n 1)
BUILD_INFO := $(shell git log --pretty=format:"%s" | head -1)
BUILD_DATE := $(shell date +%Y-%m-%d\ %H:%M:%S)

# link flags
LD_FLAGS='-X "$(APP)/version.VERSION=$(VERSION)" -X "$(APP)/version.AUTHOR=$(AUTHOR)" -X "$(APP)/version.BUILD_INFO=$(BUILD_INFO)" -X "$(APP)/version.BUILD_DATE=$(BUILD_DATE)"'

OBJECT_FILE := $(APP)

$(OBJECT_FILE):
	go build -ldflags $(LD_FLAGS) -gcflags "-N" -v -o $(OBJECT_FILE) cmd/main.go

run: $(OBJECT_FILE)
	./$(OBJECT_FILE)

.PHNOY : clean
clean: 
	-rm -f $(OBJECT_FILE)
