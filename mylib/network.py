import numpy as np
import tensorflow as tf


class Network:
    def __init__(self, v=None, n_in=672, n_out=230):
        # session
        self._sess = tf.Session()

        # in, out
        self.n_in = n_in
        self.n_out = n_out

        # create network
        if v is not None:
            self.W = tf.constant(v['W'], dtype=tf.float32)
            self.b = tf.constant(v['b'], dtype=tf.float32)
        else:
            self.W = tf.Variable(tf.zeros([self.n_in, self.n_out]))
            self.b = tf.Variable(tf.zeros([self.n_out]))

        self._x = tf.placeholder(tf.float32, shape=[None, self.n_in])
        self._t = tf.placeholder(tf.float32, shape=[None, self.n_out])
        self._y = tf.nn.sigmoid(tf.add(tf.matmul(self._x, self.W), self.b))
        self._loss = self._t * -tf.log(tf.clip_by_value(self._y, 1e-10, 1.0)) +\
            (1-self._t) * -tf.log(tf.clip_by_value(1 - self._y, 1e-10, 1.0))
        self._train_step = tf.train.AdamOptimizer().minimize(self._loss)
        _, self._auc = tf.metrics.auc(self._t, self._y)

        # session init
        self._sess.run(tf.global_variables_initializer())
        self._sess.run(tf.local_variables_initializer())

        # stat
        self.max_idx = 0
        self.max_auc = 0.5
        self.max_w = self._sess.run(self.W)
        self.max_b = self._sess.run(self.b)
        self.aucs = []

    def __enter__(self):
        return self

    def __exit__(self, type, value, traceback):
        self._sess.close()

    def train(self, x_train, y_train, x_test, y_test, iters=100, n_batch=1024):
        x = self._x
        y = self._t
        for i in range(iters):
            perm = np.random.permutation(len(x_train))
            for j in range(len(x_train) // n_batch):
                x_batch = x_train[perm[j*n_batch:(j+1)*n_batch]]
                y_batch = y_train[perm[j*n_batch:(j+1)*n_batch]]
                self._sess.run(
                    self._train_step,
                    feed_dict={self._x: x_batch, self._t: y_batch}
                )
            auc = self._sess.run(self._auc, feed_dict={
                x: x_test,
                y: y_test,
            })
            if auc > self.max_auc:
                self.max_idx = i
                self.max_auc = auc
                self.max_w = self._sess.run(self.W)
                self.max_b = self._sess.run(self.b)
                self.aucs.append(auc)
                print(i, auc, 'up')
            else:
                print(i, auc, 'down')

    def predict(self, x):
        pass

    def get_data(self):
        return {
            'w': self.max_w,
            'b': self.max_b,
            'max_idx': self.max_idx,
            'aucs': self.aucs,
        }
