from collections import defaultdict

import numpy as np
import tensorflow as tf

import protobuf.nn_auto_pb2 as pb
from lib import get_char

# element n
EN = 192

# read filters
f1 = tf.Variable(tf.random_normal([3, 32, 1, 32]), name='f1')
f2 = tf.Variable(tf.random_normal([4, 32, 1, 32]), name='f2')
f3 = tf.Variable(tf.random_normal([5, 32, 1, 32]), name='f3')
saver = tf.train.Saver()
sess = tf.Session()
saver = saver.restore(sess, './data/tf/title_filters')


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


def cnv_title(title):
    global f1, f2, f3
    x = tf.constant(str2bytes(title), dtype=tf.float32)
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
    return tf.reshape(out, (-1, 1, 96))


def wiki2vec(page):
    print(page.title)
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
                cv[cat[ch]['rank']] += 2
        c_all_v += np.where(cv > 1.0, 1.0, cv)
    c_all_v /= len(categories)
    #  -- links
    l_all_v = np.zeros(EN)
    links = page.links
    for link in links:
        cv = np.zeros(EN)
        for ch in link:
            ch = get_char(ch)
            if ch in lin:
                cv[lin[ch]['rank']] += 2
        l_all_v += np.where(cv > 1.0, 1.0, cv)
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
    return out


# load data
data = pb.Data()
wiki = pb.Wiki()
stat = pb.Statistics()

with open('data/data.bin', 'rb') as f:
    data.ParseFromString(f.read())

with open('data/wiki.bin', 'rb') as f:
    wiki.ParseFromString(f.read())

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

# create title to ids dict
title2id = defaultdict(lambda: len(title2id))
id2title = {}
inputs = np.zeros((0, 672))
on_label = {
    'page_ids': [],
    'out': [],
}

for v in wiki.pages.keys():
    id2title[title2id[v]] = v
    if v in data.taggedPages:
        out = np.zeros(230)
        for attr in data.taggedPages[v].attrs:
            out[attr-1] = 1
        on_label['page_ids'].append(title2id[v])
        on_label['out'].append(out)
    page = wiki.pages[v]
    res = sess.run(wiki2vec(page))
    inputs = np.vstack([inputs, res])

title2id = dict(title2id)
np.save('data/npy/title2id', title2id)
np.save('data/npy/id2title', id2title)
np.save('data/npy/input', inputs)
np.save('data/npy/on_label', on_label)
