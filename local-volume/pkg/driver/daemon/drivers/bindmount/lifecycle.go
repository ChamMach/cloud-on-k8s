package bindmount

import (
	"io/ioutil"
	"os"
	"path"
)

// ListKnownPVs lists the volumes path to find the existing PVs
func (d *Driver) ListKnownPVs() ([]string, error) {
	fileinfos, err := ioutil.ReadDir(ContainerMountPath)
	if err != nil {
		return nil, err
	}

	knownNames := make([]string, len(fileinfos))
	for i, fileinfo := range fileinfos {
		knownNames[i] = fileinfo.Name()
	}

	return knownNames, nil
}

// Purge recursively deletes the local volume directory
func (d *Driver) Purge(pvName string) error {
	return os.RemoveAll(path.Join(ContainerMountPath, pvName))
}
