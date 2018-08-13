package objectbox

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/flatcoredata"
	. "github.com/edgexfoundry/edgex-go/internal/pkg/objectbox"
	"github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/google/flatbuffers/go"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type ReadingBinding struct {
}

func (ReadingBinding) GetTypeId() TypeId {
	return 2
}

func (ReadingBinding) GetTypeName() string {
	return "Reading"
}

func (ReadingBinding) GetId(object interface{}) (id uint64, err error) {
	reading, ok := object.(*models.Reading)
	if !ok {
		// Programming error, OK to panic
		panic("Object has wrong type")
	}
	idString := string(reading.Id)
	if idString == "" {
		return 0, nil
	}
	return strconv.ParseUint(idString, 10, 64)
}

func (ReadingBinding) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) {
	flattenModelReading(object.(*models.Reading), fbb, id)
}

func flattenModelReading(reading *models.Reading, fbb *flatbuffers.Builder, id uint64) {
	offsetDevice := Unavailable
	if reading.Device != "" {
		offsetDevice = fbb.CreateString(reading.Device)
	}
	offsetName := Unavailable
	if reading.Name != "" {
		offsetName = fbb.CreateString(reading.Name)
	}
	offsetValue := Unavailable
	if reading.Value != "" {
		offsetValue = fbb.CreateString(reading.Value)
	}

	flatcoredata.ReadingStart(fbb)

	flatcoredata.ReadingAddId(fbb, id)
	flatcoredata.ReadingAddCreated(fbb, reading.Created)
	flatcoredata.ReadingAddOrigin(fbb, reading.Origin)
	flatcoredata.ReadingAddModified(fbb, reading.Modified)
	flatcoredata.ReadingAddPushed(fbb, reading.Pushed)

	if offsetDevice != Unavailable {
		flatcoredata.ReadingAddDevice(fbb, offsetDevice)
	}
	if offsetName != Unavailable {
		flatcoredata.ReadingAddName(fbb, offsetName)
	}
	if offsetValue != Unavailable {
		flatcoredata.ReadingAddValue(fbb, offsetValue)
	}
}

func (ReadingBinding) ToObject(bytes []byte) interface{} {
	flatReading := flatcoredata.GetRootAsReading(bytes, flatbuffers.UOffsetT(0))
	return toModelReading(flatReading)
}

func (ReadingBinding) AppendToSlice(slice interface{}, object interface{}) (sliceNew interface{}) {
	if slice == nil {
		slice = make([]models.Reading, 0, 16)
	}
	return append(slice.([]models.Reading), *object.(*models.Reading))
}

func toModelReadingFromBytes(bytesData []byte) *models.Reading {
	flatReading := flatcoredata.GetRootAsReading(bytesData, flatbuffers.UOffsetT(0))
	return toModelReading(flatReading)
}

func toModelReading(src *flatcoredata.Reading) *models.Reading {
	return &models.Reading{
		Id:       bson.ObjectId(strconv.FormatUint(src.Id(), 10)),
		Pushed:   src.Pushed(),
		Created:  src.Created(),
		Origin:   src.Origin(),
		Modified: src.Modified(),
		Device:   string(src.Device()),
		Name:     string(src.Name()),
		Value:    string(src.Value()),
	}

}
