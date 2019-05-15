def is_isogram(string):
    s = list(string.upper().replace(" ","").replace("-",""))
    while len(s) > 0:
        first = s[0]
        del s[0]
        if first in s:
            return False
    return True
