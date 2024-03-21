package model

// !!! DO NOT EDIT THIS FILE

import (
	"context"
	"encoding/json"
	"github.com/iancoleman/strcase"
	"github.com/mylxsw/eloquent/query"
	"gopkg.in/guregu/null.v3"
	"time"
)

func init() {

	// AddChannelsGlobalScope assign a global scope to a model for soft delete
	AddGlobalScopeForChannels("soft_delete", func(builder query.Condition) {
		builder.WhereNull("deleted_at")
	})

}

// ChannelsN is a Channels object, all fields are nullable
type ChannelsN struct {
	original      *channelsOriginal
	channelsModel *ChannelsModel

	Id        null.Int    `json:"id"`
	Name      null.String `json:"name"`
	Type      null.String `json:"type"`
	Server    null.String `json:"server,omitempty"`
	Secret    null.String `json:"secret,omitempty"`
	MetaJson  null.String `json:"-"`
	CreatedAt null.Time
	UpdatedAt null.Time
	DeletedAt null.Time
}

// As convert object to other type
// dst must be a pointer to struct
func (inst *ChannelsN) As(dst interface{}) error {
	return query.Copy(inst, dst)
}

// SetModel set model for Channels
func (inst *ChannelsN) SetModel(channelsModel *ChannelsModel) {
	inst.channelsModel = channelsModel
}

// channelsOriginal is an object which stores original Channels from database
type channelsOriginal struct {
	Id        null.Int
	Name      null.String
	Type      null.String
	Server    null.String
	Secret    null.String
	MetaJson  null.String
	CreatedAt null.Time
	UpdatedAt null.Time
	DeletedAt null.Time
}

// Staled identify whether the object has been modified
func (inst *ChannelsN) Staled(onlyFields ...string) bool {
	if inst.original == nil {
		inst.original = &channelsOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			return true
		}
		if inst.Name != inst.original.Name {
			return true
		}
		if inst.Type != inst.original.Type {
			return true
		}
		if inst.Server != inst.original.Server {
			return true
		}
		if inst.Secret != inst.original.Secret {
			return true
		}
		if inst.MetaJson != inst.original.MetaJson {
			return true
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			return true
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			return true
		}
		if inst.DeletedAt != inst.original.DeletedAt {
			return true
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					return true
				}
			case "name":
				if inst.Name != inst.original.Name {
					return true
				}
			case "type":
				if inst.Type != inst.original.Type {
					return true
				}
			case "server":
				if inst.Server != inst.original.Server {
					return true
				}
			case "secret":
				if inst.Secret != inst.original.Secret {
					return true
				}
			case "meta_json":
				if inst.MetaJson != inst.original.MetaJson {
					return true
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					return true
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					return true
				}
			case "deleted_at":
				if inst.DeletedAt != inst.original.DeletedAt {
					return true
				}
			default:
			}
		}
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *ChannelsN) StaledKV(onlyFields ...string) query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &channelsOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			kv["id"] = inst.Id
		}
		if inst.Name != inst.original.Name {
			kv["name"] = inst.Name
		}
		if inst.Type != inst.original.Type {
			kv["type"] = inst.Type
		}
		if inst.Server != inst.original.Server {
			kv["server"] = inst.Server
		}
		if inst.Secret != inst.original.Secret {
			kv["secret"] = inst.Secret
		}
		if inst.MetaJson != inst.original.MetaJson {
			kv["meta_json"] = inst.MetaJson
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			kv["created_at"] = inst.CreatedAt
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			kv["updated_at"] = inst.UpdatedAt
		}
		if inst.DeletedAt != inst.original.DeletedAt {
			kv["deleted_at"] = inst.DeletedAt
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					kv["id"] = inst.Id
				}
			case "name":
				if inst.Name != inst.original.Name {
					kv["name"] = inst.Name
				}
			case "type":
				if inst.Type != inst.original.Type {
					kv["type"] = inst.Type
				}
			case "server":
				if inst.Server != inst.original.Server {
					kv["server"] = inst.Server
				}
			case "secret":
				if inst.Secret != inst.original.Secret {
					kv["secret"] = inst.Secret
				}
			case "meta_json":
				if inst.MetaJson != inst.original.MetaJson {
					kv["meta_json"] = inst.MetaJson
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					kv["created_at"] = inst.CreatedAt
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					kv["updated_at"] = inst.UpdatedAt
				}
			case "deleted_at":
				if inst.DeletedAt != inst.original.DeletedAt {
					kv["deleted_at"] = inst.DeletedAt
				}
			default:
			}
		}
	}

	return kv
}

