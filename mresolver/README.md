# multiNamResolver

Reads the strandard a list of domain names from standard input end displays a list of ip address for each domain names

## Examples

```bash
$ mresolver -h
Usage of ./mresolver:
  -u    Show unresolved names


$ mresolver < domains.txt
google.com [142.251.132.46 2800:3f0:4001:834::200e]
www.google.com [172.217.30.4 2800:3f0:4001:83a::2004]


$ mresolver -u < domains.txt
google.com [142.251.132.46 2800:3f0:4001:834::200e]
www.google.com [172.217.30.4 2800:3f0:4001:83a::2004]

Domains unable to resolve:  [invalid.google.com]


$ subfinder -d nmap.org | ./mresolver

               __    _____           __
   _______  __/ /_  / __(_)___  ____/ /__  _____
  / ___/ / / / __ \/ /_/ / __ \/ __  / _ \/ ___/
 (__  ) /_/ / /_/ / __/ / / / / /_/ /  __/ /
/____/\__,_/_.___/_/ /_/_/ /_/\__,_/\___/_/

                projectdiscovery.io

[INF] Current subfinder version v2.6.6 (latest)
[INF] Loading provider config from /home/rodolfo/.config/subfinder/provider-config.yaml
[INF] Enumerating subdomains for nmap.org
[INF] Found 44 subdomains for nmap.org in 3 seconds 971 milliseconds
snanme.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
echo.nmap.org [45.33.32.156 2600:3c01::f03c:91ff:fe18:bb2f]
ascanme.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
changeme.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
analizame2.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
svn.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
scanme.nmap.org [45.33.32.156 2600:3c01::f03c:91ff:fe18:bb2f]
sname.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
scamne.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
nmap-v-ascanme.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
scnme.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
chat.nmap.org [45.33.32.156 2600:3c01::f03c:91ff:fe18:bb2f]
research.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
2fscanme.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
scame.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
scanne.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
0mhvmwpuvzrorj.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
imap4.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
Scannme.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
httpscame.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
issues.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
scannet.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
wwww.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
scan.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
mail.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
httpscanme.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
i.vpn.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
scname.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
nmap-scanme.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
scanmen.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
Scanned.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
csanme.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
ack.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
www.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
www.scanme2.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
httpscanm.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
scanmev6.nmap.org [2600:3c01::f03c:91ff:fe18:bb2f]
scamer.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
ftp0.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
scheme.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
scannme.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
httpscanned.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
www.servicioeasy.nmap.org [45.33.49.119 2600:3c01:e000:3e6::6d4e:7061]
```
