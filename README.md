# Go's Monkey Interpreter

This repository contains implementation of Monkey Interpreter using Go.

## Project Structure

Go has a special project structure, your projects should be located within its own directory inside an `src` directory, which in turn located inside your working space directory. See examples below:

    projects
      go
        bin
        pkg
        src
          <project-1>
          <project-1>
          <project-1>

Afterwards, you have to set the `GOPATH` environment variable so that go can determine your root working space directory.

    export GOPATH="$GOPATH:/mnt/c/opt/projects/go"
