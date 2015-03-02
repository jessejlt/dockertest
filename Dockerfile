FROM google/golang

# Enable UnitTests
ADD go-test /bin/go-test

# Building the app
ADD . /gopath/src/app
ENV GOPATH=/gopath/src/app/vendor
RUN mkdir /gopath/bin
WORKDIR /gopath/src/app
RUN go build .
RUN cp app /gopath/bin/app

EXPOSE 8080
CMD []
ENTRYPOINT ["/gopath/bin/app"]