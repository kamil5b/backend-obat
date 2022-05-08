package utilities

import (
	"errors"

	"golang.org/x/crypto/sha3"
)

func HashingPilihan(NIM, Nama string, kode byte) []byte {
	sbyte := []byte(Nama)
	shs := sha3.Sum256(sbyte)
	for i := 0; i < int(NIM[len(NIM)-1]); i++ {
		shs = sha3.Sum256(shs[:])
	}
	kodeUser := kode - NIM[len(NIM)-1] + NIM[len(NIM)-3]
	shs[len(NIM)-2] = kodeUser

	return shs[:]
}

func HashUser(ID int, Username string) []byte {
	sbyte := []byte(Username)
	shs := sha3.Sum256(sbyte)
	for i := 0; i < ID%5; i++ {
		shs = sha3.Sum256(shs[:])
	}
	return shs[:]
}

func HashKamil(state string) []byte {
	sbyte := []byte(state)
	shs := sha3.Sum256(sbyte)
	for i := 0; i < len(sbyte); i++ {
		shs = sha3.Sum256(shs[:])
	}
	return shs[:]
}

func HashEmail(Email string) []byte {
	sbyte := []byte(Email[:len(Email)-30])
	shs := sha3.Sum256(sbyte)
	for i := 0; i < len(sbyte); i++ {
		shs = sha3.Sum256(shs[:])
	}

	return shs[:]
}

func CompareHash(hash1, hash2 []byte) error {
	if len(hash1) != len(hash2) {
		return errors.New("hash tidak sama")
	}
	for i := 0; i < len(hash1); i++ {
		if hash1[i] != hash2[i] {
			return errors.New("hash tidak sama")
		}
	}
	return nil
}

func StatusHash(sessid, status []byte, countstatus int) (int, error) {
	//STATUS 0 : sessid[0] SWAP sessid[:-1] == status
	//STATUS 1 : sessid[1] SWAP sessid[:-2] == status
	//STATUS i : sessid[i] SWAP sessid[:-(i+1)] == status
	//COMPARE hash username dengan

	/*
		sessid : hash
		status : modified sessid
	*/

	for i := 0; i < countstatus; i++ {
		sid := sessid
		tmp := sid[i]
		sid[i] = sid[len(sid)-(i+1)]
		sid[len(sid)-(i+1)] = tmp
		if err := CompareHash(sid, status); err == nil {
			return i, nil
		}
	}

	return -1, errors.New("status invalid")
}
