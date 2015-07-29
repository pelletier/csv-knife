# csv-knife

Command line tool to manipulate CSV files.

## Installation

### From source

```bash
go install github.com/pelletier/csv-knife
```


## Usage

### Full help

    Usage of ./csv-knife:
        -d="": keep all but the given comma separated columns
        -k="": remove all but the given comma separated columns
        -rc="\"": comment delimiter for the input stream
        -rd=",": field delimiter for the input stream
        -rf=0: number of fields per record (0 = all equal to the first row, -1 = no check)
        -rl=false: allow quotes not to be closed
        -rt=true: trim leading spaces in fields
        -wc=false: use CRLF as a new line character in the output stream
        -wd=",": field delimiter for the output stream

## License

The MIT License (MIT)

Copyright (c) 2015 Thomas Pelletier

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
