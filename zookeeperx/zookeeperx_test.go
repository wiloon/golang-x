package zookeeperx

import "testing"

const ROOT_PATH = "/platform"

func TestDelete(t *testing.T) {
	Delete("/platform/environment/idc.0001/redis/cluster/hosts")
}

func TestImportWithParent(t *testing.T) {
	ImportFromFile("/tmp/import.txt", "")
}

func TestExport(t *testing.T) {
	Export("/platform/environment/idc.0001/mysql/user-center", "/tmp/local-zk-export.txt")
}
