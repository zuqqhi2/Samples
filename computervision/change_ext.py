from PIL import Image
import os

filelist = ['empire.jpg']

for infile in filelist:
    outfile = os.path.splitext(infile)[0] + ".JPG"
    if infile != outfile:
        try:
            Image.open('source/' + infile).save('dest/' + outfile)
        except IOError:
            print "cannot convert", infile
