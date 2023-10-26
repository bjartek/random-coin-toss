// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../scripts/pseudo-random-generator/next_uint64_from_address.cdc (293B)
// ../../../scripts/pseudo-random-generator/next_uint64_new_prg.cdc (263B)
// ../../../transactions/coin-toss/0_commit_coin_toss.cdc (737B)
// ../../../transactions/coin-toss/1_reveal_coin_toss.cdc (666B)
// ../../../transactions/flow-token/transfer_flow.cdc (675B)
// ../../../transactions/pseudo-random-generator/next_uint64.cdc (460B)
// ../../../transactions/pseudo-random-generator/setup_prg.cdc (675B)
// ../../../transactions/random-beacon-history/heartbeat.cdc (1.026kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _scriptsPseudoRandomGeneratorNext_uint64_from_addressCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\xb1\x4a\xc4\x40\x10\x86\xfb\x7d\x8a\x9f\x14\x92\x6d\x52\x89\xc5\x21\x1e\xc7\x15\xc1\x2e\x04\x7c\x80\xc9\xee\x5e\x5c\x48\x66\x96\xc9\x2c\x2a\xe2\xbb\x0b\x9e\x1e\x57\x78\x53\x0d\x3f\xdf\xcc\xff\xe5\xb5\x88\x1a\x9a\x61\x4b\x35\xca\x48\x1c\x65\xed\x13\x27\x25\x13\x6d\x9c\x2b\x75\xc2\xa9\x32\x56\xca\xdc\x16\x9d\x0f\x31\x6a\xda\xb6\x1d\x7e\x17\xbf\xc3\xcb\x33\xdb\xc3\x3d\x3e\x9d\x03\x00\x4d\x56\x95\x31\x27\x3b\x84\x20\x95\xed\xea\xca\x77\x73\xb2\x23\x15\x9a\xf2\x92\xed\xe3\xf1\xee\xdf\xd6\x6e\x18\xfb\xa7\xf6\xe7\xd9\xdf\xdc\xe0\xea\xb4\xe4\x30\x90\xbd\x5e\x60\xdf\x4d\xa2\x2a\x6f\xad\xbf\x44\xfb\x8e\xd3\xbb\x9d\x25\xaf\xe3\x3d\x0a\x71\x0e\x6d\x73\x94\xba\x44\xb0\x18\x4e\x99\x23\x86\xb1\x07\x19\xe8\xac\xdc\x78\xf7\xe5\xbe\x03\x00\x00\xff\xff\x6a\xe9\xce\xd4\x25\x01\x00\x00"

func scriptsPseudoRandomGeneratorNext_uint64_from_addressCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsPseudoRandomGeneratorNext_uint64_from_addressCdc,
		"scripts/pseudo-random-generator/next_uint64_from_address.cdc",
	)
}

func scriptsPseudoRandomGeneratorNext_uint64_from_addressCdc() (*asset, error) {
	bytes, err := scriptsPseudoRandomGeneratorNext_uint64_from_addressCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/pseudo-random-generator/next_uint64_from_address.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf3, 0x39, 0xc1, 0xca, 0x28, 0xa5, 0x95, 0xec, 0xa8, 0x54, 0x1d, 0x80, 0x37, 0x3a, 0xec, 0xf6, 0xf9, 0x52, 0xdb, 0xcd, 0x9f, 0x99, 0x5c, 0x93, 0x81, 0xd4, 0x25, 0xf3, 0xb3, 0xbd, 0xb, 0x38}}
	return a, nil
}

