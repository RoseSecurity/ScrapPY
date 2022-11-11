#!/usr/bin/env python3

from collections import Counter
import argparse
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
  ______                         ______  _     _  
 / _____)                       (_____ \| |   | |
( (____   ____  ____ _____ ____  _____) ) |___| |
 \____ \ / ___)/ ___|____ |  _ \|  ____/|_____  |
 _____) | (___| |   / ___ | |_| | |      _____| |
(______/ \____)_|   \_____|  __/|_|     (_______|
                          |_|                                                  
""" + NORM)
time.sleep(1)

# Arguments
parser = argparse.ArgumentParser(description='\nScrapPY enumerates documents, manuals, and sensitive PDFs for key phrases and words that can be utilized in dictionary and brute force attacks. These keywords are outputted to a default ScrapPY.txt file, or specified name with the --output flag, in the directory which the tool was ran from. This file can be read by tools such as Hydra, Dirb, and other offensive security tools for initial access and lateral movement.\t' + GREEN + tag + NORM)
parser.add_argument('-f','--file', help='PDF input file')           
parser.add_argument('-m', '--mode', choices=['word-frequency', 'full'], help='Modes of operation: full - All keywords, word-frequency - 100 most frequently used keywords', default='full')      
parser.add_argument('-o', '--output', help='Name of file to output')
args = parser.parse_args()

def read_file():
    # Open provided file
    global common_words
    pdf_file = open(args.file, 'rb')
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

    # Remove common words
    common_words = [ "and", "the", "at", "there", "some", "my", "of", "be", "use", "her", "than", "and", "this", "an", "would", "first", "a", "have", "each",   "to", "from", "which", "like", "been", "in", "or", "she", "him",  "is", "one", "do", "into", "who", "you", "had", "how", "that", "by", "their", "has", "its", "it", "if", "he", "but", "was", "not", "up", "more", "for", "are", "were", "as", "we", "with", "when", "then", "no", "come", "his", "your", "them", "way", "they", "can", "these", "could", "may", "I", "said", "so" ]

    for word in list(keywords):
        if word in common_words:
            keywords.remove(word)
    return keywords

def dedup(keywords):
    # Deduplicate keywords
    unduped_keywords = list(dict.fromkeys(keywords))
    keywords = unduped_keywords
    output_file(unduped_keywords)    

def mode(keywords):
    if args.mode == "word-frequency":
        keyword_count = Counter(keywords)
        common_keywords = []
        for word, count in keyword_count.most_common(100):
            common_keywords.append(word)
        keywords = common_keywords
        output_file(keywords)
    else:
        dedup(keywords)

def output_file(keywords):
    # Create output file    
    if args.output:
        with open(args.output, "w+") as file:
            for word in (keywords):
                file.write('%s\n' %word)
        file.close()
        print(args.output + " has been created!")
    else:
        with open("ScrapPY.txt", "w+") as file:
            for word in (keywords):
                file.write('%s\n' %word)
        file.close()
        print("ScrapPY.txt has been created!")

def main():
# Evaluate arguments
    if args.file:
        read_keywords = read_file()
        mode_keywords = mode(read_keywords)
    else:
        print(RED + "\nEnter PDF file to scrape or use -h for help menu\n" + NORM)

main()
