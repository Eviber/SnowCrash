Pretty easy (or we are just dumb), thanks Ghidra. 
We have a binary file `level07`.

- As usual we saw the file owner is flag07.
- We used Ghidra to decompile level07 :  
  We saw the function getenv("LOGNAME") which gets the value of LOGNAME (in this case it's level07).  
  The asprintf function is called, it writes `/bin/echo %s` (s: LOGNAME), in a string.  
  Then, there is the system() function that executes echo.  
- We just had to change LOGNAME's value to `\;getflag` in order to execute getflag as flag07 and run level07.

That's it. 
