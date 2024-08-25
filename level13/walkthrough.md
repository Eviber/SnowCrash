There is a single binary executable named level13. Running it prints "UID 2013 started us but we we expect 4242".

As always, start by disassembling it with Ghidra. We can see that the program checks if our UID is equal to 4242 as it said.
We can also see that if the UID is 4242, a function called `ft_des` will be called with this string as parameter: `"boe]!ai0FB@.:|L6l@A?>qJ}I"`.

Looking at `ft_des` disassembly shows a lot of complex operations, but the most important part is that it doesn't call anything except `strdup` on the parameter at the beginning. Since it looks like this function is deterministic and doesn't mess with our UID again, let's try to run it to see what we get.

- Launch gdb on the binary, set a breakpoint on main and go to it, and see the disassembly of main.
```
gdb ./level13
b main
r
disassembly
```
- Locate the call to `getuid` and step to just after it with `ni`.
- Print the return of `getuid` (which is stored in the `eax` register) and change it to 4242.
```
p $eax
set $eax=4242
p $eax
```
- Continue to run the program that will now call `ft_des` with `c`.
- The flag is printed.

Well, that was surprisingly easy.
