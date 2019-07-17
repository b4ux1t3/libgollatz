# libgollatz

`libgollatz` is a library for running the [Collatz conjecture algorithm](https://en.wikipedia.org/wiki/Collatz_conjecture). That's literally it. I'm using this to cement my understanding of packaging libraries in Go. I'm only putting it on GitHub because I'm creating another [project](https://github.com/b4ux1t3/go-rest-api) that uses this as a dependancy.

# FAQ
Q. Why are you using the `Result` struct instead of mapping an input to an array or list of outputs?

A. Because we might eventually want to store more information than just the number of steps and the outputs.
