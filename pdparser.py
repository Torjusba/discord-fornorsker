import pandas as pd
import numpy as np


class PDParser:
    def __init__(self, _url):
        df = pd.read_html(_url,header=0)[0]
        ddf = df.to_dict()

        self.importlist = list(ddf['Importord'].values())
        self.avloeyserlist = list(ddf['AvlÃ¸serord'].values())

    
    