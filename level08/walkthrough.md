We have two files : a binary file (`level08`) and `token`.  
When we execute level08 with no arguments, `./level08 [file to read]` is written. So we understand we need to give a path in arg (probably `token`).  
However, we can't read to token file at all.  
By disassembling `level08` with Ghidra we see that if the path name contains "token", the program writes we can't access to that file and exits.  

- We create a symbolic link to `token` file in /tmp.
- Then we execute level08 with /tmp/[filename] : read the file and write the content which is the token.
- Execute `su flag08` with the token as password to have the flag for level09.

That's it :) 