var _scriptsPseudoRandomGeneratorNext_uint64_new_prgCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8d\x31\x6b\xc3\x30\x10\x46\x77\xfd\x8a\x0f\x4f\x36\xb4\x9e\x4a\x29\xa6\x9d\x4d\xa7\x1a\x43\xa7\xd2\x41\xb1\xce\xc1\x60\x9f\xcc\xe9\x04\x09\x21\xff\x3d\x28\x8e\x13\x0f\x59\x84\xee\xe3\xf1\xde\x30\xcd\x5e\x14\x59\x13\x28\x3a\xdf\x5a\x76\x7e\xaa\x89\x49\xac\x7a\xc9\x8c\x99\xe3\x0e\x7d\x64\x4c\x76\xe0\x3c\x10\xb9\x0a\x7f\xbf\xdf\xac\x1f\xff\x2f\x08\x76\xd4\x0a\xe9\x7a\x7f\x2b\xd6\x0f\x4e\xc6\x00\xc0\x48\x8a\x59\xf6\xf8\x7c\xc5\x53\x77\xd9\x09\x59\xa5\xa6\xad\xf3\xe0\xa3\x74\xf4\xd3\x2f\x04\x53\x08\x15\x52\x6a\x2d\xa4\xb7\xb8\x4a\xef\x66\xb1\xec\x6e\xbd\xaf\x94\x29\x99\x0e\xba\x0c\xf9\x06\x75\x14\x54\xfc\x31\x11\x8f\x51\x48\xa3\xf0\x46\x61\xce\xe6\x12\x00\x00\xff\xff\x86\x40\x77\x62\x07\x01\x00\x00"

func scriptsPseudoRandomGeneratorNext_uint64_new_prgCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsPseudoRandomGeneratorNext_uint64_new_prgCdc,
		"scripts/pseudo-random-generator/next_uint64_new_prg.cdc",
	)
}

func scriptsPseudoRandomGeneratorNext_uint64_new_prgCdc() (*asset, error) {
	bytes, err := scriptsPseudoRandomGeneratorNext_uint64_new_prgCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/pseudo-random-generator/next_uint64_new_prg.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5d, 0xfb, 0x15, 0x1e, 0xd6, 0x59, 0x94, 0x7b, 0xcc, 0xcd, 0xdd, 0x9c, 0x2a, 0x3d, 0xa7, 0xe9, 0xd2, 0x83, 0x5e, 0x2, 0x1e, 0x70, 0xbe, 0x64, 0xfb, 0xd4, 0xd, 0x8a, 0x64, 0x35, 0xaf, 0xa7}}
	return a, nil
}

var _transactionsCoinToss0_commit_coin_tossCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x51\x3d\x8f\x9b\x40\x10\xed\xf7\x57\x8c\x29\x12\x90\x62\x68\xa2\x14\x88\x8b\x64\x9d\x74\x52\xba\x28\x77\x49\x94\x72\x80\xc1\xac\x0c\x3b\x68\x19\x4c\xac\xc8\xff\x3d\x5a\x3e\x16\xbb\x38\xbb\xb1\x98\x7d\x6f\xde\xc7\xe8\xb6\x63\x2b\x10\xbc\x34\x3c\xbe\xf1\x89\x4c\xa0\xd4\x3a\x7b\x66\x6d\xde\xb8\xef\x03\xa5\xc4\xa2\xe9\xb1\x10\xcd\x26\xcc\x49\x0e\x2d\x0f\x46\x52\xf8\xf9\xa2\xff\x7e\xf9\x1c\xc1\x3f\xa5\x00\x00\x3a\x4b\x1d\x5a\x0a\x7b\x7d\x34\x64\x53\x38\x0c\x52\x1f\x8a\xc2\x61\x1d\x06\x96\x5f\x92\xc0\x6f\x2d\x75\x69\x71\x84\xf6\x02\x39\x09\xe0\xb4\x10\x2a\xcb\xad\x1b\x79\x37\x70\xc6\xa1\x11\xcf\x6c\x48\xa0\x6a\x78\xfc\xe5\xa6\xf0\x04\xb3\x50\x9c\xb3\xb5\x3c\x66\x1f\x3c\x2d\x9e\x00\x5f\x43\xb7\x2f\x85\xa4\x17\xb6\x78\xa4\xa4\x5a\xdf\xa7\xe7\x68\x77\xb7\xd7\xd9\xc8\xf6\xdb\xfa\x78\x5c\x3c\x86\xb8\xa4\xf5\xc1\x23\x4f\xbc\xcd\xf4\xcc\x6d\xab\xc5\x27\x32\x25\x1c\xdd\x3f\x58\x2a\x48\x77\xf7\x21\x96\x99\x13\x5c\x5b\x8e\x8b\x89\xbf\x7e\xba\x9a\x53\xc8\xf6\x39\xbd\x27\x57\x53\x71\x02\xa9\x51\xe0\x1b\x94\x6c\x3e\x0a\x60\x63\x09\xcb\x0b\xd4\x78\xa6\x4d\x18\x5c\x7c\x2a\x3d\x57\x57\x6b\x6f\x72\xe9\x28\x44\x49\x37\x0f\x3f\x66\xca\xeb\x5c\xd8\x77\x94\x3a\x82\xdd\x13\x18\xdd\xdc\xdc\x6f\x3a\x35\x1a\x5d\x84\xc1\x1f\x1e\x1e\xab\xee\x82\xcd\xfd\x55\xdd\xfa\x7f\x75\xf0\xc9\xfe\xca\x10\x76\xe5\x2d\xc7\xf2\xd0\xc5\x6b\x8f\x67\x0a\xb3\xfd\x82\xfd\x04\xc2\x8f\x6d\xab\x59\xf1\xfa\x3f\x00\x00\xff\xff\xde\x75\x3c\x80\xe1\x02\x00\x00"

