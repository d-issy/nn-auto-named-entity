import json
import multiprocessing
import pickle
from collections import defaultdict
from concurrent.futures import ThreadPoolExecutor, as_completed

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


def result(page, i=None):
    global net
    if i is not None:
        print(i, page.title)
    vec = wiki2vec(page)
    preds = net.predict(vec, 0.9)

    nes = []
    for p in preds:
        d = data.namedEntities[p+1].jpName
        nes.append(d)
    return nes


page2ne = {}
ne2page = defaultdict(lambda: [])

# get result
with Network(v) as net:
    L = 3000
    step = 0
    a, b = [], list(pages.keys())
    while len(b) > 0:
        a, b = b[:L], b[L:]
        print('step', step)
        with ThreadPoolExecutor(max_workers=multiprocessing.cpu_count()) as executor:
            worker_to_page = {executor.submit(
                result, pages[t], step*L+i): t for i, t in enumerate(a)}
            for page in as_completed(worker_to_page):
                title = worker_to_page[page]
                try:
                    nes = page.result()
                except Exception as e:
                    print('error:', e)
                else:
                    if len(nes) > 0:
                        page2ne[title] = nes
                        for n in nes:
                            ne2page[n] = title
            executor.shutdown(wait=True)
        print()
        step += 1

ne2page = dict(ne2page)

with open('data/result.json', 'rb') as f:
    json.dump({'page2ne': page2ne, 'ne2page': ne2page}, f)
