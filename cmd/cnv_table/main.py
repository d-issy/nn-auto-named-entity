import pickle
from collections import defaultdict

import numpy as np

import protobuf.nn_auto_pb2 as pb

print('create cnv table...')

# load datum
data = pb.Data()
wiki = pb.Wiki()

with open('data/data.bin', 'rb') as f:
    data.ParseFromString(f.read())

with open('data/wiki.bin', 'rb') as f:
    wiki.ParseFromString(f.read())


# create cnv tables
title2id = defaultdict(lambda: len(title2id))
id2title = {}
labels = []

for title in data.taggedPages.keys():
    id2title[title2id[title]] = title
    page = data.taggedPages[title]
    label_on_hot = np.zeros(230)
    for a in page.attrs:
        label_on_hot[a-1] = 1
    labels.append(label_on_hot.tolist())


# save datum
title2id = dict(title2id)
cnv_table = {
    'title2id': title2id,
    'id2title': id2title,
    'labels': labels,
}

with open('data/nn/cnv_table.bin', 'wb') as f:
    pickle.dump(cnv_table, f)

print('end!!!')
