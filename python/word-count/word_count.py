import re

def count_words(sentence):
    hash = {}
    for word in prepare(sentence).split(' '):
        word = clean(word)
        if word == '':
            continue
        if word not in hash:
            hash[word] = 0
        hash[word] += 1
    return hash

def clean(word):
    return re.sub(r'^[^\w]+', '', re.sub(r'[^\w]+$', '', word))

def prepare(sentence):
    return re.sub(r'\t|_|,', ' ', sentence.lower())
