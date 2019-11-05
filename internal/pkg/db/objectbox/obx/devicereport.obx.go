// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package obx

import (
	"errors"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	. "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type deviceReport_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var DeviceReportBinding = deviceReport_EntityInfo{
	Entity: objectbox.Entity{
		Id: 7,
	},
	Uid: 2987965348377516725,
}

// DeviceReport_ contains type-based Property helpers to facilitate some common operations such as Queries.
var DeviceReport_ = struct {
	Created  *objectbox.PropertyInt64
	Modified *objectbox.PropertyInt64
	Origin   *objectbox.PropertyInt64
	Id       *objectbox.PropertyUint64
	Name     *objectbox.PropertyString
	Device   *objectbox.PropertyString
	Action   *objectbox.PropertyString
	Expected *objectbox.PropertyStringVector
}{
	Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &DeviceReportBinding.Entity,
		},
	},
	Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &DeviceReportBinding.Entity,
		},
	},
	Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &DeviceReportBinding.Entity,
		},
	},
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &DeviceReportBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &DeviceReportBinding.Entity,
		},
	},
	Device: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &DeviceReportBinding.Entity,
		},
	},
	Action: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &DeviceReportBinding.Entity,
		},
	},
	Expected: &objectbox.PropertyStringVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &DeviceReportBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (deviceReport_EntityInfo) GeneratorVersion() int {
	return 4
}

// AddToModel is called by ObjectBox during model build
func (deviceReport_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("DeviceReport", 7, 2987965348377516725)
	model.Property("Created", 6, 1, 144277113410594238)
	model.Property("Modified", 6, 2, 6752056937527638827)
	model.Property("Origin", 6, 3, 9138721759452247071)
	model.Property("Id", 6, 4, 8821067739042639817)
	model.PropertyFlags(1)
	model.Property("Name", 9, 5, 4017589850413680442)
	model.PropertyFlags(32)
	model.PropertyIndex(8, 6567592991408811323)
	model.Property("Device", 9, 6, 4864361410797328287)
	model.Property("Action", 9, 7, 7348386595663429576)
	model.Property("Expected", 30, 8, 3826804133873897010)
	model.EntityLastPropertyId(8, 3826804133873897010)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (deviceReport_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*DeviceReport); ok {
		return objectbox.StringIdConvertToDatabaseValue(obj.Id)
	} else {
		return objectbox.StringIdConvertToDatabaseValue(object.(DeviceReport).Id)
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (deviceReport_EntityInfo) SetId(object interface{}, id uint64) error {
	if obj, ok := object.(*DeviceReport); ok {
		var err error
		obj.Id, err = objectbox.StringIdConvertToEntityProperty(id)
		return err
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(DeviceReport).Id
		return nil
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (deviceReport_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (deviceReport_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	var obj *DeviceReport
	if objPtr, ok := object.(*DeviceReport); ok {
		obj = objPtr
	} else {
		objVal := object.(DeviceReport)
		obj = &objVal
	}

	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetDevice = fbutils.CreateStringOffset(fbb, obj.Device)
	var offsetAction = fbutils.CreateStringOffset(fbb, obj.Action)
	var offsetExpected = fbutils.CreateStringVectorOffset(fbb, obj.Expected)

	// build the FlatBuffers object
	fbb.StartObject(8)
	fbutils.SetInt64Slot(fbb, 0, obj.Timestamps.Created)
	fbutils.SetInt64Slot(fbb, 1, obj.Timestamps.Modified)
	fbutils.SetInt64Slot(fbb, 2, obj.Timestamps.Origin)
	fbutils.SetUint64Slot(fbb, 3, id)
	fbutils.SetUOffsetTSlot(fbb, 4, offsetName)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetDevice)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetAction)
	fbutils.SetUOffsetTSlot(fbb, 7, offsetExpected)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (deviceReport_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'DeviceReport' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	propId, err := objectbox.StringIdConvertToEntityProperty(fbutils.GetUint64Slot(table, 10))
	if err != nil {
		return nil, errors.New("converter objectbox.StringIdConvertToEntityProperty() failed on DeviceReport.Id: " + err.Error())
	}

	return &DeviceReport{
		Timestamps: models.Timestamps{
			Created:  fbutils.GetInt64Slot(table, 4),
			Modified: fbutils.GetInt64Slot(table, 6),
			Origin:   fbutils.GetInt64Slot(table, 8),
		},
		Id:       propId,
		Name:     fbutils.GetStringSlot(table, 12),
		Device:   fbutils.GetStringSlot(table, 14),
		Action:   fbutils.GetStringSlot(table, 16),
		Expected: fbutils.GetStringVectorSlot(table, 18),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (deviceReport_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]DeviceReport, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (deviceReport_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]DeviceReport), DeviceReport{})
	}
	return append(slice.([]DeviceReport), *object.(*DeviceReport))
}

// Box provides CRUD access to DeviceReport objects
type DeviceReportBox struct {
	*objectbox.Box
}

// BoxForDeviceReport opens a box of DeviceReport objects
func BoxForDeviceReport(ob *objectbox.ObjectBox) *DeviceReportBox {
	return &DeviceReportBox{
		Box: ob.InternalBox(7),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the DeviceReport.Id property on the passed object will be assigned the new ID as well.
func (box *DeviceReportBox) Put(object *DeviceReport) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the DeviceReport.Id property on the passed object will be assigned the new ID as well.
func (box *DeviceReportBox) Insert(object *DeviceReport) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *DeviceReportBox) Update(object *DeviceReport) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *DeviceReportBox) PutAsync(object *DeviceReport) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the DeviceReport.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the DeviceReport.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *DeviceReportBox) PutMany(objects []DeviceReport) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *DeviceReportBox) Get(id uint64) (*DeviceReport, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*DeviceReport), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is an empty object
func (box *DeviceReportBox) GetMany(ids ...uint64) ([]DeviceReport, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]DeviceReport), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *DeviceReportBox) GetManyExisting(ids ...uint64) ([]DeviceReport, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]DeviceReport), nil
}

