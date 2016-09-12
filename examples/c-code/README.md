# c-code

TODO: hook in the make template to this example

C code, compiled and run through elsy.

The key things this example illustrates are:

- How you can use elsy to compile c code without installing gcc or make.
- A customized blackbox-test phase that simply executes a make step to verify
the code is functioning.

## CI

Simply run `lc ci`

## Local Development
```
## bootstrap repo
$ lc bootstrap

## compile your code
lc package

## test your code:
lc blackbox-test
```
