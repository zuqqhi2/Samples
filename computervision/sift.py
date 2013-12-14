from PIL import Image
from numpy import *
import os

def process_image(imagename,resultname,params='--edge-thresh 10 --peak-thresh 5'):
	if imagename[-3:] != 'pgm':
		im = Image.open(imagename).convert('L')
		im.save('tmp.pgm')
		imagename = 'tmp.pgm'

	cmmd = str('sift '+imagename+' --output='+resultname+' '+params)
	os.system(cmmd)
	print 'processed', imagename, 'to', resultname

def read_features_from_file(filename):
	f = loadtxt(filename)
	return f[:,:4],f[:,:4]

def write_features_to_file(filename,locs,desc):
	savetxt(filename,hstack((locs,desc)))

def plot_features(im, locs,circle=False):
	def draw_circle(c,r):
		t = arange(0,1.01,0.01)*2*pi
		xarr = r*cos(t) + c[0]
		yarr = r*sin(t) + c[1]

		for idx_y in yarr:
			for idx_x in xarr:
				x = int(idx_x)
				y = int(idx_y)
				im[y  ,x  ] = 0
				if y+1 < im.shape[0]:
					im[y+1,x  ] = 0
				if x+1 < im.shape[1]:
					im[y  ,x+1] = 0
		#plot(x,y, 'b', linewidge=2)
	
	#imshow(im)
	if circle:
		for p in locs:
			draw_circle(p[:2],p[2])
	else:
		for y in locs[:,0]:
			for x in locs[:,1]:
				y = int(y)
				x = int(x)
				im[y  ,x  ] = 0
				if y+1 < im.shape[0]:
					im[y+1,x  ] = 0
				if x+1 < im.shape[1]:
					im[y  ,x+1] = 0
		#plot(locs[:,0],locs[:,1],'ob')
	#axis('off')
	
	pil_im = Image.fromarray(uint8(im))
	pil_im.save('dest/sift.jpg')

if __name__ == '__main__':
	imname = 'source/empire.jpg'
	im1 = array(Image.open(imname).convert('L'))
	process_image(imname, 'empire.sift')
	l1,d1 = read_features_from_file('empire.sift')
	plot_features(im1,l1,circle=True)
