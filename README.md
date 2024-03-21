# poker

This creates a randomly generated Texas Hold'em poker game and prints out the results.

Currently is only the backend, frontend eventually coming.  Can use as a basic http server to create a game or a command line to print out a sample game.  Here are some uses:

`make run` builds a binary in the `backend` directory and runs the server.  Runs on localhost port 8088.  Has 2 query parameters: `hands` as a number (> 1 or <= 10) and `print` as boolean - which will print to console log a pretty version of the game to more easily visualize the json returned (defaults to false)

`make build` builds the binary but does not run it (i.e. creates `poker` binary in `backend` directory).

`./poker deal` is cli version of the server.  It randomly generates a game and prints the results.  Check `-h` flag to see usable parameters

`make test` runs unit tests for dealer part of application (main part of logic for application, responsible for generating hands and evaulating results)

`make clean` deletes the binary built with either `make run` or `make build` 