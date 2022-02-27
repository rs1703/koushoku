// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Tag is an object representing the database table.
type Tag struct {
	ID   int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Slug string `boil:"slug" json:"slug" toml:"slug" yaml:"slug"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *tagR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tagL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TagColumns = struct {
	ID   string
	Slug string
	Name string
}{
	ID:   "id",
	Slug: "slug",
	Name: "name",
}

var TagTableColumns = struct {
	ID   string
	Slug string
	Name string
}{
	ID:   "tag.id",
	Slug: "tag.slug",
	Name: "tag.name",
}

// Generated where

var TagWhere = struct {
	ID   whereHelperint64
	Slug whereHelperstring
	Name whereHelperstring
}{
	ID:   whereHelperint64{field: "\"tag\".\"id\""},
	Slug: whereHelperstring{field: "\"tag\".\"slug\""},
	Name: whereHelperstring{field: "\"tag\".\"name\""},
}

// TagRels is where relationship names are stored.
var TagRels = struct {
	Archives string
}{
	Archives: "Archives",
}

// tagR is where relationships are stored.
type tagR struct {
	Archives ArchiveSlice `boil:"Archives" json:"Archives" toml:"Archives" yaml:"Archives"`
}

// NewStruct creates a new relationship struct
func (*tagR) NewStruct() *tagR {
	return &tagR{}
}

// tagL is where Load methods for each relationship are stored.
type tagL struct{}

var (
	tagAllColumns            = []string{"id", "slug", "name"}
	tagColumnsWithoutDefault = []string{}
	tagColumnsWithDefault    = []string{"id", "slug", "name"}
	tagPrimaryKeyColumns     = []string{"id"}
	tagGeneratedColumns      = []string{}
)

type (
	// TagSlice is an alias for a slice of pointers to Tag.
	// This should almost always be used instead of []Tag.
	TagSlice []*Tag
	// TagHook is the signature for custom Tag hook methods
	TagHook func(boil.Executor, *Tag) error

	tagQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tagType                 = reflect.TypeOf(&Tag{})
	tagMapping              = queries.MakeStructMapping(tagType)
	tagPrimaryKeyMapping, _ = queries.BindMapping(tagType, tagMapping, tagPrimaryKeyColumns)
	tagInsertCacheMut       sync.RWMutex
	tagInsertCache          = make(map[string]insertCache)
	tagUpdateCacheMut       sync.RWMutex
	tagUpdateCache          = make(map[string]updateCache)
	tagUpsertCacheMut       sync.RWMutex
	tagUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tagAfterSelectHooks []TagHook

var tagBeforeInsertHooks []TagHook
var tagAfterInsertHooks []TagHook

var tagBeforeUpdateHooks []TagHook
var tagAfterUpdateHooks []TagHook

var tagBeforeDeleteHooks []TagHook
var tagAfterDeleteHooks []TagHook

var tagBeforeUpsertHooks []TagHook
var tagAfterUpsertHooks []TagHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Tag) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range tagAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Tag) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tagBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Tag) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tagAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Tag) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range tagBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Tag) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range tagAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Tag) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range tagBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Tag) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range tagAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Tag) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tagBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Tag) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tagAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTagHook registers your hook function for all future operations.
func AddTagHook(hookPoint boil.HookPoint, tagHook TagHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		tagAfterSelectHooks = append(tagAfterSelectHooks, tagHook)
	case boil.BeforeInsertHook:
		tagBeforeInsertHooks = append(tagBeforeInsertHooks, tagHook)
	case boil.AfterInsertHook:
		tagAfterInsertHooks = append(tagAfterInsertHooks, tagHook)
	case boil.BeforeUpdateHook:
		tagBeforeUpdateHooks = append(tagBeforeUpdateHooks, tagHook)
	case boil.AfterUpdateHook:
		tagAfterUpdateHooks = append(tagAfterUpdateHooks, tagHook)
	case boil.BeforeDeleteHook:
		tagBeforeDeleteHooks = append(tagBeforeDeleteHooks, tagHook)
	case boil.AfterDeleteHook:
		tagAfterDeleteHooks = append(tagAfterDeleteHooks, tagHook)
	case boil.BeforeUpsertHook:
		tagBeforeUpsertHooks = append(tagBeforeUpsertHooks, tagHook)
	case boil.AfterUpsertHook:
		tagAfterUpsertHooks = append(tagAfterUpsertHooks, tagHook)
	}
}

// OneG returns a single tag record from the query using the global executor.
func (q tagQuery) OneG() (*Tag, error) {
	return q.One(boil.GetDB())
}

// One returns a single tag record from the query.
func (q tagQuery) One(exec boil.Executor) (*Tag, error) {
	o := &Tag{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for tag")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Tag records from the query using the global executor.
func (q tagQuery) AllG() (TagSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Tag records from the query.
func (q tagQuery) All(exec boil.Executor) (TagSlice, error) {
	var o []*Tag

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Tag slice")
	}

	if len(tagAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Tag records in the query, and panics on error.
func (q tagQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Tag records in the query.
func (q tagQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count tag rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q tagQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q tagQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if tag exists")
	}

	return count > 0, nil
}

// Archives retrieves all the archive's Archives with an executor.
func (o *Tag) Archives(mods ...qm.QueryMod) archiveQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"archive_tags\" on \"archive\".\"id\" = \"archive_tags\".\"archive_id\""),
		qm.Where("\"archive_tags\".\"tag_id\"=?", o.ID),
	)

	query := Archives(queryMods...)
	queries.SetFrom(query.Query, "\"archive\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"archive\".*"})
	}

	return query
}

// LoadArchives allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (tagL) LoadArchives(e boil.Executor, singular bool, maybeTag interface{}, mods queries.Applicator) error {
	var slice []*Tag
	var object *Tag

	if singular {
		object = maybeTag.(*Tag)
	} else {
		slice = *maybeTag.(*[]*Tag)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tagR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tagR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("\"archive\".id, \"archive\".path, \"archive\".created_at, \"archive\".updated_at, \"archive\".published_at, \"archive\".title, \"archive\".slug, \"archive\".pages, \"archive\".size, \"a\".\"tag_id\""),
		qm.From("\"archive\""),
		qm.InnerJoin("\"archive_tags\" as \"a\" on \"archive\".\"id\" = \"a\".\"archive_id\""),
		qm.WhereIn("\"a\".\"tag_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load archive")
	}

	var resultSlice []*Archive

	var localJoinCols []int64
	for results.Next() {
		one := new(Archive)
		var localJoinCol int64

		err = results.Scan(&one.ID, &one.Path, &one.CreatedAt, &one.UpdatedAt, &one.PublishedAt, &one.Title, &one.Slug, &one.Pages, &one.Size, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for archive")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice archive")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on archive")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for archive")
	}

	if len(archiveAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Archives = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &archiveR{}
			}
			foreign.R.Tags = append(foreign.R.Tags, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Archives = append(local.R.Archives, foreign)
				if foreign.R == nil {
					foreign.R = &archiveR{}
				}
				foreign.R.Tags = append(foreign.R.Tags, local)
				break
			}
		}
	}

	return nil
}

// AddArchivesG adds the given related objects to the existing relationships
// of the tag, optionally inserting them as new records.
// Appends related to o.R.Archives.
// Sets related.R.Tags appropriately.
// Uses the global database handle.
func (o *Tag) AddArchivesG(insert bool, related ...*Archive) error {
	return o.AddArchives(boil.GetDB(), insert, related...)
}

// AddArchives adds the given related objects to the existing relationships
// of the tag, optionally inserting them as new records.
// Appends related to o.R.Archives.
// Sets related.R.Tags appropriately.
func (o *Tag) AddArchives(exec boil.Executor, insert bool, related ...*Archive) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"archive_tags\" (\"tag_id\", \"archive_id\") values ($1, $2)"
		values := []interface{}{o.ID, rel.ID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, query)
			fmt.Fprintln(boil.DebugWriter, values)
		}
		_, err = exec.Exec(query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &tagR{
			Archives: related,
		}
	} else {
		o.R.Archives = append(o.R.Archives, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &archiveR{
				Tags: TagSlice{o},
			}
		} else {
			rel.R.Tags = append(rel.R.Tags, o)
		}
	}
	return nil
}

// SetArchivesG removes all previously related items of the
// tag replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Tags's Archives accordingly.
// Replaces o.R.Archives with related.
// Sets related.R.Tags's Archives accordingly.
// Uses the global database handle.
func (o *Tag) SetArchivesG(insert bool, related ...*Archive) error {
	return o.SetArchives(boil.GetDB(), insert, related...)
}

// SetArchives removes all previously related items of the
// tag replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Tags's Archives accordingly.
// Replaces o.R.Archives with related.
// Sets related.R.Tags's Archives accordingly.
func (o *Tag) SetArchives(exec boil.Executor, insert bool, related ...*Archive) error {
	query := "delete from \"archive_tags\" where \"tag_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeArchivesFromTagsSlice(o, related)
	if o.R != nil {
		o.R.Archives = nil
	}
	return o.AddArchives(exec, insert, related...)
}

// RemoveArchivesG relationships from objects passed in.
// Removes related items from R.Archives (uses pointer comparison, removal does not keep order)
// Sets related.R.Tags.
// Uses the global database handle.
func (o *Tag) RemoveArchivesG(related ...*Archive) error {
	return o.RemoveArchives(boil.GetDB(), related...)
}

// RemoveArchives relationships from objects passed in.
// Removes related items from R.Archives (uses pointer comparison, removal does not keep order)
// Sets related.R.Tags.
func (o *Tag) RemoveArchives(exec boil.Executor, related ...*Archive) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	query := fmt.Sprintf(
		"delete from \"archive_tags\" where \"tag_id\" = $1 and \"archive_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err = exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeArchivesFromTagsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Archives {
			if rel != ri {
				continue
			}

			ln := len(o.R.Archives)
			if ln > 1 && i < ln-1 {
				o.R.Archives[i] = o.R.Archives[ln-1]
			}
			o.R.Archives = o.R.Archives[:ln-1]
			break
		}
	}

	return nil
}

func removeArchivesFromTagsSlice(o *Tag, related []*Archive) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Tags {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Tags)
			if ln > 1 && i < ln-1 {
				rel.R.Tags[i] = rel.R.Tags[ln-1]
			}
			rel.R.Tags = rel.R.Tags[:ln-1]
			break
		}
	}
}

// Tags retrieves all the records using an executor.
func Tags(mods ...qm.QueryMod) tagQuery {
	mods = append(mods, qm.From("\"tag\""))
	return tagQuery{NewQuery(mods...)}
}

// FindTagG retrieves a single record by ID.
func FindTagG(iD int64, selectCols ...string) (*Tag, error) {
	return FindTag(boil.GetDB(), iD, selectCols...)
}

// FindTag retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTag(exec boil.Executor, iD int64, selectCols ...string) (*Tag, error) {
	tagObj := &Tag{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"tag\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, tagObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from tag")
	}

	if err = tagObj.doAfterSelectHooks(exec); err != nil {
		return tagObj, err
	}

	return tagObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Tag) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Tag) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tag provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tagColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tagInsertCacheMut.RLock()
	cache, cached := tagInsertCache[key]
	tagInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tagAllColumns,
			tagColumnsWithDefault,
			tagColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tagType, tagMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"tag\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"tag\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into tag")
	}

	if !cached {
		tagInsertCacheMut.Lock()
		tagInsertCache[key] = cache
		tagInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Tag record using the global executor.
// See Update for more documentation.
func (o *Tag) UpdateG(columns boil.Columns) error {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Tag.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Tag) Update(exec boil.Executor, columns boil.Columns) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(columns, nil)
	tagUpdateCacheMut.RLock()
	cache, cached := tagUpdateCache[key]
	tagUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tagAllColumns,
			tagPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update tag, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"tag\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, tagPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, append(wl, tagPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update tag row")
	}

	if !cached {
		tagUpdateCacheMut.Lock()
		tagUpdateCache[key] = cache
		tagUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q tagQuery) UpdateAllG(cols M) error {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q tagQuery) UpdateAll(exec boil.Executor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for tag")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TagSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TagSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"tag\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, tagPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in tag slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Tag) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Tag) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tag provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tagColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	tagUpsertCacheMut.RLock()
	cache, cached := tagUpsertCache[key]
	tagUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tagAllColumns,
			tagColumnsWithDefault,
			tagColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			tagAllColumns,
			tagPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert tag, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(tagPrimaryKeyColumns))
			copy(conflict, tagPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"tag\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tagType, tagMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert tag")
	}

	if !cached {
		tagUpsertCacheMut.Lock()
		tagUpsertCache[key] = cache
		tagUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single Tag record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Tag) DeleteG() error {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Tag record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Tag) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Tag provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tagPrimaryKeyMapping)
	sql := "DELETE FROM \"tag\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from tag")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

func (q tagQuery) DeleteAllG() error {
	return q.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all matching rows.
func (q tagQuery) DeleteAll(exec boil.Executor) error {
	if q.Query == nil {
		return errors.New("models: no tagQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from tag")
	}

	return nil
}

// DeleteAllG deletes all rows in the slice.
func (o TagSlice) DeleteAllG() error {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TagSlice) DeleteAll(exec boil.Executor) error {
	if len(o) == 0 {
		return nil
	}

	if len(tagBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"tag\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tagPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from tag slice")
	}

	if len(tagAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Tag) ReloadG() error {
	if o == nil {
		return errors.New("models: no Tag provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Tag) Reload(exec boil.Executor) error {
	ret, err := FindTag(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TagSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty TagSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TagSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TagSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"tag\".* FROM \"tag\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tagPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TagSlice")
	}

	*o = slice

	return nil
}

// TagExistsG checks if the Tag row exists.
func TagExistsG(iD int64) (bool, error) {
	return TagExists(boil.GetDB(), iD)
}

// TagExists checks if the Tag row exists.
func TagExists(exec boil.Executor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"tag\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if tag exists")
	}

	return exists, nil
}
