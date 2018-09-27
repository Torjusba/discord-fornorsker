import pandas as pd
import numpy as np


class PDParser:
    
    #Fikser norske bokstaver og gjør alle bokstavene små
    def fix_word_formatting(self, _word):
        ae = "Ã¦" #æ
        oe = "Ã¸" #ø
        aa="Ã¥" #å
        return str(_word).replace(ae,"æ").replace(oe,'ø').replace(aa,'å').lower()
        
    def remove_parenthesis(self, _word):
        return _word.split('(')[0]

    def __init__(self, _url):
        df = pd.read_html(_url,header=0)[0]
        ddf = df.to_dict()

        self.importlist = list(ddf['Importord'].values())
        self.avloeyserlist = list(ddf['AvlÃ¸serord'].values())

        # Fiks formatering        
        for w in self.avloeyserlist:
            self.avloeyserlist[self.avloeyserlist.index(w)] = self.fix_word_formatting(w)

        #Remove label entries
        for _w in self.importlist:
            _noparstr = self.remove_parenthesis(str(_w))

            w=str(_w) #Fikser problemer med at æøå tolkes som float
            if len(w)==1:
                self.avloeyserlist.pop(self.importlist.index(w))
                self.importlist.remove(w)
            else:
                self.importlist[self.importlist.index(_w)] = _noparstr[:-1]


    
    