func transactionsCoinToss0_commit_coin_tossCdcBytes() ([]byte, error) {
	return bindataRead(
		_transactionsCoinToss0_commit_coin_tossCdc,
		"transactions/coin-toss/0_commit_coin_toss.cdc",
	)
}

func transactionsCoinToss0_commit_coin_tossCdc() (*asset, error) {
	bytes, err := transactionsCoinToss0_commit_coin_tossCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transactions/coin-toss/0_commit_coin_toss.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x81, 0x16, 0x88, 0x77, 0x3e, 0x75, 0xcc, 0x3d, 0xb1, 0xd, 0x48, 0x95, 0x60, 0xb5, 0x35, 0x1f, 0x3a, 0x38, 0x61, 0xab, 0xf, 0xf1, 0xeb, 0x64, 0xbf, 0xaa, 0x35, 0x6b, 0x6e, 0xa1, 0x2f, 0xb9}}
	return a, nil
}

var _transactionsCoinToss1_reveal_coin_tossCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\x4f\x6f\xa3\x30\x10\xc5\xef\xfe\x14\x13\x0e\x2b\x38\x04\x72\x8e\xd8\x64\xa3\x95\x7a\xac\xaa\x34\xea\xdd\x81\x81\x58\x35\x33\xc8\x36\x45\x51\x94\xef\x5e\xf1\xc7\x0e\x89\x0f\x08\xdb\xef\xbd\xf9\xcd\x58\x35\x2d\x1b\x07\xd1\x9b\xe6\xfe\xc4\xdf\x48\x91\x10\xfe\xec\x3f\x2b\x3a\xb1\xb5\x91\x10\xce\x48\xb2\xb2\x70\x8a\x09\x6e\x42\x00\x00\xb4\x06\x5b\x69\x30\xb6\xaa\x26\x34\x5b\x38\x74\xee\x72\x28\x0a\xee\xc8\x25\x70\x1b\x25\xc3\xca\x32\xd0\x2c\x4b\x68\xae\x60\xb0\x40\xd5\x3a\xa8\x0c\x37\x60\x1d\x1b\x59\x63\xd0\x69\x74\x41\x90\xaf\x61\x4a\x4d\x07\x6b\xfe\xcf\x83\xa4\xc7\x49\xb0\x8b\x87\x88\x2d\xbc\x9e\x7f\x4e\x99\x1f\xd2\x5d\x92\x10\x3c\xac\xfd\x1e\x5a\x49\xaa\x88\xa3\x77\x86\xa3\xc7\xe0\x8e\xca\x55\x94\x88\x25\xeb\x11\x7f\x50\x6a\x38\x0f\xb4\x25\x62\xa3\xa8\x5e\xa0\x3f\xd1\xf6\x8a\x48\x51\x6d\x07\xdc\x40\x62\x46\xbf\xdf\xc6\xb3\x6f\x0b\xf9\x7a\xfe\x5d\x94\x53\x55\xc8\x48\xcf\x52\x4b\x2a\x10\x76\xb0\x49\x37\x8b\xf1\xf9\x62\x95\xe6\xfe\x4b\x76\xda\xc1\x5f\x3f\x9b\x33\x1b\xc3\x7d\xfe\x27\x3c\x5d\x3a\x0a\xfc\x70\xb2\x79\xc2\x59\xe5\xef\xc7\xeb\x64\xf5\x94\x9d\x65\x50\x62\xcb\x56\x2d\xfa\x51\xe4\x78\x68\x3a\x04\xc3\xe8\x7c\xf2\x05\x9e\x74\xb6\x4f\x55\xf3\xb5\x4f\x79\x3c\xc0\x1d\x50\x5b\x7c\xe9\xa9\x44\xeb\x0c\x5f\x43\xd1\x87\x5a\x4c\xdf\xfb\x6f\x00\x00\x00\xff\xff\xd8\xa4\xcf\x38\x9a\x02\x00\x00"

