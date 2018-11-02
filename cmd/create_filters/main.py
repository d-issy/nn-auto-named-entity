import tensorflow as tf

print('variable initalize...')

f1 = tf.Variable(tf.random_normal([3, 32, 1, 32]), name='f1')
f2 = tf.Variable(tf.random_normal([4, 32, 1, 32]), name='f2')
f3 = tf.Variable(tf.random_normal([5, 32, 1, 32]), name='f3')


saver = tf.train.Saver([f1, f2, f3])

with tf.Session() as sess:
    sess.run(tf.global_variables_initializer())
    saver.save(sess, './data/tf/title_filters')

print('done!!!')
