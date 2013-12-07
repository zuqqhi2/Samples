from PIL import Image
from numpy import *
import imtools
import os

im = array(Image.open('empire.jpg').convert('L'))
im2,cdf = imtools.histeq(im)
pil_im = Image.fromarray(uint8(im2))
try:
    pil_im.save("./histeq.jpg")
except IOError:
    print "cannot save", "empire.jpg"