func transactionsCoinToss1_reveal_coin_tossCdcBytes() ([]byte, error) {
	return bindataRead(
		_transactionsCoinToss1_reveal_coin_tossCdc,
		"transactions/coin-toss/1_reveal_coin_toss.cdc",
	)
}

func transactionsCoinToss1_reveal_coin_tossCdc() (*asset, error) {
	bytes, err := transactionsCoinToss1_reveal_coin_tossCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transactions/coin-toss/1_reveal_coin_toss.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbe, 0x8b, 0xe7, 0xf9, 0x9b, 0xdf, 0x2e, 0x57, 0x33, 0xfe, 0xa6, 0x4, 0xd6, 0x6e, 0xa3, 0x3e, 0x50, 0x2e, 0xd3, 0x2e, 0x47, 0x0, 0x5e, 0xa2, 0x6e, 0x3e, 0xc4, 0x30, 0x6a, 0xad, 0xd0, 0x67}}
	return a, nil
}

var _transactionsFlowTokenTransfer_flowCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\xcd\x6a\xeb\x30\x10\x85\xf7\x7e\x8a\xb9\x5e\x04\x1b\xee\x75\x36\x97\x2e\x4c\x7e\x08\x81\x3c\x40\x69\xbb\x97\xe5\xb1\x33\x54\x91\xc4\x78\x1c\xa7\x84\xbc\x7b\xf1\x3f\x4e\xe9\x2c\x2c\x98\xa3\x33\xf3\xe9\x60\xba\x78\xc7\x02\xe1\xa9\xb6\x25\x65\x06\xdf\xdc\x27\xda\x30\x98\xda\xc6\x35\x43\x2b\x10\x56\xb6\x52\x5a\xc8\xd9\x88\x51\x93\x27\xb4\x92\xc2\x21\xcf\x19\xab\xea\x2f\xa8\x8b\xab\xdb\xc6\xfb\x89\x6e\x2f\xff\x63\xb8\x07\x00\x00\xdd\xc7\xa0\x80\x67\x77\xa5\x1c\xf9\x43\xd5\x46\x52\x58\x4d\xa3\x93\xae\x33\xdd\x63\xd4\x48\x57\xe4\x14\x56\xf7\x05\x55\xf2\x3a\x28\x8f\x79\xb0\x67\xf4\x8a\x31\xaa\xa8\xb4\xad\xe5\x50\xcb\xf9\xa0\x75\x0b\x32\x02\xb4\x55\xa1\x29\x92\x05\x00\x6c\xa1\xf7\x24\x99\x63\x76\xcd\xe6\x99\x67\x17\x15\xec\x2e\x29\xac\x2b\x71\xac\x4a\x5c\x17\xa3\xde\xc9\xf1\x9f\xe5\xf0\x91\x1a\xb6\x50\xa2\x0c\x08\x73\x4c\x71\x52\xa2\x1c\x95\x57\x19\x19\x92\xaf\xcd\xaf\x4f\xdb\x45\x6b\x5f\x67\x86\xf4\xbc\x6f\xd4\xe2\x69\x63\x5b\x03\x77\xb4\xec\xee\xf7\xe0\x95\x25\x1d\x85\x47\x57\x9b\x1c\xac\x13\xe8\x2f\x4e\xb9\x02\x63\x81\x8c\x56\x63\xd8\x7b\x1f\x41\x77\xe0\x0d\x75\x2d\xf8\x9c\xda\x68\x4b\x72\xf4\xae\x22\x89\x16\xfb\xfa\x8c\x36\xff\x7e\x06\x9c\x34\x24\xe7\x9c\x55\xb3\x34\xb4\x35\xfe\x28\xfd\xb9\x90\xe7\xd7\x8c\x6c\x8f\xe0\x3b\x00\x00\xff\xff\xc6\xf7\x99\x11\xa3\x02\x00\x00"

