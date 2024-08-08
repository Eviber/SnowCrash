There's a file named level02.pcap, it contains a snapshot of some captured packets.

- scp the file out of the VM
- launch WireShark on it, and open level2.pcap with it
- analysis reveals that the captured packets are a login attempt
- extract the password, taking care of the `7f` bytes aka DEL
