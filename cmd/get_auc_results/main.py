import json
import multiprocessing
import pickle
import sys
from collections import defaultdict
from concurrent.futures import ThreadPoolExecutor, as_completed

import numpy as np

import protobuf.nn_auto_pb2 as pb
from mylib import Network
from mylib.wiki2vec import wiki2vec

if len(sys.argv) != 2:
    print('NID is required')
    sys.exit(-1)
try:
    NE_ID = int(sys.argv[1]) - 1
except ValueError:
    print('The format of NID is incorrect')
    sys.exit(-1)

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

# laod vector
vector = np.load('data/nn/vector.npy')
with open('data/nn/cnv_table.bin', 'rb') as f:
    cnv_table = pickle.load(f)
    title2id = cnv_table['title2id']
    id2title = cnv_table['id2title']
    labels = np.array(cnv_table['labels'])
    N_LABEL = labels.shape

check_ids = [np.where(r == 1)[0][0] for r in labels]
check_ids = np.array(check_ids)

with Network(v) as net:
    auc = net.get_auc(
        vector[np.where(check_ids == NE_ID)[0]],
        labels[np.where(check_ids == NE_ID)[0]],
    )
    print(len(np.where(check_ids == NE_ID)[0]), auc)