func transactionsFlowTokenTransfer_flowCdcBytes() ([]byte, error) {
	return bindataRead(
		_transactionsFlowTokenTransfer_flowCdc,
		"transactions/flow-token/transfer_flow.cdc",
	)
}

func transactionsFlowTokenTransfer_flowCdc() (*asset, error) {
	bytes, err := transactionsFlowTokenTransfer_flowCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transactions/flow-token/transfer_flow.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc0, 0xd4, 0xf8, 0xc3, 0xb4, 0xf4, 0xb4, 0x79, 0x72, 0x7a, 0x11, 0xd9, 0xe7, 0xc8, 0x28, 0x5b, 0x9e, 0xc2, 0x70, 0x83, 0xf1, 0xa9, 0xe6, 0x6a, 0x55, 0x98, 0x9e, 0x36, 0x8c, 0x86, 0x52, 0x62}}
	return a, nil
}

var _transactionsPseudoRandomGeneratorNext_uint64Cdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x6a\x32\x51\x0c\x85\xf7\xf3\x14\x07\x17\xff\x3f\x52\xd0\x16\x4a\x17\xa2\x05\x57\x22\x74\x21\x4a\x1f\x20\x8e\x71\xe6\xd2\x99\xe4\x92\x9b\xd1\x42\xf1\xdd\xcb\x74\x4a\xab\xd6\x2c\x2e\xdc\xe4\xe4\xe4\x7c\xa1\x89\x6a\x8e\xc1\x2a\x71\xbb\xd3\x35\xc9\x4e\x9b\x05\x0b\x1b\xb9\xda\x20\xcb\xc6\xe3\x31\x36\x74\xe0\x04\x92\x1d\xea\x20\x6f\x09\x84\xd1\x6a\xbd\x80\x71\xd2\xd6\x0a\x46\x10\x78\xc5\x48\xa1\x14\xb6\xff\x09\xc9\xd5\xa8\xe4\xaf\x8d\xd8\x6e\xeb\x50\x40\xa8\xe1\x14\xa9\xe0\xce\x30\x73\x23\x49\x54\x78\x50\xc9\xcb\xfe\x58\x50\x79\x61\x29\xbd\x9a\x60\x29\x3e\xc4\x47\x06\x00\xd1\x38\x92\x71\xde\x5b\x4f\x30\x6f\xbd\x9a\x17\x85\xb6\xbf\x92\xae\xc2\x1e\x35\x3b\xa2\x95\x98\x7d\xc7\x18\x6d\xd5\x4c\x8f\xd3\x7f\x37\xc1\xba\xfc\xcf\xf9\xde\xb4\x99\xe0\xb6\x60\xd3\x33\xac\xc8\xab\xf3\x4b\x5d\x1d\xc8\x10\x30\xc3\xfd\x45\xf7\x58\x85\x9a\x11\x30\xc5\x35\xd1\xd5\x7a\xcf\x55\x8e\x84\xdf\xfd\x75\x29\xfe\xf4\x98\x0f\xff\x08\x3a\xff\x80\x3b\x3c\x5c\x4c\x4e\xe7\x9f\x9f\xc9\x29\xeb\xdf\xd3\x67\x00\x00\x00\xff\xff\x15\x3d\x47\x96\xcc\x01\x00\x00"

