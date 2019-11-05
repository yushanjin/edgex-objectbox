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

type transmission_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var TransmissionBinding = transmission_EntityInfo{
	Entity: objectbox.Entity{
		Id: 15,
	},
	Uid: 6516899534013799700,
}

// Transmission_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Transmission_ = struct {
	Created               *objectbox.PropertyInt64
	Modified              *objectbox.PropertyInt64
	Origin                *objectbox.PropertyInt64
	ID                    *objectbox.PropertyUint64
	Notification          *objectbox.RelationToOne
	Receiver              *objectbox.PropertyString
	Channel_Type          *objectbox.PropertyString
	Channel_MailAddresses *objectbox.PropertyStringVector
	Channel_Url           *objectbox.PropertyString
	Status                *objectbox.PropertyString
	ResendCount           *objectbox.PropertyInt
	Records               *objectbox.PropertyByteVector
}{
	Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &TransmissionBinding.Entity,
		},
	},
	Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &TransmissionBinding.Entity,
		},
	},
	Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &TransmissionBinding.Entity,
		},
	},
	ID: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &TransmissionBinding.Entity,
		},
	},
	Notification: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     5,
			Entity: &TransmissionBinding.Entity,
		},
		Target: &NotificationBinding.Entity,
	},
	Receiver: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &TransmissionBinding.Entity,
		},
	},
	Channel_Type: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &TransmissionBinding.Entity,
		},
	},
	Channel_MailAddresses: &objectbox.PropertyStringVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &TransmissionBinding.Entity,
		},
	},
	Channel_Url: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &TransmissionBinding.Entity,
		},
	},
	Status: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     10,
			Entity: &TransmissionBinding.Entity,
		},
	},
	ResendCount: &objectbox.PropertyInt{
		BaseProperty: &objectbox.BaseProperty{
			Id:     11,
			Entity: &TransmissionBinding.Entity,
		},
	},
	Records: &objectbox.PropertyByteVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     12,
			Entity: &TransmissionBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (transmission_EntityInfo) GeneratorVersion() int {
	return 4
}

// AddToModel is called by ObjectBox during model build
func (transmission_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Transmission", 15, 6516899534013799700)
	model.Property("Created", 6, 1, 874268777377311919)
	model.Property("Modified", 6, 2, 4563914448884285908)
	model.Property("Origin", 6, 3, 6396561716074782386)
	model.Property("ID", 6, 4, 7018017394194855826)
	model.PropertyFlags(1)
	model.Property("Notification", 11, 5, 2587965071970041075)
	model.PropertyFlags(8712)
	model.PropertyRelation("Notification", 18, 4309985151868307894)
	model.Property("Receiver", 9, 6, 6970133143502682397)
	model.Property("Channel_Type", 9, 7, 8315623827360567320)
	model.Property("Channel_MailAddresses", 30, 8, 2372831308292201526)
	model.Property("Channel_Url", 9, 9, 7096098715837417906)
	model.Property("Status", 9, 10, 161677121815195738)
	model.Property("ResendCount", 6, 11, 1458266355346453588)
	model.Property("Records", 23, 12, 107793340757305390)
	model.EntityLastPropertyId(12, 107793340757305390)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (transmission_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*Transmission); ok {
		return objectbox.StringIdConvertToDatabaseValue(obj.ID)
	} else {
		return objectbox.StringIdConvertToDatabaseValue(object.(Transmission).ID)
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (transmission_EntityInfo) SetId(object interface{}, id uint64) error {
	if obj, ok := object.(*Transmission); ok {
		var err error
		obj.ID, err = objectbox.StringIdConvertToEntityProperty(id)
		return err
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(Transmission).ID
		return nil
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (transmission_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	if rel := &object.(*Transmission).Notification; rel != nil {
		if rId, err := NotificationBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForNotification(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (transmission_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	var obj *Transmission
	if objPtr, ok := object.(*Transmission); ok {
		obj = objPtr
	} else {
		objVal := object.(Transmission)
		obj = &objVal
	}

	var propRecords []byte
	{
		var err error
		propRecords, err = transmissionRecordsJsonToDatabaseValue(obj.Records)
		if err != nil {
			return errors.New("converter transmissionRecordsJsonToDatabaseValue() failed on Transmission.Records: " + err.Error())
		}
	}

	var offsetReceiver = fbutils.CreateStringOffset(fbb, obj.Receiver)
	var offsetChannel_Type = fbutils.CreateStringOffset(fbb, string(obj.Channel.Type))
	var offsetChannel_MailAddresses = fbutils.CreateStringVectorOffset(fbb, obj.Channel.MailAddresses)
	var offsetChannel_Url = fbutils.CreateStringOffset(fbb, obj.Channel.Url)
	var offsetStatus = fbutils.CreateStringOffset(fbb, string(obj.Status))
	var offsetRecords = fbutils.CreateByteVectorOffset(fbb, propRecords)

	var rIdNotification uint64
	if rel := &obj.Notification; rel != nil {
		if rId, err := NotificationBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdNotification = rId
		}
	}

	// build the FlatBuffers object
	fbb.StartObject(12)
	fbutils.SetInt64Slot(fbb, 0, obj.Timestamps.Created)
	fbutils.SetInt64Slot(fbb, 1, obj.Timestamps.Modified)
	fbutils.SetInt64Slot(fbb, 2, obj.Timestamps.Origin)
	fbutils.SetUint64Slot(fbb, 3, id)
	fbutils.SetUint64Slot(fbb, 4, rIdNotification)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetReceiver)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetChannel_Type)
	fbutils.SetUOffsetTSlot(fbb, 7, offsetChannel_MailAddresses)
	fbutils.SetUOffsetTSlot(fbb, 8, offsetChannel_Url)
	fbutils.SetUOffsetTSlot(fbb, 9, offsetStatus)
	fbutils.SetInt64Slot(fbb, 10, int64(obj.ResendCount))
	fbutils.SetUOffsetTSlot(fbb, 11, offsetRecords)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (transmission_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'Transmission' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	propID, err := objectbox.StringIdConvertToEntityProperty(fbutils.GetUint64Slot(table, 10))
	if err != nil {
		return nil, errors.New("converter objectbox.StringIdConvertToEntityProperty() failed on Transmission.ID: " + err.Error())
	}

	propRecords, err := transmissionRecordsJsonToEntityProperty(fbutils.GetByteVectorSlot(table, 26))
	if err != nil {
		return nil, errors.New("converter transmissionRecordsJsonToEntityProperty() failed on Transmission.Records: " + err.Error())
	}

	var relNotification *Notification
	if rId := fbutils.GetUint64Slot(table, 12); rId > 0 {
		if rObject, err := BoxForNotification(ob).Get(rId); err != nil {
			return nil, err
		} else if rObject == nil {
			relNotification = &Notification{}
		} else {
			relNotification = rObject
		}
	} else {
		relNotification = &Notification{}
	}

	return &Transmission{
		Timestamps: models.Timestamps{
			Created:  fbutils.GetInt64Slot(table, 4),
			Modified: fbutils.GetInt64Slot(table, 6),
			Origin:   fbutils.GetInt64Slot(table, 8),
		},
		ID:           propID,
		Notification: *relNotification,
		Receiver:     fbutils.GetStringSlot(table, 14),
		Channel: models.Channel{
			Type:          models.ChannelType(fbutils.GetStringSlot(table, 16)),
			MailAddresses: fbutils.GetStringVectorSlot(table, 18),
			Url:           fbutils.GetStringSlot(table, 20),
		},
		Status:      models.TransmissionStatus(fbutils.GetStringSlot(table, 22)),
		ResendCount: fbutils.GetIntSlot(table, 24),
		Records:     propRecords,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (transmission_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]Transmission, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (transmission_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]Transmission), Transmission{})
	}
	return append(slice.([]Transmission), *object.(*Transmission))
}

// Box provides CRUD access to Transmission objects
type TransmissionBox struct {
	*objectbox.Box
}

// BoxForTransmission opens a box of Transmission objects
func BoxForTransmission(ob *objectbox.ObjectBox) *TransmissionBox {
	return &TransmissionBox{
		Box: ob.InternalBox(15),
	}
}

// Put synchronously inserts/updates a single object.
// In case the ID is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Transmission.ID property on the passed object will be assigned the new ID as well.
func (box *TransmissionBox) Put(object *Transmission) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the ID is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Transmission.ID property on the passed object will be assigned the new ID as well.
func (box *TransmissionBox) Insert(object *Transmission) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *TransmissionBox) Update(object *Transmission) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *TransmissionBox) PutAsync(object *Transmission) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case IDs are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Transmission.ID property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Transmission.ID assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *TransmissionBox) PutMany(objects []Transmission) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *TransmissionBox) Get(id uint64) (*Transmission, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Transmission), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is an empty object
func (box *TransmissionBox) GetMany(ids ...uint64) ([]Transmission, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]Transmission), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *TransmissionBox) GetManyExisting(ids ...uint64) ([]Transmission, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]Transmission), nil
}

