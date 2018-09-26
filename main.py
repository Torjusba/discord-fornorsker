#! /usr/bin/python3

# Nettsiden:
# https://www.sprakradet.no/sprakhjelp/Skriverad/Avloeysarord/
from pdparser import PDParser


def recognize(_word, _importlist, _avloeyserlist):
    _alt = ""
    if _word in _importlist:
        _alt = _avloeyserlist[_importlist.index(_word)]
    return _alt


def main():
    pdparser = PDParser('https://www.sprakradet.no/sprakhjelp/Skriverad/Avloeysarord/')
    while True:
        inputstring = input("Ord >").lower()
        alt = recognize(inputstring, pdparser.importlist, pdparser.avloeyserlist)

        if len(alt)>0:
            print("Norske alternativer: \n {}".format(alt))

if __name__=="__main__":
    main()