package zookeeperx

import "testing"

const ROOT_PATH = "/platform"

func TestDelete(t *testing.T) {
	Delete("/parent/path/to/delete")
}

func TestImportFromFile(t *testing.T) {
	ImportFromFile("/tmp/local-zk-export.txt.bak", "")
	Export("/platform/environment/idc.0001/time_series", "/tmp/local-zk-export.txt")
}

func TestReImportWithParent(t *testing.T) {
	root := "/platform/project/GB04"
	Delete(root)
	ImportFromFile("/tmp/local-zk-export.txt.bak", "")
	Export(root, "/tmp/local-zk-export.txt")
}

func TestExport(t *testing.T) {
	// /platform/environment/idc.0001/influxdb
	Export("/platform/project/GB04", "/tmp/local-zk-export.txt")
}