// GetAll reads all stored objects
func (box *TransmissionBox) GetAll() ([]Transmission, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]Transmission), nil
}

// Remove deletes a single object
func (box *TransmissionBox) Remove(object *Transmission) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *TransmissionBox) RemoveMany(objects ...*Transmission) (uint64, error) {
	var ids = make([]uint64, len(objects))
	var err error
	for k, object := range objects {
		ids[k], err = objectbox.StringIdConvertToDatabaseValue(object.ID)
		if err != nil {
			return 0, errors.New("converter objectbox.StringIdConvertToDatabaseValue() failed on Transmission.ID: " + err.Error())
		}
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Transmission_ struct to create conditions.
// Keep the *TransmissionQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *TransmissionBox) Query(conditions ...objectbox.Condition) *TransmissionQuery {
	return &TransmissionQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Transmission_ struct to create conditions.
// Keep the *TransmissionQuery if you intend to execute the query multiple times.
func (box *TransmissionBox) QueryOrError(conditions ...objectbox.Condition) (*TransmissionQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &TransmissionQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See TransmissionAsyncBox for more information.
func (box *TransmissionBox) Async() *TransmissionAsyncBox {
	return &TransmissionAsyncBox{AsyncBox: box.Box.Async()}
}

// TransmissionAsyncBox provides asynchronous operations on Transmission objects.
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
type TransmissionAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForTransmission creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use TransmissionBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForTransmission(ob *objectbox.ObjectBox, timeoutMs uint64) *TransmissionAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 15, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 15: %s" + err.Error())
	}
	return &TransmissionAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the ID property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *TransmissionAsyncBox) Put(object *Transmission) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The ID property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *TransmissionAsyncBox) Insert(object *Transmission) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *TransmissionAsyncBox) Update(object *Transmission) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *TransmissionAsyncBox) Remove(object *Transmission) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all Transmission which ID is either 42 or 47:
// 		box.Query(Transmission_.ID.In(42, 47)).Find()
type TransmissionQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *TransmissionQuery) Find() ([]Transmission, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]Transmission), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *TransmissionQuery) Offset(offset uint64) *TransmissionQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *TransmissionQuery) Limit(limit uint64) *TransmissionQuery {
	query.Query.Limit(limit)
	return query
}
