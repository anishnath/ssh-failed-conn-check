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

Output:
```bash
❯ ./ssh-failed-connections  -host=72.20.6.27 -username=ubuntu
Apr 28 18:45:55 host5722-host5722-vie-0 sshd[1365789]: Connection closed by authenticating user root 41.111.227.75 port 47306 [preauth]
Apr 28 18:45:55 host5722-host5722-vie-0 sshd[1365792]: Connection closed by authenticating user ubuntu 41.111.227.75 port 47356 [preauth]
Apr 28 18:45:55 host5722-host5722-vie-0 sshd[1365796]: Connection closed by authenticating user root 41.111.227.75 port 47322 [preauth]
Apr 28 18:45:56 host5722-host5722-vie-0 sshd[1365816]: Connection closed by authenticating user root 41.111.227.75 port 47318 [preauth]
Apr 28 18:55:04 host5722-host5722-vie-0 sshd[1366175]: Connection closed by authenticating user ubuntu 103.208.70.204 port 33977 [preauth]
Apr 28 18:57:31 host5722-host5722-vie-0 sshd[1366230]: Connection closed by authenticating user sol 103.208.70.204 port 39383 [preauth]
Apr 28 19:03:25 host5722-host5722-vie-0 sshd[1366346]: Connection closed by authenticating user root 101.43.1.7 port 56838 [preauth]
Apr 28 19:03:26 host5722-host5722-vie-0 sshd[1366362]: Connection closed by authenticating user ubuntu 101.43.1.7 port 56806 [preauth]
Apr 28 19:35:06 host5722-host5722-vie-0 sshd[1366916]: Connection closed by authenticating user ubuntu 103.208.70.204 port 45954 [preauth]
Apr 28 19:35:28 host5722-host5722-vie-0 sshd[1366926]: Connection closed by authenticating user ubuntu 103.208.70.204 port 54116 [preauth]
Apr 28 19:35:37 host5722-host5722-vie-0 sshd[1366929]: Connection closed by authenticating user ubuntu 103.208.70.204 port 43460 [preauth]
Apr 28 19:35:58 host5722-host5722-vie-0 sshd[1366937]: Connection closed by authenticating user ubuntu 103.208.70.204 port 36053 [preauth]
Apr 28 19:52:28 host5722-host5722-vie-0 sshd[1367539]: Connection closed by authenticating user nobody 5.10.250.44 port 54006 [preauth]
Apr 28 20:35:55 host5722-host5722-vie-0 sshd[1369919]: Connection closed by authenticating user ubuntu 103.208.70.204 port 7207 [preauth]
Apr 28 20:35:57 host5722-host5722-vie-0 sshd[1369921]: Connection closed by authenticating user ubuntu 103.208.70.204 port 65517 [preauth]
Apr 28 20:37:17 host5722-host5722-vie-0 sshd[1370047]: Connection closed by authenticating user ubuntu 103.208.70.204 port 40957 [preauth]
Apr 28 20:37:27 host5722-host5722-vie-0 sshd[1370051]: Connection closed by authenticating user ubuntu 103.208.70.204 port 11866 [preauth]
Apr 28 20:42:03 host5722-host5722-vie-0 sshd[1370331]: Connection closed by authenticating user root 14.36.134.124 port 41764 [preauth]
Apr 28 20:44:10 host5722-host5722-vie-0 sshd[1370459]: Connection closed by authenticating user ubuntu 103.208.70.204 port 4340 [preauth]

````

Running locally 
```bash
./ssh-failed-connections 
```



## License

This program is licensed under the MIT License. See the LICENSE file for details.