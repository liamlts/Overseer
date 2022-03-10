# Overseer
This is a program I wrote for myself therefore it assumes you have password authentication diabled for OpenSSH and are using SSH keys. Eventually it will have an
option for those who use password auth. At the moment it simply relays information to the user about possible malicious hosts, my goal is to add
firewall functionality by default without user input. 

To run Overseer currently(3/9/22) simply use go build logMonitor.go, then sudo ./logMonitor. 

Please note this is far from a finished project.
