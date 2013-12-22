import math
import numpy as np


def mv_norm(x, mu, sigma):
  norm1 = 1 / (math.pow(2 * math.pi, len(x)/2.0) * math.pow(np.linalg.det(sigma), 1.0/2.0))
  x_mu = np.matrix(x-mu)
  norm2 = np.exp(-0.5 * x_mu * np.linalg.inv(sigma) * x_mu.T)
  return float(norm1 * norm2)

# test data
data = np.array([[-1,-1],[-1,0],[0,1],[1,1],[1,2]])

# dimension
K = 2
# time
N = len(data)

# average
mu = np.array([np.array([0,0]), np.array([1,0])])
# variance
sigma = np.array([np.eye(2), np.eye(2)])
# pi parameter
pi_k = [0.5, 0.5]

L = []
mu_iter = []
sigma_iter = []
pi_k_iter = []
diff = 1

num_iter = 1
while diff > 0.1:
  print 'num iteration : ', num_iter

  # E-step
  likelihood = np.zeros((N,K))
  gamma_nk = np.zeros((N,K))

  for k in range(K):
    likelihood[:,k] = [mv_norm(d, mu[k], np.array(sigma[k]))*pi_k[k] for d in data]

  for n in range(N):
    gamma_nk[n,:] = likelihood[n,:] / sum(likelihood[n,:])
  
  # M-step
  N_k = np.array([sum(gamma_nk[:,k]) for k in range(K)])
  #pi_k = N_k/sum(N_k)
  pi_k = N_k / len(data)
  mu = np.dot(gamma_nk.T, data)/N_k
  for k in range(K):
    sig = np.zeros((2,2))
    for n in range(N):
      x_mu = data[n,:] - mu[:,k]
      sig += gamma_nk[n,k] * np.outer(x_mu, x_mu.T)
    sigma[k] = np.array(sig/N_k[k])

  # iteration
  mu_iter.append(mu)
  sigma_iter.append(sigma)
  pi_k_iter.append(pi_k)

  l = sum(map(np.log,[sum(likelihood[n,:]) for n in range(N)]))/N
  print 'log likelihood: ', l

  if L:
    diff = math.fabs(L[-1] - l)
    L.append(l)
  else:
    L.append(l)

  num_iter += 1

# normal value
print 'which distribution the data is beloging : '
print [mv_norm(np.array([-1,-1]), mu[k], np.array(sigma[k])) for k in range(K)]
print [mv_norm(np.array([1,1]), mu[k], np.array(sigma[k])) for k in range(K)]
# outlier
print [mv_norm(np.array([10,-10]), mu[k], np.array(sigma[k])) for k in range(K)]

# density function's result
print 'p(training data) :'
for d in data:
  print sum([mv_norm(d, mu[k], np.array(sigma[k]))*pi_k[k] for k in range(K)])
# outlier
print 'p(outlier) :'
print sum([mv_norm(np.array([-1,-2]), mu[k], np.array(sigma[k]))*pi_k[k] for k in range(K)])
print sum([mv_norm(np.array([2,-1]), mu[k], np.array(sigma[k]))*pi_k[k] for k in range(K)])
print sum([mv_norm(np.array([3,0]), mu[k], np.array(sigma[k]))*pi_k[k] for k in range(K)])
print sum([mv_norm(np.array([10,-10]), mu[k], np.array(sigma[k]))*pi_k[k] for k in range(K)])