func transactionsPseudoRandomGeneratorNext_uint64CdcBytes() ([]byte, error) {
	return bindataRead(
		_transactionsPseudoRandomGeneratorNext_uint64Cdc,
		"transactions/pseudo-random-generator/next_uint64.cdc",
	)
}

func transactionsPseudoRandomGeneratorNext_uint64Cdc() (*asset, error) {
	bytes, err := transactionsPseudoRandomGeneratorNext_uint64CdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transactions/pseudo-random-generator/next_uint64.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1e, 0x3e, 0x73, 0xb, 0xfd, 0x37, 0xdc, 0x97, 0x70, 0xa6, 0xb4, 0x92, 0x4a, 0x19, 0xf8, 0x34, 0xbf, 0xa8, 0xd, 0x50, 0x6e, 0xad, 0x32, 0xe4, 0xb6, 0x7, 0xba, 0x49, 0x59, 0xd0, 0x41, 0x61}}
	return a, nil
}

var _transactionsPseudoRandomGeneratorSetup_prgCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x51\x4f\x4b\xc3\x50\x0c\xbf\xf7\x53\xc4\x1d\xb4\x85\xb9\x5d\x44\xa4\x4c\x61\xa7\xe2\xc9\xd2\xe1\x49\x3c\xc4\x36\xdb\x1e\x76\x79\x25\x49\x07\x22\xfb\xee\xf2\xba\x39\x6c\xdd\x60\xef\x50\xda\xbe\xdf\xdf\xc4\x6d\x1a\x2f\x06\xa3\x5c\xa9\xad\x7c\x81\x5c\xf9\x4d\x46\x4c\x82\xe6\x65\x14\x45\xd3\xe9\x14\x16\xb8\x25\x05\xe4\x0a\x6a\xc7\x9f\x0a\x08\x93\xbc\xc8\x40\x48\x7d\x2b\x25\x81\x63\xb0\x35\x81\xba\x15\x93\xdc\x28\xa8\x79\xc1\x15\x75\x8c\xa6\xfd\xa8\x5d\x09\x8c\x1b\xd2\x06\x4b\x0a\x82\x91\x09\xb2\x62\x69\xce\x73\xac\x44\x55\x0a\x6f\xaf\xcf\x6c\x0f\xef\x63\x50\xac\x2d\x85\xf0\x75\x7f\x97\xc0\x77\x04\x00\xd0\x08\x35\x28\x14\xef\x0d\x52\x98\xb7\xb6\x9e\x97\xa5\x6f\xd9\x7e\x21\xe1\xb8\xe5\x21\xc2\xc4\xbe\x1a\x8a\xd1\x52\x38\xd9\x6a\xb2\xd8\xe7\xcb\xd1\xd6\x09\x5c\x3d\x02\xbb\xfa\x8f\x4c\x38\x42\xd6\x0a\x1f\x7f\xed\x8e\x6f\x07\x03\xc5\x2d\xc5\x3d\x06\xcc\x6e\xcf\xb8\x95\x42\x68\x94\x17\xd9\x80\x10\xd4\xba\xf9\xbd\x2c\xf7\x14\x26\xd5\x14\xc2\x3c\xc6\xff\x91\xdd\x58\xc2\xb3\x77\x95\xf4\x91\xe6\x2f\xa8\x7c\x64\x24\xc3\x56\x61\xbb\xb3\xeb\xd3\x02\x79\x91\x3d\xf5\x0b\x9c\xc1\x75\xfb\x0e\x3e\x83\x68\x28\x2b\xba\x64\x23\x83\x78\xbb\x68\xf7\x13\x00\x00\xff\xff\x06\x15\xf9\x4b\xa3\x02\x00\x00"

func transactionsPseudoRandomGeneratorSetup_prgCdcBytes() ([]byte, error) {
	return bindataRead(
		_transactionsPseudoRandomGeneratorSetup_prgCdc,
		"transactions/pseudo-random-generator/setup_prg.cdc",
	)
}

