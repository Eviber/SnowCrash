There is only one file, level11.lua, which already is running. It is a simple server that asks for a password, compare its SHA1 hash with a predetermined hash, and respond with a message depending on if the hash is correct.

After following the red herring of the hash, we realize that since the server is directly concatenating the password in a shell command, an injection is very easy to do.

- Connect to the server with netcat: `nc localhost 5151`
- The server asks for a password, write `; getflag > /tmp/flag`
- Read the created file to get the flag `cat /tmp/flag`
- And voilà !

The hash was matching the string "NotSoEasy" by the way...
