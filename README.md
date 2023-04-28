# ssh-failed-connections-checker

This is a command-line program that parses the SSH connection log and extracts all the usernames that have failed authentication for the past N days. The program can connect to a remote SSH server or run on the localhost, depending on the command-line arguments.

## Motivation

As a sysadmin task's some times it's needed which user is failing most authentication regularly

To find this we can use  `journalctl` here is the complete command which can tail the journald log's for `ssh.service`

`journalctl -f  _SYSTEMD_UNIT=ssh.service`

The log's will tell which user's is failed the authenictaion, **Note** this log's shows only the user exists in the Linux system but failed the authentcaion
```
....
Apr 28 18:55:04 host5722-host5722-vie-0 sshd[1366175]: Connection closed by authenticating user ubuntu 103.208.70.204 port 33977 [preauth]
....
Apr 28 18:57:31 host5722-host5722-vie-0 sshd[1366230]: Connection closed by authenticating user sol 103.208.70.204 port 39383 [preauth]
```

In case we need to find out who all user failed the authentication for last `N` day's a simple go lang script can be written to achieve that

## Prerequisites

- Go programming language installed  
- Access to the SSH private key for the user on the remote server

## Usage
To use the program, run the following command:

```bash
./ssh-failed-connections -host=192.168.1.100 -days=10 -username=ec2-user
```

Running locally 
```bash
./ssh-failed-connections 
```



## License

This program is licensed under the MIT License. See the LICENSE file for details.