package mithril

import (
	"fmt"
)

const (
	ModuleName                           = "2000-cardano-mithril"
	KeyFirstCertificateInEpochPrefix     = "fcInEpoch"
	KeyFirstCertificateMsdInEpochPrefix  = "fcMsdInEpoch"
	KeyFirstCertificateTsInEpochPrefix   = "fcTsInEpoch"
	KeyLatestCertificateMsdInEpochPrefix = "LcMsdInEpoch"
	KeyLatestCertificateTsInEpochPrefix  = "LcTsInEpoch"
	KeyMSDCertificateHashPrefix          = "MSDCertificateHash"
)

func FcInEpochKey(epoch uint64) []byte {
	return []byte(FcInEpochPath(epoch))
}

func FcInEpochPath(epoch uint64) string {
	return fmt.Sprintf("%s/%v", KeyFirstCertificateInEpochPrefix, epoch)
}

func MSDCertificateHashKey(hash string) []byte {
	return []byte(MSDCertificateHashPath(hash))
}

func MSDCertificateHashPath(hash string) string {
	return fmt.Sprintf("%s/%s", KeyMSDCertificateHashPrefix, hash)
}

func FcMsdInEpochKey(epoch uint64) []byte {
	return []byte(FcMsdInEpochPath(epoch))
}

func FcMsdInEpochPath(epoch uint64) string {
	return fmt.Sprintf("%s/%v", KeyFirstCertificateMsdInEpochPrefix, epoch)
}

func FcTsInEpochKey(epoch uint64) []byte {
	return []byte(FcTsInEpochPath(epoch))
}

func FcTsInEpochPath(epoch uint64) string {
	return fmt.Sprintf("%s/%v", KeyFirstCertificateTsInEpochPrefix, epoch)
}

func LcMsdInEpochKey(epoch uint64) []byte {
	return []byte(LcMsdInEpochPath(epoch))
}

func LcMsdInEpochPath(epoch uint64) string {
	return fmt.Sprintf("%s/%v", KeyLatestCertificateMsdInEpochPrefix, epoch)
}

func LcTsInEpochKey(epoch uint64) []byte {
	return []byte(LcTsInEpochPath(epoch))
}

func LcTsInEpochPath(epoch uint64) string {
	return fmt.Sprintf("%s/%v", KeyLatestCertificateTsInEpochPrefix, epoch)
}
