#-*- coding:utf-8 -*-

from scipy.ndimage import filters
from PIL import Image
from numpy import *
from pylab import *

def compute_harris_response(im, sigma=3):
    """
    グレースケール画像の各ピクセルについて
    Harrisコーナー検出器の応答関数を定義する
    """

    # Differencial Coefficient
    imx = zeros(im.shape)
    filters.gaussian_filter(im, (sigma,sigma), (0,1), imx)
    imy = zeros(im.shape)
    filters.gaussian_filter(im, (sigma,sigma), (1,0), imy)

    # Calculate element of Harris Matrix
    Wxx = filters.gaussian_filter(imx*imx, sigma)
    Wxy = filters.gaussian_filter(imx*imy, sigma)
    Wyy = filters.gaussian_filter(imy*imy, sigma)

    # Discriminant
    Wdet = Wxx*Wyy - Wxy**2
    Wtr = Wxx + Wyy

    return Wdet / Wtr

def get_harris_points(harrisim, min_dist=10, threshold=0.1):
    """
    Harris応答画像からコーナーを返す。
    min_distはコーナーや画像境界から分離する最小ピクセル数
    """

    # しきい値thresholdを超えるコーナー候補を見つける
    corner_threshold = harrisim.max() * threshold
    harrisim_t = (harrisim > corner_threshold) * 1

    # Get candidate cordinates
    coords = array(harrisim_t.nonzero()).T

    # Get candidate's value
    candidate_values = [harrisim[c[0],c[1]] for c in coords]

    # Sor candidates
    index = argsort(candidate_values)

    # Pussh coordinates for array
    allowed_locations = zeros(harrisim.shape)
    allowed_locations[min_dist:-min_dist,min_dist:-min_dist] = 1

    #Get best point with minist distance
    filtered_coords = []
    for i in index:
        if allowed_locations[coords[i,0],coords[i,1]] == 1:
            filtered_coords.append(coords[i])
            allowed_locations[(coords[i,0]-min_dist):(coords[i,0]+min_dist),(coords[i,1]-min_dist):(coords[i,1]+min_dist)] = 0

    return filtered_coords

def plot_harris_points(image, filtered_coords):
    """ 画像中に見つかったコーナーを描画 """
    figure()
    gray()
    imshow(image)
    plot([p[1] for p in filtered_coords],[p[0] for p in filtered_coords], '*')
    axis('off')
    show()

def get_descriptors(image, filtered_coords,wid=5):
	""" 各点について、点の周辺で幅2*wid+1の近傍ピクセル値を返す。（点の最小距離min_distance > widを仮定する) """
	desc = []
	for coords in filtered_coords:
		patch = image[coords[0]-wid:coords[0]+wid+1,coords[1]-wid:coords[1]+wid+1].flatten()
		desc.append(patch)
	
	return desc

def match(desc1,desc2,threshold=0.5):
	""" 正規化相互相関を用いて、第一の画像の各コーナー点記述子について第２の画像の対応点を選択する。  """
	n = len(desc1[0])

	d = -ones((len(desc1), len(desc2)))
	for i in range(len(desc1)):
		for j in range(len(desc2)):
			d1 = (desc1[i] - mean(desc1[i])) / std(desc1[i])
			d2 = (desc2[j] - mean(desc2[j])) / std(desc2[j])
			ncc_value = sum(d1 * d2) / (n-1)
			if ncc_value > threshold:
				d[i,j] = ncc_value
	
	ndx = argsort(-d)
	matchscores = ndx[:,0]

	return matchscores

def match_twosided(desc1,desc2,threshold=0.5):
	""" match()双方向で一致を調べるバージョン  """
	matches_12 = match(desc1,desc2, threshold)
	matches_21 = match(desc2,desc1, threshold)

	ndx_12 = where(matches_12 >= 0)[0]

	for n in ndx_12:
		if matches_21[matches_12[n]] != n:
			matches_12[n] = -1
	
	return matches_12

def appendimages(im1, im2):
	""" 2つの画像を左右に並べた画像を返す """

	rows1 = im1.shape[0]
	rows2 = im2.shape[0]

	if rows1 < rows2:
		im1 = concatenate((im1,zeros((rows2-rows1,im1.shape[1]))), axis=0)
	elif rows1 > rows2:
		im2 = concatenate((im2,zeros((rows1-rows2,im2.shape[1]))), axis=0)
	
	return concatenate((im1,im2), axis=1)

def plot_matches(im1,im2,locs1,locs2,matchscores,show_below=True):
	im3 = appendimages(im1,im2)
	if show_below:
		im3 = vstack((im3,im3))
	
	cols1 = im1.shape[1]

	for i,m in enumerate(matchscores):
		if m > 0:
			for t in range(0,10000):
				norm_t = t / 15000.0
				y1 = int( locs2[m][0]          * norm_t + locs1[i][0] * (1.0 - norm_t))
				x  = int((locs2[m][1] + cols1) * norm_t + locs1[i][1] * (1.0 - norm_t))
				y2 = y1-1
				y3 = y1+1
				if y1 < im3.shape[0] and x < im3.shape[1] and y3 < im3.shape[0] and y2 >= 0:
					im3[y1,x] = 0
					im3[y2,x] = 0
					im3[y3,x] = 0
	
	return im3

wid = 5
im1 = array(Image.open('source/data/crans_1_small.jpg').convert('L'))
pil_im = Image.fromarray(uint8(im1))
pil_im.save('dest/harris-match-original1.jpg')
harrisim = compute_harris_response(im1, wid)
filtered_coords1 = get_harris_points(harrisim, wid+1)
d1 = get_descriptors(im1,filtered_coords1,wid)
imsave('dest/harris-match-response1.pdf', harrisim)

im2 = array(Image.open('source/data/crans_2_small.jpg').convert('L'))
pil_im = Image.fromarray(uint8(im2))
pil_im.save('dest/harris-match-original2.jpg')
harrisim = compute_harris_response(im2,wid)
filtered_coords2 = get_harris_points(harrisim, wid+1)
d2 = get_descriptors(im2,filtered_coords2,wid)
imsave('dest/harris-match-response2.pdf', harrisim)

matches = match_twosided(d1,d2)
if len(matches) > 100:
	matches = matches[:100]

im3 = plot_matches(im1,im2,filtered_coords1,filtered_coords2,matches)
pil_im = Image.fromarray(uint8(im3))
pil_im.save('dest/harris-match-result.jpg')

for p in filtered_coords1:
    for idxY in range(3):
        for idxX in range(3):
            if idxY < 0:
                continue
            if idxX < 0:
                continue
            y = idxY - 1
            x = idxX - 1
            im1[p[0]+x,p[1]+y] = 0
pil_im = Image.fromarray(uint8(im1))
pil_im.save('dest/harris-match-im1.jpg')

for p in filtered_coords2:
    for idxY in range(3):
        for idxX in range(3):
            if idxY < 0:
                continue
            if idxX < 0:
                continue
            y = idxY - 1
            x = idxX - 1
            im2[p[0]+x,p[1]+y] = 0
pil_im = Image.fromarray(uint8(im2))
pil_im.save('dest/harris-match-im2.jpg')

