import re
def abbreviate(words):
    word_array = re.split('-| |_', words)
    abb = ""
    for word in word_array:
        if len(word) > 0:
            abb += word[0]
    return abb.upper()