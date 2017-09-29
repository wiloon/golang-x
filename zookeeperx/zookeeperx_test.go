package zookeeperx

import "testing"

const ROOT_PATH = "/platform"

func TestDelete(t *testing.T) {
	Delete("/parent/path/to/delete")
}

func TestImportWithParent(t *testing.T) {
	ImportFromFile("/tmp/local-zk-export.txt.bak", "")
}

func TestExport(t *testing.T) {
	Export("/platform/environment/idc.0001/mysql/user-center", "/tmp/local-zk-export.txt")
}
