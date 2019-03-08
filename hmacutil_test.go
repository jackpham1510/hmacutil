package hmacutil

import "testing"

var (
	secretKey = "secret_key"
	data      = "github.com/tiendung1510/hmacutil"
)

func TestHexStringEncode(t *testing.T) {
	var mac string
	testcases := []struct {
		algo, expect string
	}{
		{MD5, "42223b2dfbcd5fb9aa7637edc87a6180"},
		{SHA1, "bb0d06319c8f8a320ffb654f994b6c8071e5c785"},
		{SHA256, "de0e2263c35c1b8e09acd85fc23fb065135b9d1d88f7574c452f6770cd3317a2"},
		{SHA512, "57d89fb36dd6a4b4253447bf57062ed9cc6bfc036ff5bef8d95694476213f499e32abebf052602f37bdd571ef1e2bad797501b1584e5873f0f547d29e462ae57"},
	}

	for _, testcase := range testcases {
		mac = HexStringEncode(testcase.algo, secretKey, data)
		if mac != testcase.expect {
			t.Errorf("HexStringEncode(%v) = %v; expect %v", testcase.algo, mac, testcase.expect)
		}
	}
}
