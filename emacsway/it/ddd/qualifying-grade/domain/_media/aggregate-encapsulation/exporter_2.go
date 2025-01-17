package qualifying_grade_2

import (
    "time"
)

type Exportable[T any] interface {
    Export() T
}

type ExportableUint uint

func (e ExportableUint) Export() uint {
    return uint(e)
}

type RecognizerId ExportableUint
type MemberId ExportableUint
type Grade ExportableUint
type EndorsementCount ExportableUint

type RecognizerState struct {
    Id                        uint
    MemberId                  uint
    Grade                     uint
    AvailableEndorsementCount uint
    PendingEndorsementCount   uint
    Version                   uint
    CreatedAt                 time.Time
}

type Recognizer struct {
    id                        RecognizerId
    memberId                  MemberId
    grade                     Grade
    availableEndorsementCount EndorsementCount
    pendingEndorsementCount   EndorsementCount
    version                   uint
    createdAt                 time.Time
}

func (r Recognizer) Export() RecognizerState {
    return RecognizerState{
        r.id.Export(), r.memberId.Export(), r.grade.Export(),
        r.availableEndorsementCount.Export(),
        r.pendingEndorsementCount.Export(),
        r.version, r.createdAt,
    }
}
