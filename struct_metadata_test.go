package structmeta

import (
	"encoding/json"
	"testing"
	"time"
)

var testTagName = "column"

type TestAsset struct {
	//StructMetadata

	Id          int        `column:"Asset.asset_id"`
	Status      string     `column:"mdAssetStatus.assetStatus_desc,Asset.assetStatus_id=mdAssetStatus.assetStatus_id,enum,ro"`
	CreatedDate *time.Time `column:"Asset.created_date"`
	UpdatedDate *time.Time `column:"Asset.last_modified_date"`
	CreatedBy   *int       `column:"Asset.created_by"`
	UpdatedBy   *int       `column:"Asset.last_modified_by"`
	Name        *string    `column:"Asset.file_name"`
	FilmId      *int       `column:"Asset.Asset_Film_Detail_Id"`
	ParentId    *int       `column:"Asset_ParentAsset.parent_asset_id,Asset.asset_id=Asset_ParentAsset.asset_id"`
	Version     *int       `column:"Asset.version_number"`
	OwnerId     *int       `column:"AssetFilmDetail.Content_Owner_Id,Asset.Asset_Film_Detail_Id=AssetFilmDetail.Asset_Film_Detail_Id"`
	Size        *int       `column:"Asset.file_size_int"`
	Checksum    *string    `column:"Asset.Checksum"`
	Notes       *string    `column:"AssetNotes.notes,AssetNotes.asset_id=Asset.asset_id"`

	// Read only field
	BugCount int `column:"(select count(1) from AssetBug where AssetBug.asset_id=Asset.asset_id),ro"`
}

func Test_StructMetadata(t *testing.T) {
	f := TestAsset{
		Status: "Received",
		Name:   new(string),
		Notes:  new(string),
	}
	*f.Name = "foo"
	*f.Notes = "test note"

	m := ParseStructMetadata(&f, testTagName, false)
	t.Logf("%s Fields: %d", len(m))

	if len(m.HasArg("ro")) != 2 {
		t.Fatal("Should have 2 ro fields")
	}

	if len(m.HasArg("enum")) != 1 {
		t.Fatal("Should have 2 ro fields")
	}

	if m.FieldByKey("Asset.Checksum").Field != "Checksum" {
		t.Fatal("Field -> key mismatch")
	}

	if len(m.NotHasArg("ro")) != len(m)-2 {
		t.Fatal("arg length mismatch")
	}

	if len(m.FieldNames()) != len(m) || len(m.Keys()) != len(m) {
		t.Fatal("field/key length mismatch")
	}

	if len(m.Values()) != len(m) {
		t.Fatal("length mismatch")
	}

	if m.FieldByName("Name") == nil {
		t.Fatal("field not found")
	}

	b, _ := json.MarshalIndent(m, "", "  ")
	t.Logf("%s\n", b)
}

func Test_StructMetadata_Slice(t *testing.T) {
	f := TestAsset{
		Status: "Received",
		Name:   new(string),
		Notes:  new(string),
	}
	*f.Name = "foo"
	*f.Notes = "test note"

	fs := []TestAsset{f}
	m := ParseStructMetadata(&fs, testTagName, false)
	fld := m.FieldByName("Id")
	t.Log(fld.Key)
}
