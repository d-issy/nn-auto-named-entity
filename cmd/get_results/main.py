import pickle
from collections import defaultdict

import numpy as np

import protobuf.nn_auto_pb2 as pb
from mylib import Network
from mylib.wiki2vec import wiki2vec

# load data
data = pb.Data()
with open('data/data.bin', 'rb') as f:
    data.ParseFromString(f.read())

wiki = pb.Wiki()
with open('data/wiki.bin', 'rb') as f:
    wiki.ParseFromString(f.read())

pages = wiki.pages


# load nn data
with open('data/nn/nn_data.bin', 'rb') as f:
    nn_data = pickle.load(f)

    max_vid = 0
    max_auc = 0.0

    for vid in range(5):
        idx = nn_data[vid]['max_idx']
        auc = nn_data[vid]['aucs'][idx]
        if auc > max_auc:
            max_vid = vid
            max_auc = auc

v = {
    'w': nn_data[max_vid]['w'],
    'b': nn_data[max_vid]['b'],
}


page2ne = {}
ne2page = defaultdict(lambda: [])

with Network(v) as net:
    for i, title in enumerate(pages.keys()):
        print(i, title)
        page = pages[title]
        vec = wiki2vec(page)
        preds = net.predict(vec, 0.9)

        nes = []
        for p in preds:
            d = data.namedEntities[p+1].jpName
            ne2page[d].append(page.title)
            nes.append(d)
        if len(nes) > 0:
            page2ne[page.title] = nes
