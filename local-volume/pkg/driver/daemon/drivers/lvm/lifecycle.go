package lvm

import log "github.com/sirupsen/logrus"

// ListKnownPVs lists the logical volumes from the volume group name to find the PV names
func (d *Driver) ListKnownPVs() ([]string, error) {
	vg, err := LookupVolumeGroup(d.options.VolumeGroupName)
	if err != nil {
		return nil, err
	}

	lvs, err := vg.ListLogicalVolumes()
	if err != nil {
		return nil, err
	}

	var knownNames []string
	for _, lv := range lvs {
		knownNames = append(knownNames, lv.name)
	}
	return knownNames, nil
}

// Purge deletes a logical volume
func (d *Driver) Purge(pvName string) error {
	vg, err := LookupVolumeGroup(d.options.VolumeGroupName)
	if err != nil {
		if err == ErrVolumeGroupNotFound {
			// we're deleting, missing volume group means the lv must be gone as well
			log.Infof("Volume group %s not found during purging of %s, skipping.", d.options.VolumeGroupName, pvName)
			return nil
		}
		return err
	}

	lv, err := vg.LookupLogicalVolume(pvName)
	if err != nil {
		if err == ErrLogicalVolumeNotFound {
			// we're deleting, so not found is fine.
			log.Infof("Logical volume %s not found during purging, skipping.", pvName)
			return nil
		}
		return err
	}

	return lv.Remove()
}
