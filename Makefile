.PHONY: all clean
all: $(addsuffix /go-ut60g, 386 amd64 arm)
	
%/go-ut60g: go-ut60g.go
	GOARCH=$(@D) go build -ldflags -s -o $@ $<

clean:
	rm 386/*
	rm amd64/*
	rm arm/*
