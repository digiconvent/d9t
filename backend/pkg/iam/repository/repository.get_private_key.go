package iam_repository

import "crypto/rsa"

func (r *IamRepository) GetPrivateKey() *rsa.PrivateKey {
	return r.privateKey
}
