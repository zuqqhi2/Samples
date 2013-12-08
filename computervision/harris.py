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

im = array(Image.open('source/empire.jpg').convert('L'))
harrisim = compute_harris_response(im)
filtered_coords = get_harris_points(harrisim, 6, 0.1)
#plot_harris_points(im, filtered_coords)

pil_im = Image.fromarray(uint8(im))
pil_im.save('dest/harris-original.jpg')
imsave('dest/harris-original.pdf', im)
pil_harrisim = Image.fromarray(uint8(harrisim))
pil_harrisim.save('dest/harris-response.jpg')
imsave('dest/harris-response.pdf', harrisim)

for p in filtered_coords:
    for idxY in range(3):
        for idxX in range(3):
            if idxY < 0:
                continue
            if idxX < 0:
                continue
            y = idxY - 1
            x = idxX - 1
            im[p[0]+x,p[1]+y] = 0

pil_im = Image.fromarray(uint8(im))
pil_im.save('dest/harris.jpg')
imsave('dest/harris.pdf', im)
