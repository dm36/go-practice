def isAnagram(x, y):
    print sorted(x)
    print sorted(y)
    return ''.join(sorted(x)) == ''.join(sorted(y))


print isAnagram("dog", "god")
print isAnagram("swag", "swags")
