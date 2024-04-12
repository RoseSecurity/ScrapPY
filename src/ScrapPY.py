#!/usr/bin/env python3

from scipy.stats import entropy
from collections import Counter
from datetime import datetime
import pandas as pd
import argparse
import PyPDF2
import textract
import time
import re
import sys


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
parser = argparse.ArgumentParser(description='\nScrapPY enumerates documents, manuals, and sensitive PDFs for key phrases and words that can be utilized in dictionary and brute force attacks. These keywords are outputted to a text file (ScrapPY.txt in the directory which the tool was ran from)that can be read by tools such as Hydra, Dirb, and other offensive security tools for initial access and lateral movement.\t' + GREEN + tag + NORM)
parser.add_argument('-f','--file', help='PDF input file')
parser.add_argument('-m', '--mode', choices=['word-frequency', 'full', 'metadata', 'entropy'], help='Modes of operation: full - All keywords, word-frequency - 100 most frequently used keywords, metadata - Title, author, and extracted metadata, entropy - 100 most random keywords, potentially disclosing password hashes or hardcoded keys', default='full')
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
    if args.mode == "metadata":
        with open(args.file, 'rb') as pdf_file:
            read_pdf = PyPDF2.PdfFileReader(pdf_file)
            pdf_info = read_pdf.getDocumentInfo()
            author = pdf_info.author
            creator = pdf_info.creator
            producer = pdf_info.producer
            title = pdf_info.title
            subject = pdf_info.subject
            creation_date = str(pdf_info.creation_date)
            print(BLUE + "Title: " + NORM + title)
            print(BLUE + "Subject: " + NORM + subject)
            print(BLUE + "Author: " + NORM + author)
            if creator == producer:
                print(BLUE + "Creator: " + NORM + creator)
            else:
                print(BLUE + "Producer: " + NORM + producer)
                print(BLUE + "Creator: " + NORM + creator)
            print(BLUE + "Creation Date: " + NORM + creation_date + "\n")
            sys.exit(0)
    if args.mode == "entropy":
        unduped_keywords = list(dict.fromkeys(keywords))
        # Pass deduplicated keywords to byte array function
        entropy_conv(unduped_keywords)
    else:
        dedup(keywords)

# Convert dedplicated keywords to byte array and pass to entropy calculation function
def entropy_conv(unduped_keywords):
    keyword_bytearray_list = []
    for word in unduped_keywords:
        keyword_bytearray_list.append(bytearray(word, "utf-8"))
    entropy_calc(keyword_bytearray_list)
    
def entropy_calc(keyword_bytearray_list):
    # Detetmine entropy score
    entropy_score = []
    for word in keyword_bytearray_list:
        byte_series = pd.Series(word)
        entropy_score.append(float(entropy(byte_series.value_counts())))
    # Decode keyword list
    keywords_list = []
    for word in keyword_bytearray_list:
        keywords_list.append(str(word.decode("utf-8")))
    # Create dictionary of keywords and entropy scores
    score_dict = {keywords_list[i]: entropy_score[i] for i in range(len(keywords_list))}
    sorted_score_dict = Counter(dict(sorted(score_dict.items(),
                           key=lambda item: item[1],
                           reverse=True)))   
    for key,value in sorted_score_dict.most_common(100):
        print( BLUE + "Keyword: " + NORM + key + BLUE + "\t\tEntropy Score: " + NORM + str(value))

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
