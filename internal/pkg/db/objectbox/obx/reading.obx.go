// Code generated by ObjectBox; DO NOT EDIT.

package obx

import (
	. "github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type reading_EntityInfo struct {
	Id  objectbox.TypeId
	Uid uint64
}

var ReadingBinding = reading_EntityInfo{
	Id:  1,
	Uid: 3265861363892097666,
}

// Reading_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Reading_ = struct {
	Id       *objectbox.PropertyStringUint64
	Pushed   *objectbox.PropertyInt64
	Created  *objectbox.PropertyInt64
	Origin   *objectbox.PropertyInt64
	Modified *objectbox.PropertyInt64
	Device   *objectbox.PropertyString
	Name     *objectbox.PropertyString
	Value    *objectbox.PropertyString
}{
	Id: &objectbox.PropertyStringUint64{
		Property: &objectbox.Property{
			Id: 1,
		},
	},
	Pushed: &objectbox.PropertyInt64{
		Property: &objectbox.Property{
			Id: 2,
		},
	},
	Created: &objectbox.PropertyInt64{
		Property: &objectbox.Property{
			Id: 3,
		},
	},
	Origin: &objectbox.PropertyInt64{
		Property: &objectbox.Property{
			Id: 4,
		},
	},
	Modified: &objectbox.PropertyInt64{
		Property: &objectbox.Property{
			Id: 5,
		},
	},
	Device: &objectbox.PropertyString{
		Property: &objectbox.Property{
			Id: 6,
		},
	},
	Name: &objectbox.PropertyString{
		Property: &objectbox.Property{
			Id: 7,
		},
	},
	Value: &objectbox.PropertyString{
		Property: &objectbox.Property{
			Id: 8,
		},
	},
}

// GeneratorVersion is called by the ObjectBox to verify the compatibility of the generator used to generate this code
func (reading_EntityInfo) GeneratorVersion() int {
	return 1
}

// AddToModel is called by the ObjectBox during model build
func (reading_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Reading", 1, 3265861363892097666)
	model.Property("Id", objectbox.PropertyType_Long, 1, 2576545678551266117)
	model.PropertyFlags(objectbox.PropertyFlags_ID)
	model.Property("Pushed", objectbox.PropertyType_Long, 2, 3103381566796494532)
	model.Property("Created", objectbox.PropertyType_Long, 3, 4633520858368257570)
	model.Property("Origin", objectbox.PropertyType_Long, 4, 8768181063589245434)
	model.Property("Modified", objectbox.PropertyType_Long, 5, 5309626729124006671)
	model.Property("Device", objectbox.PropertyType_String, 6, 5429580072525114953)
	model.Property("Name", objectbox.PropertyType_String, 7, 6236109473949377403)
	model.Property("Value", objectbox.PropertyType_String, 8, 9137277974737209706)
	model.EntityLastPropertyId(8, 9137277974737209706)
}

// GetId is called by the ObjectBox during Put operations to check for existing ID on an object
func (reading_EntityInfo) GetId(object interface{}) (uint64, error) {
	if len(object.(*Reading).Id) == 0 {
		return 0, nil
	} else {
		return strconv.ParseUint(string(object.(*Reading).Id), 10, 64)
	}
}

// SetId is called by the ObjectBox during Put to update an ID on an object that has just been inserted
func (reading_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Reading).Id = bson.ObjectId(strconv.FormatUint(id, 10))
	return nil
}

// Flatten is called by the ObjectBox to transform an object to a FlatBuffer
func (reading_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) {
	obj := object.(*Reading)
	var offsetDevice = fbutils.CreateStringOffset(fbb, obj.Device)
	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetValue = fbutils.CreateStringOffset(fbb, obj.Value)

	// build the FlatBuffers object
	fbb.StartObject(8)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetInt64Slot(fbb, 1, obj.Pushed)
	fbutils.SetInt64Slot(fbb, 2, obj.Created)
	fbutils.SetInt64Slot(fbb, 3, obj.Origin)
	fbutils.SetInt64Slot(fbb, 4, obj.Modified)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetDevice)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetName)
	fbutils.SetUOffsetTSlot(fbb, 7, offsetValue)
}

