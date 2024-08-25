We arrive in this final level to find absolutely nothing.

After searching in vain for a while, we decide to directly attack `getflag`.
After analysis, it turns out that after a few integrity checks, `getflag` gets the user's UID and prints the corresponding message.
The catch being that the flags are not stored in clear, but are decrypted on the fly by a function called `ft_des`, like in the level 13.

Two solutions arise:
- Directly attack getflag with GDB and bypass the integrity tests, and then manually change the return of `getuid` to the UID of the target (in this case flag14, with the UID 3014).
- Extract `ft_des` from `getflag` with Ghidra, and create a simple program with it to decrypt the flag.

The first one is fun albeit pretty long, so we provide [decrypt.c](level14/decrypt.c) for the second solution.

And that's it! Many thanks to our evaluators, we hope you had a good time. :)
