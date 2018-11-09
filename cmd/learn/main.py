import os
import pickle

import numpy as np
from sklearn.model_selection import StratifiedKFold

from mylib import Network

os.environ['TF_CPP_MIN_LOG_LEVEL'] = '2'

# load cnv_table
with open('data/nn/cnv_table.bin', 'rb') as f:
    cnv_table = pickle.load(f)
    title2id = cnv_table['title2id']
    id2title = cnv_table['id2title']
    labels = np.array(cnv_table['labels'])
    N_LABEL = labels.shape[1]

# load vector
vector = np.load('data/nn/vector.npy')

# cross validation
skf = StratifiedKFold(n_splits=5, shuffle=True)
cross_labels = []
for y in labels:
    try:
        label = np.random.choice(np.where(y)[0])
    except:
        label = np.random.choice(N_LABEL)
    cross_labels.append(label)


# train
nn_data = []
vid = 1
for train_idx, test_idx in skf.split(vector, cross_labels):
    x_train = vector[train_idx]
    y_train = labels[train_idx]
    x_test = vector[test_idx]
    y_test = labels[test_idx]
    with Network(n_out=N_LABEL) as net:
        print('validate:', vid)
        vid += 1
        net.train(x_train, y_train, x_test, y_test, iters=500)
        data = net.get_data()
        nn_data.append(data)
        print()

with open('data/nn/nn_data.bin', 'wb') as f:
    pickle.dump(nn_data, f)
