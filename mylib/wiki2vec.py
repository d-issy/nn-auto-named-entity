import os
import pickle

import numpy as np
import tensorflow as tf

import protobuf.nn_auto_pb2 as pb
from .get_char import get_char

# disable cuda
os.environ['TF_CPP_MIN_LOG_LEVEL'] = '2'

# element n
EN = 192

# read filters
with open('data/nn/filters.bin', 'rb') as f:
    filters = pickle.load(f)

f1 = filters['f1']
f2 = filters['f2']
f3 = filters['f3']

# load cnv table
with open('data/nn/cnv_table.bin', 'rb') as f:
    cnv_table = pickle.load(f)

title2id = cnv_table['title2id']
id2title = cnv_table['id2title']

# load statistics
stat = pb.Statistics()
with open('data/stat.bin', 'rb') as f:
    stat.ParseFromString(f.read())

# arrange stats
cat = {}
for i, v in enumerate(stat.categories[:EN]):
    cat[v.key] = {'rank': i, 'count': v.count}

lin = {}
for i, v in enumerate(stat.links[:EN]):
    lin[v.key] = {'rank': i, 'count': v.count}

tem = {}
for i, v in enumerate(stat.templates[:EN]):
    tem[v.key] = {'rank': i, 'count': v.count}
del stat


def str2bytes(st):
    out = np.zeros([0, 32])
    for ch in st[-32:]:
        c = [bin(c)[2:] for c in ch.encode()]
        c = ''.join(c).zfill(32)
        c = list(map(np.float32, c))
        c = np.array(c)
        out = np.vstack((out, c))
    out = np.pad(out, ((32-len(out), 0), (0, 0)), 'constant')
    return np.reshape(out, (1, 32, 32, 1))


def cnv_title(title_text):
    x = tf.constant(str2bytes(title_text))
    c1 = tf.nn.conv2d(x, f1, strides=(1, 1, 1, 1), padding='VALID')
    pool1 = tf.nn.max_pool(c1, ksize=(1, 30, 1, 1),
                           strides=(1, 1, 1, 1), padding='VALID')

    c2 = tf.nn.conv2d(x, f2, strides=(1, 1, 1, 1), padding='VALID')
    pool2 = tf.nn.max_pool(c2, ksize=(1, 29, 1, 1),
                           strides=(1, 1, 1, 1), padding='VALID')

    c3 = tf.nn.conv2d(x, f3, strides=(1, 1, 1, 1), padding='VALID')
    pool3 = tf.nn.max_pool(c3, ksize=(1, 28, 1, 1),
                           strides=(1, 1, 1, 1), padding='VALID')

    out = tf.concat([pool1, pool2, pool3], axis=3)
    out = tf.reshape(out, (-1, 1, 96))
    return out


def wiki2vec(i, page):
    print(i, page.title)
    #  -- title
    tv = cnv_title(page.title)
    #  -- cat
    c_all_v = np.zeros(EN)
    categories = page.categories
    for category in categories:
        cv = np.zeros(EN)
        for ch in category:
            ch = get_char(ch)
            if ch in cat:
                cv[cat[ch]['rank']] += 1
        c_all_v += np.where(cv > 1.0, 1.0, cv)
    if len(categories) > 0:
        c_all_v /= len(categories)
    #  -- links
    l_all_v = np.zeros(EN)
    links = page.links
    for link in links:
        cv = np.zeros(EN)
        for ch in link:
            ch = get_char(ch)
            if ch in lin:
                cv[lin[ch]['rank']] += 1
        l_all_v += np.where(cv > 1.0, 1.0, cv)
    if len(links) > 0:
        l_all_v /= len(links)
    #  -- links
    t_all_v = np.zeros(EN)
    templates = page.templates
    for template in templates:
        if template in tem:
            t_all_v[tem[template]['rank']] = 1
    out = np.vstack([t_all_v, l_all_v, t_all_v]).reshape((-1, 6, 96))
    out = tf.concat([tv, out], axis=1)
    out = tf.reshape(out, (-1, 672))
    with tf.Session() as sess:
        out = sess.run(out)
    return out