// Save create a new model or update it
func (inst *ChannelsN) Save(ctx context.Context, onlyFields ...string) error {
	if inst.channelsModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.channelsModel.SaveOrUpdate(ctx, *inst, onlyFields...)
	if err != nil {
		return err
	}

	inst.Id = null.IntFrom(id)
	return nil
}

// Delete remove a channels
func (inst *ChannelsN) Delete(ctx context.Context) error {
	if inst.channelsModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.channelsModel.DeleteById(ctx, inst.Id.Int64)
	if err != nil {
		return err
	}

	return nil
}

// String convert instance to json string
func (inst *ChannelsN) String() string {
	rs, _ := json.Marshal(inst)
	return string(rs)
}

type channelsScope struct {
	name  string
	apply func(builder query.Condition)
}

var channelsGlobalScopes = make([]channelsScope, 0)
var channelsLocalScopes = make([]channelsScope, 0)

// AddGlobalScopeForChannels assign a global scope to a model
func AddGlobalScopeForChannels(name string, apply func(builder query.Condition)) {
	channelsGlobalScopes = append(channelsGlobalScopes, channelsScope{name: name, apply: apply})
}

// AddLocalScopeForChannels assign a local scope to a model
func AddLocalScopeForChannels(name string, apply func(builder query.Condition)) {
	channelsLocalScopes = append(channelsLocalScopes, channelsScope{name: name, apply: apply})
}

func (m *ChannelsModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range channelsGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range channelsLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *ChannelsModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *ChannelsModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}

	return true
}

