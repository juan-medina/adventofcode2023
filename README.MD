# Advent of Code 2023

Day exercises for [Advent of Code 2023](https://adventofcode.com/2023) done in [Go](https://go.dev/)

## Run via makefile

### specific day
```sh
make run DAY="1" PART="2"
```

### solution test
```sh
make run DAY="1" PART="2" TEST="true"
```

## Run the compiled binary

First build the binary with

```sh
make build
```

### specific day
```sh
./build/adventofcode2023 -day 1 -part 1
```

### solution test
```sh
./build/adventofcode2023 -day 1 -part 1 -test
```

## Additional make commands

- clean : clean the build
- format : format the code
- vet : basic checks
- update : update dependencies & remove unused

## VS Code

Include a `launch.json` that allow to debug in VS Code, change the parameters for different days or test

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/internal/app",
            "args": ["-day", "1", "-part", "1", "-test"],
            "cwd": "${workspaceFolder}"
        }
    ]
}
```

## LICENSE
```
 Copyright (c) 2023 Juan Antonio Medina Iglesias
 
 Permission is hereby granted, free of charge, to any person obtaining a copy
 of this software and associated documentation files (the "Software"), to deal
 in the Software without restriction, including without limitation the rights
 to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 copies of the Software, and to permit persons to whom the Software is
 furnished to do so, subject to the following conditions:
 
 The above copyright notice and this permission notice shall be included in
 all copies or substantial portions of the Software.
 
 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 THE SOFTWARE.
```