// ToObject is called by the ObjectBox to load an object from a FlatBuffer
func (reading_EntityInfo) ToObject(bytes []byte) interface{} {
	table := &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	return &Reading{
		Id:       bson.ObjectId(strconv.FormatUint(table.GetUint64Slot(4, 0), 10)),
		Pushed:   table.GetInt64Slot(6, 0),
		Created:  table.GetInt64Slot(8, 0),
		Origin:   table.GetInt64Slot(10, 0),
		Modified: table.GetInt64Slot(12, 0),
		Device:   fbutils.GetStringSlot(table, 14),
		Name:     fbutils.GetStringSlot(table, 16),
		Value:    fbutils.GetStringSlot(table, 18),
	}
}

// MakeSlice is called by the ObjectBox to construct a new slice to hold the read objects
func (reading_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Reading, 0, capacity)
}

// AppendToSlice is called by the ObjectBox to fill the slice of the read objects
func (reading_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]*Reading), object.(*Reading))
}

// Box provides CRUD access to Reading objects
type ReadingBox struct {
	*objectbox.Box
}

// BoxForReading opens a box of Reading objects
func BoxForReading(ob *objectbox.ObjectBox) *ReadingBox {
	return &ReadingBox{
		Box: ob.InternalBox(1),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Reading.Id property on the passed object will be assigned the new ID as well.
func (box *ReadingBox) Put(object *Reading) (string, error) {
	if id, err := box.Box.Put(object); err != nil {
		return "", err
	} else {
		return strconv.FormatUint(id, 10), nil
	}
}

// PutAsync asynchronously inserts/updates a single object.
// When inserting, the Reading.Id property on the passed object will be assigned the new ID as well.
//
// It's executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "Put & Forget:" you gain faster puts as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
//
// In situations with (extremely) high async load, this method may be throttled (~1ms) or delayed (<1s).
// In the unlikely event that the object could not be enqueued after delaying, an error will be returned.
//
// Note that this method does not give you hard durability guarantees like the synchronous Put provides.
// There is a small time window (typically 3 ms) in which the data may not have been committed durably yet.
func (box *ReadingBox) PutAsync(object *Reading) (string, error) {
	if id, err := box.Box.PutAsync(object); err != nil {
		return "", err
	} else {
		return strconv.FormatUint(id, 10), nil
	}
}

// PutAll inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Reading.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Reading.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *ReadingBox) PutAll(objects []*Reading) ([]string, error) {
	ids, err := box.Box.PutAll(objects)
	if err != nil || len(ids) == 0 {
		return []string{}, err
	}

	var stringIds = make([]string, len(ids))
	for i, id := range ids {
		stringIds[i] = strconv.FormatUint(id, 10)
	}

	return stringIds, nil
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *ReadingBox) Get(id string) (*Reading, error) {
	idUint64, parseErr := strconv.ParseUint(id, 10, 64)
	if parseErr != nil {
		return nil, parseErr
	}

	object, err := box.Box.Get(idUint64)

	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Reading), nil
}

// Get reads all stored objects
func (box *ReadingBox) GetAll() ([]*Reading, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Reading), nil
}

// Remove deletes a single object
func (box *ReadingBox) Remove(object *Reading) (err error) {
	idUint64, parseErr := strconv.ParseUint(string(object.Id), 10, 64)
	if parseErr != nil {
		return parseErr
	}

	return box.Box.Remove(idUint64)
}

// Creates a query with the given conditions. Use the fields of the Reading_ struct to create conditions.
// Keep the *ReadingQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *ReadingBox) Query(conditions ...objectbox.Condition) *ReadingQuery {
	return &ReadingQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Reading_ struct to create conditions.
// Keep the *ReadingQuery if you intend to execute the query multiple times.
func (box *ReadingBox) QueryOrError(conditions ...objectbox.Condition) (*ReadingQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &ReadingQuery{query}, nil
	}
}

// Query provides a way to search stored objects
//
// For example, you can find all Reading which Id is either 42 or 47:
// 		box.Query(Reading_.Id.In(42, 47)).Find()
type ReadingQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *ReadingQuery) Find() ([]*Reading, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Reading), nil
}