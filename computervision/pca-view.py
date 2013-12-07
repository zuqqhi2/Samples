from PIL import Image
from numpy import *
import imtools
import pca
import os

imlist = imtools.get_imlist("data/a_thumbs")
im = array(Image.open(imlist[0]))
m,n = im.shape[0:2]
imnbr = len(imlist)

immatrix = array([array(Image.open(im)).flatten() for im in imlist], 'f')

V,S,immean = pca.pca(immatrix)

pil_im = Image.fromarray(uint8(immean.reshape(m,n)))
try:
    pil_im.save("./pca.jpg")
except IOError:
    print "cannot save", "pca.jpg"

for i in range(7):
    tmp = (V[i] - min(V[i])) / (max(V[i]) - min(V[i])) * 255
    pil_im = Image.fromarray(uint8(tmp.reshape(m,n)))
    
    print min(V[i]), max(V[i])
    print tmp.reshape(m,n)
    try:
        pil_im.save("./pca-" + str(i) + ".jpg")
    except IOError:
        print "cannnot save", "pca-" + str(i) + ".jpg"
