package zookeeperx

import "testing"

const ROOT_PATH = "/k0"

func TestExport(t *testing.T) {
	Export(ROOT_PATH)

}

func TestDelete0(t *testing.T) {
	Delete(ROOT_PATH)
}

func TestDelete1(t *testing.T) {
	Delete("/k0/k01")
}

func TestImport(t *testing.T) {
	ImportFromFile()
}
