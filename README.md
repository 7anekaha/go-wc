# CCWC - Coding Challenge WC
## Description
Simple go version of wc unix command that counts the number of bytes, word, lines and characters in a file.

It supports a file as input or a stream from stdin.

Challenge from [here](https://codingchallenges.fyi/challenges/challenge-wc/)

*Flags supported:*
|    Flag     |    Description   |      
| ----------- |:----------------:|
|    -c       | byte count       |
|    -m       | character count  |
|    -w       | word count       |
|    -l       | newline count    |

## Usage
First you need to build the binary:
```bash
make build
```
Then you can run the binary:
```bash
.\bin\ccwc.exe test.txt
```
## Examples
```bash
> .\bin\ccwc.exe -l test.txt
   7145 test.txt
```
```bash
> Get-Content test.txt | .\bin\ccwc.exe -l
   7145
```

