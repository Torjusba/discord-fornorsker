from parsehtml import HTMLTableParser

def main():
    hp = HTMLTableParser()
    table = hp.parse_url(url)[0][1] # Grabbing the table from the tuple
    table.head()

if __name__=="__main__":
    main()