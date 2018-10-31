# -*- coding: utf-8 -*-
from privatetoken import doUglyUtfHack
import pandas as pd
import numpy as np


class PDParser:
    
    #Fikser norske bokstaver og gjør alle bokstavene små
    def fix_word_formatting(self, _word):
        ae = "Ã¦" #æ
        oe = "Ã¸" #ø
        aa="Ã¥" #å
        uu='Ã¼' #ü
        return str(_word).replace(ae,"æ").replace(oe,'ø').replace(aa,'å').replace(uu, 'u').lower()
        
    def remove_parenthesis(self, _word):
        return _word.split('(')[0]

    def __init__(self, _url):
        df = pd.read_html(_url,header=0)[0]
        ddf = df.to_dict()
        
        if not (doUglyUtfHack()):
            _akey= 'Avl\xc3\xb8serord'
        else:
            _akey=b'Avl\xc3\xb8serord'.decode("utf-8", "strict")
        

        self.importlist = list(ddf['Importord'].values())
        self.avloeyserlist = list(ddf[_akey].values())

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
            elif _noparstr != w:
                if _noparstr[-1] == " ":
                    self.importlist[self.importlist.index(_w)] = _noparstr[:-1].lower()
                else:
                    self.importlist[self.importlist.index(_w)] = _noparstr.lower()
                    
            else:
                self.importlist[self.importlist.index(_w)] = w.lower()   