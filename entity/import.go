package entity

type ImportState int32

const (
	ImportState_ImportPending      ImportState = 0
	ImportState_ImportFailed       ImportState = 1
	ImportState_ImportStarted      ImportState = 2
	ImportState_ImportDownloaded   ImportState = 3
	ImportState_ImportParsed       ImportState = 4
	ImportState_ImportPersisted    ImportState = 5
	ImportState_ImportCompleted    ImportState = 6
	ImportState_ImportAllocSegment ImportState = 10
)

type ImportTaskState struct {
	Id           int64
	State        ImportState
	RowCount     int64
	IdList       []int64
	Infos        map[string]string
	CollectionId int64
	SegmentIds   []int64
}
