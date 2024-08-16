We are presented with two files:
- `level06.php` is a script that takes two arguments, and makes some substitutions on the file given as the first argument
- `level06` is a binary with a flag06 SUID bit that will pass its arguments to `level06.php`

Basically the idea is to find a way to make the php script evaluate a call to `getflag`.

- `level06.php` uses the function preg_replace with the `/e` flag at the beginning of the function x. This is a serious security issue as it means it will execute the given code after substitution.
- Create a file `/tmp/test` and write in it `[x {${exec($z)}}]`. This will match the regex with the `/e` flag and will execute the payload inside.
- Call `level06` with `/tmp/test` as a first argument and whatever function you want to be run as flag06 as a second argument. For example `./level06 /tmp/test getflag`.
- The output of the command will be read as a variable by php, which will fail, and an error will be shown with the name of the "variable", wich gives us the output!

PHP sucks btw.
