# Nettsiden:
# https://www.sprakradet.no/sprakhjelp/Skriverad/Avloeysarord/


import pandas as pd
import numpy as np



df = pd.read_html('http://www.sprakradet.no/sprakhjelp/Skriverad/Avloeysarord/',header=0)[0]

ddf = df.to_dict()
#ndf = df.values

importlist = list(ddf['Importord'].values())
avloeyserlist = list(ddf['AvlÃ¸serord'].values())

print(importlist)

while True:
    inp = input("Ord >")
    if inp in importlist:
        print("Gjenkjent")
        index = importlist.index(inp)
        print("Index: {}".format(index))
        alternativ = avloeyserlist[index]
        print(alternativ)
