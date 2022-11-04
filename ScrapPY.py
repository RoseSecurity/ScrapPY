#!/usr/bin/env python3

import PyPDF2
import textract
import time
import re
import sys
import os

# Variables
BLUE = '\033[34m'
RED = '\033[91m'
GREEN = '\033[92m'
NORM = '\x1b[0m'
tag = "@RoseSecurity"
# Banner
print(BLUE + """

 ▄▀▀▀▀▄  ▄▀▄▄▄▄   ▄▀▀▄▀▀▀▄  ▄▀▀█▄   ▄▀▀▄▀▀▀▄  ▄▀▀▄▀▀▀▄  ▄▀▀▄ ▀▀▄ 
█ █   ▐ █ █    ▌ █   █   █ ▐ ▄▀ ▀▄ █   █   █ █   █   █ █   ▀▄ ▄▀ 
   ▀▄   ▐ █      ▐  █▀▀█▀    █▄▄▄█ ▐  █▀▀▀▀  ▐  █▀▀▀▀  ▐     █   
▀▄   █    █       ▄▀    █   ▄▀   █    █         █            █   
 █▀▀▀    ▄▀▄▄▄▄▀ █     █   █   ▄▀   ▄▀        ▄▀           ▄▀    
 ▐      █     ▐  ▐     ▐   ▐   ▐   █         █             █     
        ▐                          ▐         ▐             ▐     

""")
time.sleep(1)
print("\nScrapPY enumerates documents, manuals, and sensitive PDFs for key phrases and words that can be utilized in dictionary and brute force attacks. These keywords are outputted to a text file (ScrapPY.txt in the directory which the tool was ran from)that can be read by tools such as Hydra, Dirb, and other offensive security tools for initial access and lateral movement.\n\nUsage: python3 ScrapPY.py test_doc.pdf\n\n" + GREEN + tag + NORM)

# Evaluate sysargv
if len(sys.argv) > 1:
    filename = sys.argv[1]
else:
    print(RED + "\nEnter PDF file to scrape\n" + NORM)
# Open provided file
pdf_file = open(filename, 'rb')
read_pdf = PyPDF2.PdfFileReader(pdf_file)
page_nums = int(read_pdf.numPages)

# Loop through each page and convert to text
loop_count = 0
page_text = " "

while loop_count < page_nums:
    page_obj = read_pdf.getPage(loop_count)
    loop_count += 1
    page_text += page_obj.extractText()

# Lowercase each word
page_text = str(page_text.encode('ascii','ignore').lower())

# Extracting keywords
keywords = re.findall(r'[a-zA-Z]\w+', page_text)

# Deduplicate keywords
keywords = list(dict.fromkeys(keywords))

# Remove common words
common_words = [ "and", "the", "at", "there", "some", "my", "of", "be", "use", "her", "than", "and", "this", "an", "would", "first", "a", "have", "each",   "to", "from", "which", "like", "been", "in", "or", "she", "him",  "is", "one", "do", "into", "who", "you", "had", "how", "that", "by", "their", "has", "its", "it", "if", "he", "but", "was", "not", "up", "more", "for", "are", "were", "as", "we", "with", "when", "then", "no", "come", "his", "your", "them", "way", "they", "can", "these", "could", "may", "I", "said", "so" ]

for word in list(keywords):
    if word in common_words:
        keywords.remove(word)

# Create output file
with open("ScrapPY.txt", "w+") as file:
    for word in (keywords):
        file.write('%s\n' %word)
file.close()
