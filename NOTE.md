# note


keygen from the website

./nokubectl --as admin orch conncheck

./nokubectl --as admin orch add cl --clusterid test-cs

 - ./nokubectl read

 - 149220e323812e51201998cd90e5e998

./nokubectl --as admin orch install cl --clusterid test-cs --targetip 192.168.50.94:22 --targetid ubuntu --targetpw ubuntu --localip 192.168.50.94 --osnm ubuntu --cv 1.27 --updatetoken 62920f8ad52b85034374c36f82a8208c


./nokubectl --as admin orch install cl log --clusterid test-cs --targetip 192.168.50.94:22 --targetid ubuntu --targetpw ubuntu 


```bash
total 8
drwxr-xr-x  2 root root 4096 Nov 14 01:41 .
drwxr-xr-x 20 root root 4096 Nov 14 01:41 ..

----------ROOT NPIA CREATED----------
[sudo] password for ubuntu: 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
100 19.7M  100 19.7M    0     0  7841k      0  0:00:02  0:00:02 --:--:-- 9258k

----------NPIA BIN DOWNLOADED----------
[sudo] password for ubuntu: 

----------NPIA BIN INSTALLED----------
[sudo] password for ubuntu: 
existing configuration: 
MODE: test
BASE_URL: http://localhost:13337
EMAIL: seantywork@gmail.com
-----INITLOG-----
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
100 18.5M  100 18.5M    0     0  8209k      0  0:00:02  0:00:02 --:--:-- 10.5M
lib/
lib/base/
lib/base/hacfg
lib/base/admininstanctrlol-ubuntu20
lib/base/admininstenvres-ubuntu20
lib/base/admininit-ubuntu20
lib/base/settingcrtvol-ubuntu20
lib/base/admininstwkol-ubuntu20
lib/base/admininstvolol-ubuntu20
lib/base/admininsttkol-ubuntu20
lib/base/admininstctrl-ubuntu20
lib/base/admininstenv-ubuntu20
lib/base/settingcrtmon-ubuntu20
lib/bin/
lib/bin/docker-compose
lib/bin/kompose
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1699  100  1699    0     0   5041      0 --:--:-- --:--:-- --:--:--  5041
Reading package lists...
Building dependency tree...
Reading state information...
The following NEW packages will be installed:
  apt-transport-https
0 upgraded, 1 newly installed, 0 to remove and 46 not upgraded.
Need to get 1704 B of archives.
After this operation, 162 kB of additional disk space will be used.
Get:1 http://kr.archive.ubuntu.com/ubuntu focal-updates/universe amd64 apt-transport-https all 2.0.9 [1704 B]
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "en_US.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
locale: Cannot set LC_CTYPE to default locale: No such file or directory
locale: Cannot set LC_MESSAGES to default locale: No such file or directory
locale: Cannot set LC_ALL to default locale: No such file or directory
Fetched 1704 B in 0s (3485 B/s)
Selecting previously unselected package apt-transport-https.
(Reading database ... 72314 files and directories currently installed.)
Preparing to unpack .../apt-transport-https_2.0.9_all.deb ...
Unpacking apt-transport-https (2.0.9) ...
Setting up apt-transport-https (2.0.9) ...
deb [arch=amd64 signed-by=/usr/share/keyrings/helm.gpg] https://baltocdn.com/helm/stable/debian/ all main
Hit:1 http://kr.archive.ubuntu.com/ubuntu focal InRelease
Get:2 https://baltocdn.com/helm/stable/debian all InRelease [7651 B]
Get:3 https://baltocdn.com/helm/stable/debian all/main amd64 Packages [3852 B]
Get:4 http://kr.archive.ubuntu.com/ubuntu focal-updates InRelease [114 kB]
Get:5 http://kr.archive.ubuntu.com/ubuntu focal-backports InRelease [108 kB]
Get:6 http://kr.archive.ubuntu.com/ubuntu focal-security InRelease [114 kB]
Get:7 http://kr.archive.ubuntu.com/ubuntu focal-updates/main amd64 Packages [2959 kB]
Get:8 http://kr.archive.ubuntu.com/ubuntu focal-updates/main Translation-en [481 kB]
Get:9 http://kr.archive.ubuntu.com/ubuntu focal-updates/main amd64 c-n-f Metadata [17.2 kB]
Get:10 http://kr.archive.ubuntu.com/ubuntu focal-updates/restricted amd64 Packages [2479 kB]
Get:11 http://kr.archive.ubuntu.com/ubuntu focal-updates/restricted Translation-en [346 kB]
Get:12 http://kr.archive.ubuntu.com/ubuntu focal-updates/restricted amd64 c-n-f Metadata [552 B]
Get:13 http://kr.archive.ubuntu.com/ubuntu focal-updates/universe amd64 Packages [1129 kB]
Get:14 http://kr.archive.ubuntu.com/ubuntu focal-updates/universe Translation-en [270 kB]
Get:15 http://kr.archive.ubuntu.com/ubuntu focal-updates/universe amd64 c-n-f Metadata [25.7 kB]
Get:16 http://kr.archive.ubuntu.com/ubuntu focal-updates/multiverse amd64 Packages [25.8 kB]
Get:17 http://kr.archive.ubuntu.com/ubuntu focal-updates/multiverse amd64 c-n-f Metadata [620 B]
Get:18 http://kr.archive.ubuntu.com/ubuntu focal-security/main amd64 Packages [2569 kB]
Get:19 http://kr.archive.ubuntu.com/ubuntu focal-security/main Translation-en [398 kB]
Get:20 http://kr.archive.ubuntu.com/ubuntu focal-security/main amd64 c-n-f Metadata [13.2 kB]
Get:21 http://kr.archive.ubuntu.com/ubuntu focal-security/restricted amd64 Packages [2367 kB]
Get:22 http://kr.archive.ubuntu.com/ubuntu focal-security/restricted Translation-en [329 kB]
Get:23 http://kr.archive.ubuntu.com/ubuntu focal-security/restricted amd64 c-n-f Metadata [552 B]
Get:24 http://kr.archive.ubuntu.com/ubuntu focal-security/universe amd64 Packages [898 kB]
Get:25 http://kr.archive.ubuntu.com/ubuntu focal-security/universe Translation-en [188 kB]
Get:26 http://kr.archive.ubuntu.com/ubuntu focal-security/universe amd64 c-n-f Metadata [19.2 kB]
Get:27 http://kr.archive.ubuntu.com/ubuntu focal-security/multiverse amd64 Packages [23.6 kB]
Get:28 http://kr.archive.ubuntu.com/ubuntu focal-security/multiverse amd64 c-n-f Metadata [548 B]
Fetched 14.9 MB in 9s (1569 kB/s)
Reading package lists...
Reading package lists...
Building dependency tree...
Reading state information...
The following NEW packages will be installed:
  helm
0 upgraded, 1 newly installed, 0 to remove and 80 not upgraded.
Need to get 16.0 MB of archives.
After this operation, 50.6 MB of additional disk space will be used.
Get:1 https://baltocdn.com/helm/stable/debian all/main amd64 helm amd64 3.12.1-1 [16.0 MB]
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "en_US.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
locale: Cannot set LC_CTYPE to default locale: No such file or directory
locale: Cannot set LC_MESSAGES to default locale: No such file or directory
locale: Cannot set LC_ALL to default locale: No such file or directory
Fetched 16.0 MB in 2s (6468 kB/s)
Selecting previously unselected package helm.
(Reading database ... 72318 files and directories currently installed.)
Preparing to unpack .../helm_3.12.1-1_amd64.deb ...
Unpacking helm (3.12.1-1) ...
Setting up helm (3.12.1-1) ...
Processing triggers for man-db (2.9.1-1) ...
npia init success

-----------------
successfully initiated

----------NPIA INITIATED----------


----------**********----------

---------*   LOG    *---------

----------**********----------

[sudo] password for ubuntu: 
----------MESSAGE----------
+ HOME=/root
+ sudo swapoff -a
+ crontab -l
+ crontab -
+ echo '@reboot /sbin/swapoff -a'
+ sudo apt-get update -y
Hit:1 http://kr.archive.ubuntu.com/ubuntu focal InRelease
Hit:2 https://baltocdn.com/helm/stable/debian all InRelease
Hit:3 http://kr.archive.ubuntu.com/ubuntu focal-updates InRelease
Hit:4 http://kr.archive.ubuntu.com/ubuntu focal-backports InRelease
Hit:5 http://kr.archive.ubuntu.com/ubuntu focal-security InRelease
Reading package lists...
+ OS=xUbuntu_20.04
+ OSNM=ubuntu
+ VERSION=1.27
+ cat
+ sudo tee /etc/modules-load.d/crio.conf
overlay
br_netfilter
+ sudo modprobe overlay
+ sudo modprobe br_netfilter
+ cat
+ sudo tee /etc/sysctl.d/99-kubernetes-cri.conf
net.bridge.bridge-nf-call-iptables  = 1
net.ipv4.ip_forward                 = 1
net.bridge.bridge-nf-call-ip6tables = 1
+ sudo sysctl --system
* Applying /etc/sysctl.d/10-console-messages.conf ...
kernel.printk = 4 4 1 7
* Applying /etc/sysctl.d/10-ipv6-privacy.conf ...
net.ipv6.conf.all.use_tempaddr = 2
net.ipv6.conf.default.use_tempaddr = 2
* Applying /etc/sysctl.d/10-kernel-hardening.conf ...
kernel.kptr_restrict = 1
* Applying /etc/sysctl.d/10-link-restrictions.conf ...
fs.protected_hardlinks = 1
fs.protected_symlinks = 1
* Applying /etc/sysctl.d/10-magic-sysrq.conf ...
kernel.sysrq = 176
* Applying /etc/sysctl.d/10-network-security.conf ...
net.ipv4.conf.default.rp_filter = 2
net.ipv4.conf.all.rp_filter = 2
* Applying /etc/sysctl.d/10-ptrace.conf ...
kernel.yama.ptrace_scope = 1
* Applying /etc/sysctl.d/10-zeropage.conf ...
vm.mmap_min_addr = 65536
* Applying /usr/lib/sysctl.d/50-default.conf ...
net.ipv4.conf.default.promote_secondaries = 1
sysctl: setting key "net.ipv4.conf.all.promote_secondaries": Invalid argument
net.ipv4.ping_group_range = 0 2147483647
net.core.default_qdisc = fq_codel
fs.protected_regular = 1
fs.protected_fifos = 1
* Applying /usr/lib/sysctl.d/50-pid-max.conf ...
kernel.pid_max = 4194304
* Applying /etc/sysctl.d/99-kubernetes-cri.conf ...
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-ip6tables = 1
* Applying /etc/sysctl.d/99-sysctl.conf ...
* Applying /usr/lib/sysctl.d/protect-links.conf ...
fs.protected_fifos = 1
fs.protected_hardlinks = 1
fs.protected_regular = 2
fs.protected_symlinks = 1
* Applying /etc/sysctl.conf ...
+ sudo tee /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list
+ cat
deb https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/xUbuntu_20.04/ /
+ sudo tee /etc/apt/sources.list.d/devel:kubic:libcontainers:stable:cri-o:1.27.list
+ cat
deb http://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable:/cri-o:/1.27/xUbuntu_20.04/ /
+ sudo apt-key --keyring /etc/apt/trusted.gpg.d/libcontainers.gpg add -
+ curl -L https://download.opensuse.org/repositories/devel:kubic:libcontainers:stable:cri-o:1.27/xUbuntu_20.04/Release.key
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0Warning: apt-key output should not be parsed (stdout is not a terminal)
100   393  100   393    0     0    487      0 --:--:-- --:--:-- --:--:--   486
100   394  100   394    0     0    373      0  0:00:01  0:00:01 --:--:--   373
100   395  100   395    0     0    301      0  0:00:01  0:00:01 --:--:--   301
100   396  100   396    0     0    253      0  0:00:01  0:00:01 --:--:--   253
100   397  100   397    0     0    219      0  0:00:01  0:00:01 --:--:--     0
```