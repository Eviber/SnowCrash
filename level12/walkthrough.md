There is a perl script in the home directory. It's similar to level04, where the script is run by the apache server. There are two functions :
- Function N which prints `..` if its first parameter is equal to 1 or `.`.
- Function T takes two parameters. The function sets the first one to uppercase and removes all characters after a space. There are comparisons with the second parameter.

We see that the most important thing is a shell command that uses `egrep` with the first parameter. This hints that we can inject a command here.

The issue here is : the first parameter is set to uppercase, so it limits what we can do.
First we try to use an environment variable, to pass a script, because we see the shebang is : `#!/usr/bin/env/ perl` but it doesn't work. Indeed the script is not executed by us so the environment isn't shared.
We look for an another way to circumvent the uppercase issue. We realize that non alphabetic characters are unaffected. This means that we can use a wildcard to specify a path.

Thus, we can write a script in `/tmp/[YOURSCRIPT]` and as long as the script's name is all uppercase, we'll be able to refer to it with `/*/[YOURSCRIPT]`.  
Here is the script (don't forget to set the execution rights) :
```
#!/bin/bash

getflag > /tmp/[yourfile]
chmod +r /tmp/[youfile]
```

Now, all we need to do is to give `` `/*/[YOURSCRIPT]` ``, and because the parameter is between backticks it will be executed by the server as a bash command.
The entire command is :
```
curl 'localhost:4646?x=`/*/[YOURSCRIPT]`'
```
