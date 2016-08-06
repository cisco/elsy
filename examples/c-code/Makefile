.RECIPEPREFIX = >

clean:
>  rm foo || true

foo: foo.c
> make clean
> gcc -o foo foo.c

blackbox-test:
> ./foo | grep -s "foo and such" >/dev/null
