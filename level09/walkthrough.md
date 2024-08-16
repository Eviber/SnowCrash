There are two files:
- `level09`, an executable binary
- `token`, a file that contains some unreadable data

After some testing, it appears that `level09` encrypts the argument it gets. The encryption is pretty simple, it just adds the index of each character to its value.

- Write a simple script to reverse the encryption, it just needs to remove the index of every byte to its value
- Give the contents of `token` to the decryption script
- That's it!

⛄