func transactionsPseudoRandomGeneratorSetup_prgCdc() (*asset, error) {
	bytes, err := transactionsPseudoRandomGeneratorSetup_prgCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transactions/pseudo-random-generator/setup_prg.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf4, 0x8e, 0xc1, 0x8a, 0x68, 0x90, 0x27, 0x2, 0x89, 0xa5, 0x88, 0x1d, 0xdc, 0xb4, 0xc3, 0xf0, 0x99, 0x95, 0x5, 0x25, 0x87, 0x4c, 0x32, 0xbd, 0xeb, 0x87, 0x35, 0x51, 0x88, 0xb2, 0xa9, 0x6f}}
	return a, nil
}

var _transactionsRandomBeaconHistoryHeartbeatCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x3f\x6f\xdc\x30\x0c\xc5\x77\x7f\x0a\xe2\x86\xd6\xb7\x9c\xd7\xc2\x68\x1a\x24\xe9\x90\x2e\x4d\x91\xa4\x53\xd1\x41\x96\x79\x67\x22\xb2\xe8\x52\x54\x82\x43\x91\xef\x5e\x58\xfe\x53\x1b\xb8\x1c\xa2\xc9\x02\xc5\xf7\x7e\x7a\x32\xa9\xed\x58\x14\x36\xf7\xc6\xd7\xdc\x5e\xa3\xb1\xec\x6f\x29\x28\xcb\x71\x93\x65\x45\x51\xc0\x0d\xb7\x2d\x69\x00\x6d\x10\x02\x47\xb1\x08\xbc\x07\x49\xe7\x3d\x86\x00\x7b\x96\x54\x14\xfc\x13\x31\x28\xd6\x50\x39\xb6\x4f\xd0\x20\x1d\x1a\x85\xbd\x70\x9b\xea\x27\x2c\xe0\x16\x8d\x68\x85\x46\x7b\xab\x64\xf7\x9d\x15\x4b\xb0\x93\x29\x9f\x6a\xdb\xcd\x6d\xf0\x42\xce\x41\x85\x7d\x43\xe7\xb0\x37\x27\x3f\xa0\x1e\x83\x62\x0b\xb6\x89\xfe\x09\x54\x8c\x0f\xc6\x2a\xb1\x87\xea\x38\xd4\x51\x9e\xc9\x62\xf2\x34\xd6\x72\xf4\xba\x83\xc7\x86\xc2\xea\x30\x05\x20\x6f\x5d\xac\xb1\x4e\xf7\x24\x4f\x4a\xc6\x81\x62\x50\xf2\x07\xe8\xa2\x74\x1c\x30\x40\xf4\x4a\x6e\x4c\xe5\x21\xa5\x34\xb2\xe6\xdb\x5e\xc4\x3c\x1b\x72\xa6\x72\x08\xe4\x07\x4b\x0f\xd8\x46\x67\x94\x05\xaa\x48\xae\x4e\x01\x2c\xac\xf3\xc0\x52\xc2\xaf\x9f\xdf\xbc\x7e\xfa\xbd\x85\xbf\x19\x00\x40\x27\xd8\x19\xc1\x7c\x84\xbf\x1a\xb8\x4b\xb8\x8a\xda\x8c\x9b\xe9\x68\xbf\x8a\x02\xae\x59\x84\x5f\xde\xca\x7f\x11\xe4\x3d\x8e\x8f\x3b\xbf\x57\xa0\x83\xef\xef\x38\x9a\x4d\x29\xcd\xea\x0e\x75\xbc\xf0\x4a\xf3\xbf\xe4\x05\xac\x39\x77\x55\x82\xf9\xfc\xe1\x2c\xc9\x97\x7c\x76\x98\x56\x8f\x54\x9e\xe7\x7f\x50\x16\x73\xc0\x1f\x46\x9b\x55\xfb\x16\x2e\x2f\xa1\x33\x9e\x6c\xbe\xb9\xe1\xe8\x6a\xff\x51\x61\xe0\x78\x67\x20\x9b\x6d\xb6\x0c\xf4\xf1\xee\xeb\xdd\x72\xdf\xa7\xb0\x7c\x2a\xb8\x38\xfd\x17\xac\x44\x86\x99\x7a\x7b\xa4\xcc\x50\xb3\x51\x04\xbd\x0e\xf3\x34\x8c\xd3\xac\x72\x2e\xf9\x5d\x33\x7d\xe5\x27\x58\xca\x9e\x77\x9b\x84\x5e\xb3\xd7\xec\x5f\x00\x00\x00\xff\xff\x83\x51\xe6\x3c\x02\x04\x00\x00"

