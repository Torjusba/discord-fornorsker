import pandas as pd
import numpy as np


class PDParser:
    
    def fix_norwegian_letters(self, _word):
        #æ
        ae = "Ã¦"
        #ø
        oe = "Ã¸"
        #å
        aa="Ã¥"
        return str(_word).replace(ae,"æ").replace(oe,'ø').replace(aa,'å')

    def __init__(self, _url):
        df = pd.read_html(_url,header=0)[0]
        ddf = df.to_dict()

        self.importlist = list(ddf['Importord'].values())
        self.avloeyserlist = list(ddf['AvlÃ¸serord'].values())
        
        for w in self.avloeyserlist:
            self.avloeyserlist[self.avloeyserlist.index(w)] = self.fix_norwegian_letters(w)
    
    
    