type Channels struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Server    string `json:"server,omitempty"`
	Secret    string `json:"secret,omitempty"`
	MetaJson  string `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (w Channels) ToChannelsN(allows ...string) ChannelsN {
	if len(allows) == 0 {
		return ChannelsN{

			Id:        null.IntFrom(int64(w.Id)),
			Name:      null.StringFrom(w.Name),
			Type:      null.StringFrom(w.Type),
			Server:    null.StringFrom(w.Server),
			Secret:    null.StringFrom(w.Secret),
			MetaJson:  null.StringFrom(w.MetaJson),
			CreatedAt: null.TimeFrom(w.CreatedAt),
			UpdatedAt: null.TimeFrom(w.UpdatedAt),
			DeletedAt: null.TimeFrom(w.DeletedAt),
		}
	}

	res := ChannelsN{}
	for _, al := range allows {
		switch strcase.ToSnake(al) {

		case "id":
			res.Id = null.IntFrom(int64(w.Id))
		case "name":
			res.Name = null.StringFrom(w.Name)
		case "type":
			res.Type = null.StringFrom(w.Type)
		case "server":
			res.Server = null.StringFrom(w.Server)
		case "secret":
			res.Secret = null.StringFrom(w.Secret)
		case "meta_json":
			res.MetaJson = null.StringFrom(w.MetaJson)
		case "created_at":
			res.CreatedAt = null.TimeFrom(w.CreatedAt)
		case "updated_at":
			res.UpdatedAt = null.TimeFrom(w.UpdatedAt)
		case "deleted_at":
			res.DeletedAt = null.TimeFrom(w.DeletedAt)
		default:
		}
	}

	return res
}

// As convert object to other type
// dst must be a pointer to struct
func (w Channels) As(dst interface{}) error {
	return query.Copy(w, dst)
}

func (w *ChannelsN) ToChannels() Channels {
	return Channels{

		Id:        w.Id.Int64,
		Name:      w.Name.String,
		Type:      w.Type.String,
		Server:    w.Server.String,
		Secret:    w.Secret.String,
		MetaJson:  w.MetaJson.String,
		CreatedAt: w.CreatedAt.Time,
		UpdatedAt: w.UpdatedAt.Time,
		DeletedAt: w.DeletedAt.Time,
	}
}

// ChannelsModel is a model which encapsulates the operations of the object
type ChannelsModel struct {
	db        *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes  []string

	query query.SQLBuilder
}

var channelsTableName = "channels"

// ChannelsTable return table name for Channels
func ChannelsTable() string {
	return channelsTableName
}

const (
	FieldChannelsId        = "id"
	FieldChannelsName      = "name"
	FieldChannelsType      = "type"
	FieldChannelsServer    = "server"
	FieldChannelsSecret    = "secret"
	FieldChannelsMetaJson  = "meta_json"
	FieldChannelsCreatedAt = "created_at"
	FieldChannelsUpdatedAt = "updated_at"
	FieldChannelsDeletedAt = "deleted_at"
)

// ChannelsFields return all fields in Channels model
func ChannelsFields() []string {
	return []string{
		"id",
		"name",
		"type",
		"server",
		"secret",
		"meta_json",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

func SetChannelsTable(tableName string) {
	channelsTableName = tableName
}

// NewChannelsModel create a ChannelsModel
func NewChannelsModel(db query.Database) *ChannelsModel {
	return &ChannelsModel{
		db:                  query.NewDatabaseWrap(db),
		tableName:           channelsTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes:  make([]string, 0),
		query:               query.Builder(),
	}
}

// GetDB return database instance
func (m *ChannelsModel) GetDB() query.Database {
	return m.db.GetDB()
}

// WithTrashed force soft deleted models to appear in a result set
func (m *ChannelsModel) WithTrashed() *ChannelsModel {
	return m.WithoutGlobalScopes("soft_delete")
}

func (m *ChannelsModel) clone() *ChannelsModel {
	return &ChannelsModel{
		db:                  m.db,
		tableName:           m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes:  append([]string{}, m.includeLocalScopes...),
		query:               m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *ChannelsModel) WithoutGlobalScopes(names ...string) *ChannelsModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *ChannelsModel) WithLocalScopes(names ...string) *ChannelsModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Condition add query builder to model
func (m *ChannelsModel) Condition(builder query.SQLBuilder) *ChannelsModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *ChannelsModel) Find(ctx context.Context, id int64) (*ChannelsN, error) {
	return m.First(ctx, m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *ChannelsModel) Exists(ctx context.Context, builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(ctx, builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *ChannelsModel) Count(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {
	sqlStr, params := m.query.
		Merge(builders...).
		Table(m.tableName).
		AppendCondition(m.applyScope()).
		ResolveCount()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	rows.Next()
	var res int64
	if err := rows.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (m *ChannelsModel) Paginate(ctx context.Context, page int64, perPage int64, builders ...query.SQLBuilder) ([]ChannelsN, query.PaginateMeta, error) {
	if page <= 0 {
		page = 1
	}

	if perPage <= 0 {
		perPage = 15
	}

	meta := query.PaginateMeta{
		PerPage: perPage,
		Page:    page,
	}

	count, err := m.Count(ctx, builders...)
	if err != nil {
		return nil, meta, err
	}

	meta.Total = count
	meta.LastPage = count / perPage
	if count%perPage != 0 {
		meta.LastPage += 1
	}

	res, err := m.Get(ctx, append([]query.SQLBuilder{query.Builder().Limit(perPage).Offset((page - 1) * perPage)}, builders...)...)
	if err != nil {
		return res, meta, err
	}

	return res, meta, nil
}

// Get retrieve all results for given query
func (m *ChannelsModel) Get(ctx context.Context, builders ...query.SQLBuilder) ([]ChannelsN, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"id",
			"name",
			"type",
			"server",
			"secret",
			"meta_json",
			"created_at",
			"updated_at",
			"deleted_at",
		)
	}

	fields := b.GetFields()
	selectFields := make([]query.Expr, 0)

	for _, f := range fields {
		switch strcase.ToSnake(f.Value) {

		case "id":
			selectFields = append(selectFields, f)
		case "name":
			selectFields = append(selectFields, f)
		case "type":
			selectFields = append(selectFields, f)
		case "server":
			selectFields = append(selectFields, f)
		case "secret":
			selectFields = append(selectFields, f)
		case "meta_json":
			selectFields = append(selectFields, f)
		case "created_at":
			selectFields = append(selectFields, f)
		case "updated_at":
			selectFields = append(selectFields, f)
		case "deleted_at":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*ChannelsN, []interface{}) {
		var channelsVar ChannelsN
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {

			case "id":
				scanFields = append(scanFields, &channelsVar.Id)
			case "name":
				scanFields = append(scanFields, &channelsVar.Name)
			case "type":
				scanFields = append(scanFields, &channelsVar.Type)
			case "server":
				scanFields = append(scanFields, &channelsVar.Server)
			case "secret":
				scanFields = append(scanFields, &channelsVar.Secret)
			case "meta_json":
				scanFields = append(scanFields, &channelsVar.MetaJson)
			case "created_at":
				scanFields = append(scanFields, &channelsVar.CreatedAt)
			case "updated_at":
				scanFields = append(scanFields, &channelsVar.UpdatedAt)
			case "deleted_at":
				scanFields = append(scanFields, &channelsVar.DeletedAt)
			}
		}

		return &channelsVar, scanFields
	}

	sqlStr, params := b.Fields(selectFields...).ResolveQuery()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	channelss := make([]ChannelsN, 0)
	for rows.Next() {
		channelsReal, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		channelsReal.original = &channelsOriginal{}
		_ = query.Copy(channelsReal, channelsReal.original)

		channelsReal.SetModel(m)
		channelss = append(channelss, *channelsReal)
	}

	return channelss, nil
}

// First return first result for given query
func (m *ChannelsModel) First(ctx context.Context, builders ...query.SQLBuilder) (*ChannelsN, error) {
	res, err := m.Get(ctx, append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, query.ErrNoResult
	}

	return &res[0], nil
}

// Create save a new channels to database
func (m *ChannelsModel) Create(ctx context.Context, kv query.KV) (int64, error) {

	if _, ok := kv["created_at"]; !ok {
		kv["created_at"] = time.Now()
	}

	if _, ok := kv["updated_at"]; !ok {
		kv["updated_at"] = time.Now()
	}

	sqlStr, params := m.query.Table(m.tableName).ResolveInsert(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// SaveAll save all channelss to database
func (m *ChannelsModel) SaveAll(ctx context.Context, channelss []ChannelsN) ([]int64, error) {
	ids := make([]int64, 0)
	for _, channels := range channelss {
		id, err := m.Save(ctx, channels)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a channels to database
func (m *ChannelsModel) Save(ctx context.Context, channels ChannelsN, onlyFields ...string) (int64, error) {
	return m.Create(ctx, channels.StaledKV(onlyFields...))
}

// SaveOrUpdate save a new channels or update it when it has a id > 0
func (m *ChannelsModel) SaveOrUpdate(ctx context.Context, channels ChannelsN, onlyFields ...string) (id int64, updated bool, err error) {
	if channels.Id.Int64 > 0 {
		_, _err := m.UpdateById(ctx, channels.Id.Int64, channels, onlyFields...)
		return channels.Id.Int64, true, _err
	}

	_id, _err := m.Save(ctx, channels, onlyFields...)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *ChannelsModel) UpdateFields(ctx context.Context, kv query.KV, builders ...query.SQLBuilder) (int64, error) {
	if len(kv) == 0 {
		return 0, nil
	}

	kv["updated_at"] = time.Now()

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).
		Table(m.tableName).
		ResolveUpdate(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Update update a model for given query
func (m *ChannelsModel) Update(ctx context.Context, builder query.SQLBuilder, channels ChannelsN, onlyFields ...string) (int64, error) {
	return m.UpdateFields(ctx, channels.StaledKV(onlyFields...), builder)
}

// UpdateById update a model by id
func (m *ChannelsModel) UpdateById(ctx context.Context, id int64, channels ChannelsN, onlyFields ...string) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).UpdateFields(ctx, channels.StaledKV(onlyFields...))
}

// ForceDelete permanently remove a soft deleted model from the database
func (m *ChannelsModel) ForceDelete(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {
	m2 := m.WithTrashed()

	sqlStr, params := m2.query.Merge(builders...).AppendCondition(m2.applyScope()).Table(m2.tableName).ResolveDelete()

	res, err := m2.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// ForceDeleteById permanently remove a soft deleted model from the database by id
func (m *ChannelsModel) ForceDeleteById(ctx context.Context, id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).ForceDelete(ctx)
}

// Restore restore a soft deleted model into an active state
func (m *ChannelsModel) Restore(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {
	m2 := m.WithTrashed()
	return m2.UpdateFields(ctx, query.KV{
		"deleted_at": nil,
	}, builders...)
}

// RestoreById restore a soft deleted model into an active state by id
func (m *ChannelsModel) RestoreById(ctx context.Context, id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Restore(ctx)
}

// Delete remove a model
func (m *ChannelsModel) Delete(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {

	return m.UpdateFields(ctx, query.KV{
		"deleted_at": time.Now(),
	}, builders...)

}

// DeleteById remove a model by id
func (m *ChannelsModel) DeleteById(ctx context.Context, id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Delete(ctx)
}
