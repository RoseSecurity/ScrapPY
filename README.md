# :dog2: ScrapPY: PDF Scraping Made Easy

<p align="center">
<img width=40% height=40% src="https://user-images.githubusercontent.com/72598486/200046477-94c17a93-2dc8-418b-96eb-2b554227dce2.png">
</p>

ScrapPY is a Python utility for scraping manuals, documents, and other sensitive PDFs to generate wordlists that can be utilized by offensive security tools to perform brute force, forced browsing, and dictionary attacks against targets. ScrapPY performs word frequency analysis and can run in full output modes to craft custom wordlists for targeted attacks. The tool dives deep to discover keywords and phrases leading to potential passwords or hidden directories, outputting to a text file that is readable by tools such as Hydra, Dirb, and Nmap. Expedite initial access, vulnerability discovery, and lateral movement with ScrapPY!

# Demo:

https://user-images.githubusercontent.com/72598486/201235531-6b037daf-d1f3-4d33-b256-8411e3a0b3da.mov

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
$ pip3 install Counter
```

ScrapPY Usage:

```
usage: ScrapPY.py [-h] [-f FILE] [-m {word-frequency,full}] [-o OUTPUT]
```

```
# Output top 100 frequently used keywords to a file name ```Top_100_Keywords.txt```

$ python3 ScrapPY.py -f example.pdf -m word-frequency -o Top_100_Keywords.txt

# Output all keywords to default ScrapPY.txt file

$ python3 ScrapPY.py -f example.pdf
```

ScrapPY Output:

```
# ScrapPY outputs the ScrapPY.txt file or specified name file to the directory in which the tool was ran. To view the first fifty lines of the file, run this command:

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
URL_BASE: http://192.168.1.123/
WORDLIST_FILES: /root/ScrapPY/ScrapPY.txt

-----------------

GENERATED WORDS: 4592

---- Scanning URL: http://192.168.1.123/ ----
==> DIRECTORY: http://192.168.1.123/vi/
+ http://192.168.1.123/programming (CODE:200|SIZE:2726)
+ http://192.168.1.123/s7-logic/ (CODE:403|SIZE:1122)
==> DIRECTORY: http://192.168.1.123/config/
==> DIRECTORY: http://192.168.1.123/docs/
==> DIRECTORY: http://192.168.1.123/external/
```

Utilize ScrapPY with Hydra for advanced brute force attacks:

```
root@RoseSecurity:~# hydra -l root -P /root/ScrapPY/ScrapPY.txt -t 6 ssh://192.168.1.123
Hydra v7.6 (c)2013 by van Hauser/THC & David Maciejak - for legal purposes only

Hydra (http://www.thc.org/thc-hydra) starting at 2014-05-19 07:53:33
[DATA] 6 tasks, 1 server, 1003 login tries (l:1/p:1003), ~167 tries per task
[DATA] attacking service ssh on port 22
```

Enhance Nmap scripts with ScrapPY wordlists:

```
nmap -p445 --script smb-brute.nse --script-args userdb=users.txt,passdb=ScrapPY.txt 192.168.1.123
```

## Future Development:

- [x] Allow for custom output file naming and increased verbosity
- [x] Integrate different modes of operation including word frequency analysis
- [ ] Incorporate ```pyexiftool``` for metadata analysis
