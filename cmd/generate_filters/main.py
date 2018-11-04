import pickle

import numpy as np

print('generate filters...')

f1 = np.random.randn(3, 32, 1, 32)
f2 = np.random.randn(4, 32, 1, 32)
f3 = np.random.randn(5, 32, 1, 32)

filters = {
    'f1': f1,
    'f2': f2,
    'f3': f3,
}

with open('data/nn/filters.bin', 'wb') as f:
    pickle.dump(filters, f)

print('done')