func transactionsRandomBeaconHistoryHeartbeatCdcBytes() ([]byte, error) {
	return bindataRead(
		_transactionsRandomBeaconHistoryHeartbeatCdc,
		"transactions/random-beacon-history/heartbeat.cdc",
	)
}

func transactionsRandomBeaconHistoryHeartbeatCdc() (*asset, error) {
	bytes, err := transactionsRandomBeaconHistoryHeartbeatCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transactions/random-beacon-history/heartbeat.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9b, 0x5d, 0x24, 0xb5, 0x5a, 0xb4, 0x43, 0x26, 0x46, 0x51, 0xa3, 0xe8, 0x1, 0x9, 0x6d, 0x5d, 0xcf, 0x2, 0x89, 0x7, 0xc9, 0x20, 0xe2, 0xff, 0xce, 0x78, 0xe, 0xc4, 0xaf, 0xf5, 0xbf, 0xec}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"scripts/pseudo-random-generator/next_uint64_from_address.cdc": scriptsPseudoRandomGeneratorNext_uint64_from_addressCdc,
	"scripts/pseudo-random-generator/next_uint64_new_prg.cdc":      scriptsPseudoRandomGeneratorNext_uint64_new_prgCdc,
	"transactions/coin-toss/0_commit_coin_toss.cdc":                transactionsCoinToss0_commit_coin_tossCdc,
	"transactions/coin-toss/1_reveal_coin_toss.cdc":                transactionsCoinToss1_reveal_coin_tossCdc,
	"transactions/flow-token/transfer_flow.cdc":                    transactionsFlowTokenTransfer_flowCdc,
	"transactions/pseudo-random-generator/next_uint64.cdc":         transactionsPseudoRandomGeneratorNext_uint64Cdc,
	"transactions/pseudo-random-generator/setup_prg.cdc":           transactionsPseudoRandomGeneratorSetup_prgCdc,
	"transactions/random-beacon-history/heartbeat.cdc":             transactionsRandomBeaconHistoryHeartbeatCdc,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"scripts": {nil, map[string]*bintree{
		"pseudo-random-generator": {nil, map[string]*bintree{
			"next_uint64_from_address.cdc": {scriptsPseudoRandomGeneratorNext_uint64_from_addressCdc, map[string]*bintree{}},
			"next_uint64_new_prg.cdc": {scriptsPseudoRandomGeneratorNext_uint64_new_prgCdc, map[string]*bintree{}},
		}},
	}},
	"transactions": {nil, map[string]*bintree{
		"coin-toss": {nil, map[string]*bintree{
			"0_commit_coin_toss.cdc": {transactionsCoinToss0_commit_coin_tossCdc, map[string]*bintree{}},
			"1_reveal_coin_toss.cdc": {transactionsCoinToss1_reveal_coin_tossCdc, map[string]*bintree{}},
		}},
		"flow-token": {nil, map[string]*bintree{
			"transfer_flow.cdc": {transactionsFlowTokenTransfer_flowCdc, map[string]*bintree{}},
		}},
		"pseudo-random-generator": {nil, map[string]*bintree{
			"next_uint64.cdc": {transactionsPseudoRandomGeneratorNext_uint64Cdc, map[string]*bintree{}},
			"setup_prg.cdc": {transactionsPseudoRandomGeneratorSetup_prgCdc, map[string]*bintree{}},
		}},
		"random-beacon-history": {nil, map[string]*bintree{
			"heartbeat.cdc": {transactionsRandomBeaconHistoryHeartbeatCdc, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
