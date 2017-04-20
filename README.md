# Go Action SDK (port of the google action java sdk)

## Warning
This is currently an extremely alpha SDK, maybe even pre-alpha, it is not ready for prime time and comes with no guarantees that it will actually work for you.  It has worked for my limited test cases, if you do use it and run into issues feel free to open an issue however it's possible I will not maintain this in the long term depending on life and other factors.

## About
This is a port of a port.  Put simply my javascript experience and knowledge is minimal at best, however [frogermcs](https://github.com/frogermcs/Google-Actions-Java-SDK) ported the action sdk google released from javascript to java.  I know java and since I wanted to work in go to make containerization easier and smaller I have ported his work to Go which is my goto language currently.

This is currently under active development and is not ready.  I will merge it into the master branch when it is ready to be used/tested/broken.  That being said I do not guarantee any functionality or that it will be kept up to date if google changes their API.

## TODO:
* Setup unit tests.
* Setup CI job in travis-ci to run the unit tests on commit.
* Test and (probably) fix the askResponse function in ResponseBuilder.go.
* Add various other handlers to abstract interacting with the API based on documentation [here.](https://developers.google.com/actions/reference/conversation)
* Add examples/sample code of how to use it.
* More to come