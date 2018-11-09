import multiprocessing
import pickle
from concurrent.futures import ThreadPoolExecutor, as_completed

import numpy as np

import protobuf.nn_auto_pb2 as pb
from mylib.wiki2vec import wiki2vec

# load cnv table
with open('data/nn/cnv_table.bin', 'rb') as f:
    cnv_table = pickle.load(f)

title2id = cnv_table['title2id']
id2title = cnv_table['id2title']

# load wiki
wiki = pb.Wiki()

with open('data/wiki.bin', 'rb') as f:
    wiki.ParseFromString(f.read())


pages = wiki.pages
pre_vector = {}

with ThreadPoolExecutor(max_workers=multiprocessing.cpu_count()) as executor:
    worker_to_page = {executor.submit(
        wiki2vec, pages[t], i): t for i, t in enumerate(title2id.keys())}
    for page in as_completed(worker_to_page):
        title = worker_to_page[page]
        try:
            data = page.result()
        except Exception as e:
            print('error:', e)
        else:
            pre_vector[title] = data
    executor.shutdown()

pL = len(title2id)
vector = np.zeros((pL, 672))
for title in pre_vector.keys():
    i = title2id[title]
    vector[i] = pre_vector[title]

np.save('data/nn/vector.npy', vector)
