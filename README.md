# :closed_book: ScrapPY: PDF Scraping Made Easy

<img src="https://user-images.githubusercontent.com/72598486/200032795-29e5bb5d-db3b-4ee3-9e46-f6b732848a6f.png" width=40% height=40%>

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


