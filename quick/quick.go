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

package quick

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sync"

	"github.com/fatih/structs"
	"github.com/tintedwild/pkg/v3/safe"
	etcd "go.etcd.io/etcd/client/v3"
)

// Config - generic config interface functions
type Config interface {
	String() string
	Version() string
	Save(string) error
	Load(string) error
	Data() interface{}
	Diff(Config) ([]structs.Field, error)
	DeepDiff(Config) ([]structs.Field, error)
}

// config - implements quick.Config interface
type config struct {
	data interface{}
	clnt *etcd.Client
	lock *sync.RWMutex
}

// Version returns the current config file format version
func (d config) Version() string {
	st := structs.New(d.data)
	f := st.Field("Version")
	return f.Value().(string)
}

// String converts JSON config to printable string
func (d config) String() string {
	configBytes, _ := json.MarshalIndent(d.data, "", "\t")
	return string(configBytes)
}

// Save writes config data to a file. Data format
// is selected based on file extension or JSON if
// not provided.
func (d config) Save(filename string) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.clnt != nil {
		return saveFileConfigEtcd(filename, d.clnt, d.data)
	}

	// Backup if given file exists
	oldData, err := os.ReadFile(filename)
	if err != nil {
		// Ignore if file does not exist.
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		// Save read data to the backup file.
		backupFilename := filename + ".old"
		if err = writeFile(backupFilename, oldData); err != nil {
			return err
		}
	}

	// Save data.
	return saveFileConfig(filename, d.data)
}

// Load - loads config from file and merge with currently set values
// File content format is guessed from the file name extension, if not
// available, consider that we have JSON.
func (d config) Load(filename string) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.clnt != nil {
		return loadFileConfigEtcd(filename, d.clnt, d.data)
	}
	return loadFileConfig(filename, d.data)
}

// Data - grab internal data map for reading
func (d config) Data() interface{} {
	return d.data
}

// Diff  - list fields that are in A but not in B
func (d config) Diff(c Config) ([]structs.Field, error) {
	var fields []structs.Field

	currFields := structs.Fields(d.Data())
	newFields := structs.Fields(c.Data())

	var found bool
	for _, currField := range currFields {
		found = false
		for _, newField := range newFields {
			if reflect.DeepEqual(currField.Name(), newField.Name()) {
				found = true
			}
		}
		if !found {
			fields = append(fields, *currField)
		}
	}
	return fields, nil
}

// DeepDiff  - list fields in A that are missing or not equal to fields in B
func (d config) DeepDiff(c Config) ([]structs.Field, error) {
	var fields []structs.Field

	currFields := structs.Fields(d.Data())
	newFields := structs.Fields(c.Data())

	var found bool
	for _, currField := range currFields {
		found = false
		for _, newField := range newFields {
			if reflect.DeepEqual(currField.Value(), newField.Value()) {
				found = true
			}
		}
		if !found {
			fields = append(fields, *currField)
		}
	}
	return fields, nil
}

// CheckData - checks the validity of config data. Data should be of
// type struct and contain a string type field called "Version".
func CheckData(data interface{}) error {
	if !structs.IsStruct(data) {
		return fmt.Errorf("interface must be struct type")
	}

	st := structs.New(data)
	f, ok := st.FieldOk("Version")
	if !ok {
		return fmt.Errorf("struct ‘%s’ must have field ‘Version’", st.Name())
	}

	if f.Kind() != reflect.String {
		return fmt.Errorf("‘Version’ field in struct ‘%s’ must be a string type", st.Name())
	}

	return nil
}

// writeFile writes data to a file named by filename.
// If the file does not exist, writeFile creates it;
// otherwise writeFile truncates it before writing.
func writeFile(filename string, data []byte) error {
	safeFile, err := safe.CreateFile(filename)
	if err != nil {
		return err
	}
	_, err = safeFile.Write(data)
	if err != nil {
		return err
	}
	return safeFile.Close()
}

// GetVersion - extracts the version information.
func GetVersion(filename string, clnt *etcd.Client) (version string, err error) {
	var qc Config
	qc, err = LoadConfig(filename, clnt, &struct {
		Version string
	}{})
	if err != nil {
		return "", err
	}
	return qc.Version(), nil
}

// LoadConfig - loads json config from filename for the a given struct data
func LoadConfig(filename string, clnt *etcd.Client, data interface{}) (qc Config, err error) {
	qc, err = NewConfig(data, clnt)
	if err != nil {
		return nil, err
	}
	return qc, qc.Load(filename)
}

// SaveConfig - saves given configuration data into given file as JSON.
func SaveConfig(data interface{}, filename string, clnt *etcd.Client) (err error) {
	if err = CheckData(data); err != nil {
		return err
	}
	var qc Config
	qc, err = NewConfig(data, clnt)
	if err != nil {
		return err
	}
	return qc.Save(filename)
}

// NewConfig loads config from etcd client if provided, otherwise loads from a local filename.
// fails when all else fails.
func NewConfig(data interface{}, clnt *etcd.Client) (cfg Config, err error) {
	if err := CheckData(data); err != nil {
		return nil, err
	}

	d := new(config)
	d.data = data
	d.clnt = clnt
	d.lock = new(sync.RWMutex)
	return d, nil
}
