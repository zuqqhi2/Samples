#-*- coding:utf-8 -*-
from numpy import *
from numpy import random
from scipy.ndimage import filters
from scipy.misc import imsave

def denoise(im, U_init, tolerance=0.1, tau=0.125, tv_weight=100):
    """
    A. Chambolle (2005)の数式(11)記載の計算手順に基づく
    Rudin-Osher-Fatemi(ROF)ノイズ除去モデルの実装。

    入力：ノイズのある入力像(グレースケール)、
    　　　Uの初期ガウス分布、
    　　　終了判断基準の許容誤差(tolerance)、
    　　　ステップ長(tau)、
    　　　TV正規化項の重み(tv_weight)
    出力：ノイズ除去された画像、残余テクスチャ
    """

    m,n = im.shape

    # Initialize
    U = U_init
    Px = im
    Py = im
    error = 1

    while (error > tolerance):
        Uold = U

        # Gradient of main variable
        GradUx = roll(U, -1, axis=1) - U
        GradUy = roll(U, -1, axis=0) - U

        # Update binary variable
        PxNew = Px + (tau / tv_weight) * GradUx
        PyNew = Py + (tau / tv_weight) * GradUy
        NormNew = maximum(1,sqrt(PxNew**2+PyNew**2))

        # Update dual variable (x,y) element
        Px = PxNew/NormNew
        Py = PyNew/NormNew

        # Update main variable
        RxPx = roll(Px, 1, axis=1)
        RyPy = roll(Py, 1, axis=0)

        # Divergence of binary area
        DivP = (Px-RxPx) + (Py-RyPy)

        # Update main variable
        U = im + tv_weight * DivP

        # Update error
        error = linalg.norm(U-Uold) / sqrt(n*m)

    # Denoise Image and Remain Texture
    return U, im-U

# Make noise image
im = zeros((500, 500))
im[100:400, 100:400] = 128
im[200:300, 200:300] = 255
im = im + 30*random.standard_normal((500,500))

U,T = denoise(im,im)
G = filters.gaussian_filter(im,10)

# Save the result
imsave('synth_original.pdf', im)
imsave('synth_rof.pdf', U)
imsave('synth_gaussian.pdf', G)
