When we log in level05, the sentence "you have new mail" appeared.
- We looked for the mail (Youva checked his mailbox hahahahaha), and we found it, level05 in `/var/mail`.
`*/2 * * * * su -c "sh /usr/sbin/openarenaserver" - flag05`
- We noticed that the format of the file looked like crontab, a tool which allows to execute something (for example a script) regularly (in this case every 2 minutes).
- Then, we looked for the script `openarenaserver` and we saw it executed and deleted all scripts in `/opt/openarenaserver/`.
- So, we just wrote a simple script:
```
#!/bin/bash 

getflag > /tmp/filename.txt
chmod +r /tmp/filename.txt
```
- We just had to wait 2 minutes and get the flag in `/tmp/filename.txt`.

Period.
