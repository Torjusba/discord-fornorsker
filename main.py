#! /usr/bin/python3

# Nettsiden:
# https://www.sprakradet.no/sprakhjelp/Skriverad/Avloeysarord/
from pdparser import PDParser


def recognize(_importlist, _avloeyserlist):
    inp = input("Ord >")
    if inp in _importlist:
        print("Gjenkjent")
        index = _importlist.index(inp)
        print("Index: {}".format(index))
        alternativ = _avloeyserlist[index]
        print(alternativ)


def main():
    pdparser = PDParser('https://www.sprakradet.no/sprakhjelp/Skriverad/Avloeysarord/')
    while True:
        recognize(pdparser.importlist, pdparser.avloeyserlist)

if __name__=="__main__":
    main()