// GetAll reads all stored objects
func (box *DeviceReportBox) GetAll() ([]DeviceReport, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]DeviceReport), nil
}

// Remove deletes a single object
func (box *DeviceReportBox) Remove(object *DeviceReport) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *DeviceReportBox) RemoveMany(objects ...*DeviceReport) (uint64, error) {
	var ids = make([]uint64, len(objects))
	var err error
	for k, object := range objects {
		ids[k], err = objectbox.StringIdConvertToDatabaseValue(object.Id)
		if err != nil {
			return 0, errors.New("converter objectbox.StringIdConvertToDatabaseValue() failed on DeviceReport.Id: " + err.Error())
		}
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the DeviceReport_ struct to create conditions.
// Keep the *DeviceReportQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *DeviceReportBox) Query(conditions ...objectbox.Condition) *DeviceReportQuery {
	return &DeviceReportQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the DeviceReport_ struct to create conditions.
// Keep the *DeviceReportQuery if you intend to execute the query multiple times.
func (box *DeviceReportBox) QueryOrError(conditions ...objectbox.Condition) (*DeviceReportQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &DeviceReportQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See DeviceReportAsyncBox for more information.
func (box *DeviceReportBox) Async() *DeviceReportAsyncBox {
	return &DeviceReportAsyncBox{AsyncBox: box.Box.Async()}
}

// DeviceReportAsyncBox provides asynchronous operations on DeviceReport objects.
//
// Asynchronous operations are executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "execute & forget:" you gain faster put/remove operations as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
// In situations with (extremely) high async load, an async method may be throttled (~1ms) or delayed up to 1 second.
// In the unlikely event that the object could still not be enqueued (full queue), an error will be returned.
//
// Note that async methods do not give you hard durability guarantees like the synchronous Box provides.
// There is a small time window in which the data may not have been committed durably yet.
type DeviceReportAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForDeviceReport creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use DeviceReportBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForDeviceReport(ob *objectbox.ObjectBox, timeoutMs uint64) *DeviceReportAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 7, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 7: %s" + err.Error())
	}
	return &DeviceReportAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *DeviceReportAsyncBox) Put(object *DeviceReport) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *DeviceReportAsyncBox) Insert(object *DeviceReport) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *DeviceReportAsyncBox) Update(object *DeviceReport) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *DeviceReportAsyncBox) Remove(object *DeviceReport) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all DeviceReport which Id is either 42 or 47:
// 		box.Query(DeviceReport_.Id.In(42, 47)).Find()
type DeviceReportQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *DeviceReportQuery) Find() ([]DeviceReport, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]DeviceReport), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *DeviceReportQuery) Offset(offset uint64) *DeviceReportQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *DeviceReportQuery) Limit(limit uint64) *DeviceReportQuery {
	query.Query.Limit(limit)
	return query
}
