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

}

// CreativeIslandN is a CreativeIsland object, all fields are nullable
type CreativeIslandN struct {
	original            *creativeIslandOriginal
	creativeIslandModel *CreativeIslandModel

	Id                     null.Int    `json:"id"`
	IslandId               null.String `json:"island_id"`
	Title                  null.String `json:"title"`
	TitleColor             null.String `json:"title_color"`
	Description            null.String `json:"description"`
	Category               null.String `json:"category"`
	ModelType              null.String `json:"model_type"`
	WordCount              null.Int    `json:"word_count"`
	Hint                   null.String `json:"hint"`
	Vendor                 null.String `json:"vendor"`
	Model                  null.String `json:"model"`
	StylePreset            null.String `json:"style_preset,omitempty"`
	Prompt                 null.String `json:"prompt"`
	BgImage                null.String `json:"bg_image,omitempty"`
	BgEmbeddedImage        null.String `json:"bg_embedded_image,omitempty"`
	Label                  null.String `json:"label,omitempty"`
	LabelColor             null.String `json:"label_color,omitempty"`
	SubmitBtnText          null.String `json:"submit_btn_text,omitempty"`
	PromptInputTitle       null.String `json:"prompt_input_title,omitempty"`
	WaitSeconds            null.Int    `json:"wait_seconds,omitempty"`
	ShowImageStyleSelector null.Int    `json:"show_image_style_selector,omitempty"`
	NoPrompt               null.Int    `json:"no_prompt,omitempty"`
	VersionMin             null.String `json:"version_min,omitempty"`
	VersionMax             null.String `json:"version_max,omitempty"`
	Status                 null.Int    `json:"status"`
	Priority               null.Int    `json:"priority,omitempty"`
	Ext                    null.String `json:"ext,omitempty"`
	CreatedAt              null.Time
	UpdatedAt              null.Time
}

// As convert object to other type
// dst must be a pointer to struct
func (inst *CreativeIslandN) As(dst interface{}) error {
	return query.Copy(inst, dst)
}

// SetModel set model for CreativeIsland
func (inst *CreativeIslandN) SetModel(creativeIslandModel *CreativeIslandModel) {
	inst.creativeIslandModel = creativeIslandModel
}

// creativeIslandOriginal is an object which stores original CreativeIsland from database
type creativeIslandOriginal struct {
	Id                     null.Int
	IslandId               null.String
	Title                  null.String
	TitleColor             null.String
	Description            null.String
	Category               null.String
	ModelType              null.String
	WordCount              null.Int
	Hint                   null.String
	Vendor                 null.String
	Model                  null.String
	StylePreset            null.String
	Prompt                 null.String
	BgImage                null.String
	BgEmbeddedImage        null.String
	Label                  null.String
	LabelColor             null.String
	SubmitBtnText          null.String
	PromptInputTitle       null.String
	WaitSeconds            null.Int
	ShowImageStyleSelector null.Int
	NoPrompt               null.Int
	VersionMin             null.String
	VersionMax             null.String
	Status                 null.Int
	Priority               null.Int
	Ext                    null.String
	CreatedAt              null.Time
	UpdatedAt              null.Time
}

// Staled identify whether the object has been modified
func (inst *CreativeIslandN) Staled(onlyFields ...string) bool {
	if inst.original == nil {
		inst.original = &creativeIslandOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			return true
		}
		if inst.IslandId != inst.original.IslandId {
			return true
		}
		if inst.Title != inst.original.Title {
			return true
		}
		if inst.TitleColor != inst.original.TitleColor {
			return true
		}
		if inst.Description != inst.original.Description {
			return true
		}
		if inst.Category != inst.original.Category {
			return true
		}
		if inst.ModelType != inst.original.ModelType {
			return true
		}
		if inst.WordCount != inst.original.WordCount {
			return true
		}
		if inst.Hint != inst.original.Hint {
			return true
		}
		if inst.Vendor != inst.original.Vendor {
			return true
		}
		if inst.Model != inst.original.Model {
			return true
		}
		if inst.StylePreset != inst.original.StylePreset {
			return true
		}
		if inst.Prompt != inst.original.Prompt {
			return true
		}
		if inst.BgImage != inst.original.BgImage {
			return true
		}
		if inst.BgEmbeddedImage != inst.original.BgEmbeddedImage {
			return true
		}
		if inst.Label != inst.original.Label {
			return true
		}
		if inst.LabelColor != inst.original.LabelColor {
			return true
		}
		if inst.SubmitBtnText != inst.original.SubmitBtnText {
			return true
		}
		if inst.PromptInputTitle != inst.original.PromptInputTitle {
			return true
		}
		if inst.WaitSeconds != inst.original.WaitSeconds {
			return true
		}
		if inst.ShowImageStyleSelector != inst.original.ShowImageStyleSelector {
			return true
		}
		if inst.NoPrompt != inst.original.NoPrompt {
			return true
		}
		if inst.VersionMin != inst.original.VersionMin {
			return true
		}
		if inst.VersionMax != inst.original.VersionMax {
			return true
		}
		if inst.Status != inst.original.Status {
			return true
		}
		if inst.Priority != inst.original.Priority {
			return true
		}
		if inst.Ext != inst.original.Ext {
			return true
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			return true
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			return true
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					return true
				}
			case "island_id":
				if inst.IslandId != inst.original.IslandId {
					return true
				}
			case "title":
				if inst.Title != inst.original.Title {
					return true
				}
			case "title_color":
				if inst.TitleColor != inst.original.TitleColor {
					return true
				}
			case "description":
				if inst.Description != inst.original.Description {
					return true
				}
			case "category":
				if inst.Category != inst.original.Category {
					return true
				}
			case "model_type":
				if inst.ModelType != inst.original.ModelType {
					return true
				}
			case "word_count":
				if inst.WordCount != inst.original.WordCount {
					return true
				}
			case "hint":
				if inst.Hint != inst.original.Hint {
					return true
				}
			case "vendor":
				if inst.Vendor != inst.original.Vendor {
					return true
				}
			case "model":
				if inst.Model != inst.original.Model {
					return true
				}
			case "style_preset":
				if inst.StylePreset != inst.original.StylePreset {
					return true
				}
			case "prompt":
				if inst.Prompt != inst.original.Prompt {
					return true
				}
			case "bg_image":
				if inst.BgImage != inst.original.BgImage {
					return true
				}
			case "bg_embedded_image":
				if inst.BgEmbeddedImage != inst.original.BgEmbeddedImage {
					return true
				}
			case "label":
				if inst.Label != inst.original.Label {
					return true
				}
			case "label_color":
				if inst.LabelColor != inst.original.LabelColor {
					return true
				}
			case "submit_btn_text":
				if inst.SubmitBtnText != inst.original.SubmitBtnText {
					return true
				}
			case "prompt_input_title":
				if inst.PromptInputTitle != inst.original.PromptInputTitle {
					return true
				}
			case "wait_seconds":
				if inst.WaitSeconds != inst.original.WaitSeconds {
					return true
				}
			case "show_image_style_selector":
				if inst.ShowImageStyleSelector != inst.original.ShowImageStyleSelector {
					return true
				}
			case "no_prompt":
				if inst.NoPrompt != inst.original.NoPrompt {
					return true
				}
			case "version_min":
				if inst.VersionMin != inst.original.VersionMin {
					return true
				}
			case "version_max":
				if inst.VersionMax != inst.original.VersionMax {
					return true
				}
			case "status":
				if inst.Status != inst.original.Status {
					return true
				}
			case "priority":
				if inst.Priority != inst.original.Priority {
					return true
				}
			case "ext":
				if inst.Ext != inst.original.Ext {
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
			default:
			}
		}
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *CreativeIslandN) StaledKV(onlyFields ...string) query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &creativeIslandOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			kv["id"] = inst.Id
		}
		if inst.IslandId != inst.original.IslandId {
			kv["island_id"] = inst.IslandId
		}
		if inst.Title != inst.original.Title {
			kv["title"] = inst.Title
		}
		if inst.TitleColor != inst.original.TitleColor {
			kv["title_color"] = inst.TitleColor
		}
		if inst.Description != inst.original.Description {
			kv["description"] = inst.Description
		}
		if inst.Category != inst.original.Category {
			kv["category"] = inst.Category
		}
		if inst.ModelType != inst.original.ModelType {
			kv["model_type"] = inst.ModelType
		}
		if inst.WordCount != inst.original.WordCount {
			kv["word_count"] = inst.WordCount
		}
		if inst.Hint != inst.original.Hint {
			kv["hint"] = inst.Hint
		}
		if inst.Vendor != inst.original.Vendor {
			kv["vendor"] = inst.Vendor
		}
		if inst.Model != inst.original.Model {
			kv["model"] = inst.Model
		}
		if inst.StylePreset != inst.original.StylePreset {
			kv["style_preset"] = inst.StylePreset
		}
		if inst.Prompt != inst.original.Prompt {
			kv["prompt"] = inst.Prompt
		}
		if inst.BgImage != inst.original.BgImage {
			kv["bg_image"] = inst.BgImage
		}
		if inst.BgEmbeddedImage != inst.original.BgEmbeddedImage {
			kv["bg_embedded_image"] = inst.BgEmbeddedImage
		}
		if inst.Label != inst.original.Label {
			kv["label"] = inst.Label
		}
		if inst.LabelColor != inst.original.LabelColor {
			kv["label_color"] = inst.LabelColor
		}
		if inst.SubmitBtnText != inst.original.SubmitBtnText {
			kv["submit_btn_text"] = inst.SubmitBtnText
		}
		if inst.PromptInputTitle != inst.original.PromptInputTitle {
			kv["prompt_input_title"] = inst.PromptInputTitle
		}
		if inst.WaitSeconds != inst.original.WaitSeconds {
			kv["wait_seconds"] = inst.WaitSeconds
		}
		if inst.ShowImageStyleSelector != inst.original.ShowImageStyleSelector {
			kv["show_image_style_selector"] = inst.ShowImageStyleSelector
		}
		if inst.NoPrompt != inst.original.NoPrompt {
			kv["no_prompt"] = inst.NoPrompt
		}
		if inst.VersionMin != inst.original.VersionMin {
			kv["version_min"] = inst.VersionMin
		}
		if inst.VersionMax != inst.original.VersionMax {
			kv["version_max"] = inst.VersionMax
		}
		if inst.Status != inst.original.Status {
			kv["status"] = inst.Status
		}
		if inst.Priority != inst.original.Priority {
			kv["priority"] = inst.Priority
		}
		if inst.Ext != inst.original.Ext {
			kv["ext"] = inst.Ext
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			kv["created_at"] = inst.CreatedAt
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			kv["updated_at"] = inst.UpdatedAt
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					kv["id"] = inst.Id
				}
			case "island_id":
				if inst.IslandId != inst.original.IslandId {
					kv["island_id"] = inst.IslandId
				}
			case "title":
				if inst.Title != inst.original.Title {
					kv["title"] = inst.Title
				}
			case "title_color":
				if inst.TitleColor != inst.original.TitleColor {
					kv["title_color"] = inst.TitleColor
				}
			case "description":
				if inst.Description != inst.original.Description {
					kv["description"] = inst.Description
				}
			case "category":
				if inst.Category != inst.original.Category {
					kv["category"] = inst.Category
				}
			case "model_type":
				if inst.ModelType != inst.original.ModelType {
					kv["model_type"] = inst.ModelType
				}
			case "word_count":
				if inst.WordCount != inst.original.WordCount {
					kv["word_count"] = inst.WordCount
				}
			case "hint":
				if inst.Hint != inst.original.Hint {
					kv["hint"] = inst.Hint
				}
			case "vendor":
				if inst.Vendor != inst.original.Vendor {
					kv["vendor"] = inst.Vendor
				}
			case "model":
				if inst.Model != inst.original.Model {
					kv["model"] = inst.Model
				}
			case "style_preset":
				if inst.StylePreset != inst.original.StylePreset {
					kv["style_preset"] = inst.StylePreset
				}
			case "prompt":
				if inst.Prompt != inst.original.Prompt {
					kv["prompt"] = inst.Prompt
				}
			case "bg_image":
				if inst.BgImage != inst.original.BgImage {
					kv["bg_image"] = inst.BgImage
				}
			case "bg_embedded_image":
				if inst.BgEmbeddedImage != inst.original.BgEmbeddedImage {
					kv["bg_embedded_image"] = inst.BgEmbeddedImage
				}
			case "label":
				if inst.Label != inst.original.Label {
					kv["label"] = inst.Label
				}
			case "label_color":
				if inst.LabelColor != inst.original.LabelColor {
					kv["label_color"] = inst.LabelColor
				}
			case "submit_btn_text":
				if inst.SubmitBtnText != inst.original.SubmitBtnText {
					kv["submit_btn_text"] = inst.SubmitBtnText
				}
			case "prompt_input_title":
				if inst.PromptInputTitle != inst.original.PromptInputTitle {
					kv["prompt_input_title"] = inst.PromptInputTitle
				}
			case "wait_seconds":
				if inst.WaitSeconds != inst.original.WaitSeconds {
					kv["wait_seconds"] = inst.WaitSeconds
				}
			case "show_image_style_selector":
				if inst.ShowImageStyleSelector != inst.original.ShowImageStyleSelector {
					kv["show_image_style_selector"] = inst.ShowImageStyleSelector
				}
			case "no_prompt":
				if inst.NoPrompt != inst.original.NoPrompt {
					kv["no_prompt"] = inst.NoPrompt
				}
			case "version_min":
				if inst.VersionMin != inst.original.VersionMin {
					kv["version_min"] = inst.VersionMin
				}
			case "version_max":
				if inst.VersionMax != inst.original.VersionMax {
					kv["version_max"] = inst.VersionMax
				}
			case "status":
				if inst.Status != inst.original.Status {
					kv["status"] = inst.Status
				}
			case "priority":
				if inst.Priority != inst.original.Priority {
					kv["priority"] = inst.Priority
				}
			case "ext":
				if inst.Ext != inst.original.Ext {
					kv["ext"] = inst.Ext
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					kv["created_at"] = inst.CreatedAt
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					kv["updated_at"] = inst.UpdatedAt
				}
			default:
			}
		}
	}

	return kv
}

// Save create a new model or update it
func (inst *CreativeIslandN) Save(ctx context.Context, onlyFields ...string) error {
	if inst.creativeIslandModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.creativeIslandModel.SaveOrUpdate(ctx, *inst, onlyFields...)
	if err != nil {
		return err
	}

	inst.Id = null.IntFrom(id)
	return nil
}

// Delete remove a creative_island
func (inst *CreativeIslandN) Delete(ctx context.Context) error {
	if inst.creativeIslandModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.creativeIslandModel.DeleteById(ctx, inst.Id.Int64)
	if err != nil {
		return err
	}

	return nil
}

// String convert instance to json string
func (inst *CreativeIslandN) String() string {
	rs, _ := json.Marshal(inst)
	return string(rs)
}

type creativeIslandScope struct {
	name  string
	apply func(builder query.Condition)
}

var creativeIslandGlobalScopes = make([]creativeIslandScope, 0)
var creativeIslandLocalScopes = make([]creativeIslandScope, 0)

// AddGlobalScopeForCreativeIsland assign a global scope to a model
func AddGlobalScopeForCreativeIsland(name string, apply func(builder query.Condition)) {
	creativeIslandGlobalScopes = append(creativeIslandGlobalScopes, creativeIslandScope{name: name, apply: apply})
}

// AddLocalScopeForCreativeIsland assign a local scope to a model
func AddLocalScopeForCreativeIsland(name string, apply func(builder query.Condition)) {
	creativeIslandLocalScopes = append(creativeIslandLocalScopes, creativeIslandScope{name: name, apply: apply})
}

func (m *CreativeIslandModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range creativeIslandGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range creativeIslandLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *CreativeIslandModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *CreativeIslandModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}

	return true
}

type CreativeIsland struct {
	Id                     int64  `json:"id"`
	IslandId               string `json:"island_id"`
	Title                  string `json:"title"`
	TitleColor             string `json:"title_color"`
	Description            string `json:"description"`
	Category               string `json:"category"`
	ModelType              string `json:"model_type"`
	WordCount              int64  `json:"word_count"`
	Hint                   string `json:"hint"`
	Vendor                 string `json:"vendor"`
	Model                  string `json:"model"`
	StylePreset            string `json:"style_preset,omitempty"`
	Prompt                 string `json:"prompt"`
	BgImage                string `json:"bg_image,omitempty"`
	BgEmbeddedImage        string `json:"bg_embedded_image,omitempty"`
	Label                  string `json:"label,omitempty"`
	LabelColor             string `json:"label_color,omitempty"`
	SubmitBtnText          string `json:"submit_btn_text,omitempty"`
	PromptInputTitle       string `json:"prompt_input_title,omitempty"`
	WaitSeconds            int64  `json:"wait_seconds,omitempty"`
	ShowImageStyleSelector int64  `json:"show_image_style_selector,omitempty"`
	NoPrompt               int64  `json:"no_prompt,omitempty"`
	VersionMin             string `json:"version_min,omitempty"`
	VersionMax             string `json:"version_max,omitempty"`
	Status                 int64  `json:"status"`
	Priority               int64  `json:"priority,omitempty"`
	Ext                    string `json:"ext,omitempty"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

func (w CreativeIsland) ToCreativeIslandN(allows ...string) CreativeIslandN {
	if len(allows) == 0 {
		return CreativeIslandN{

			Id:                     null.IntFrom(int64(w.Id)),
			IslandId:               null.StringFrom(w.IslandId),
			Title:                  null.StringFrom(w.Title),
			TitleColor:             null.StringFrom(w.TitleColor),
			Description:            null.StringFrom(w.Description),
			Category:               null.StringFrom(w.Category),
			ModelType:              null.StringFrom(w.ModelType),
			WordCount:              null.IntFrom(int64(w.WordCount)),
			Hint:                   null.StringFrom(w.Hint),
			Vendor:                 null.StringFrom(w.Vendor),
			Model:                  null.StringFrom(w.Model),
			StylePreset:            null.StringFrom(w.StylePreset),
			Prompt:                 null.StringFrom(w.Prompt),
			BgImage:                null.StringFrom(w.BgImage),
			BgEmbeddedImage:        null.StringFrom(w.BgEmbeddedImage),
			Label:                  null.StringFrom(w.Label),
			LabelColor:             null.StringFrom(w.LabelColor),
			SubmitBtnText:          null.StringFrom(w.SubmitBtnText),
			PromptInputTitle:       null.StringFrom(w.PromptInputTitle),
			WaitSeconds:            null.IntFrom(int64(w.WaitSeconds)),
			ShowImageStyleSelector: null.IntFrom(int64(w.ShowImageStyleSelector)),
			NoPrompt:               null.IntFrom(int64(w.NoPrompt)),
			VersionMin:             null.StringFrom(w.VersionMin),
			VersionMax:             null.StringFrom(w.VersionMax),
			Status:                 null.IntFrom(int64(w.Status)),
			Priority:               null.IntFrom(int64(w.Priority)),
			Ext:                    null.StringFrom(w.Ext),
			CreatedAt:              null.TimeFrom(w.CreatedAt),
			UpdatedAt:              null.TimeFrom(w.UpdatedAt),
		}
	}

	res := CreativeIslandN{}
	for _, al := range allows {
		switch strcase.ToSnake(al) {

		case "id":
			res.Id = null.IntFrom(int64(w.Id))
		case "island_id":
			res.IslandId = null.StringFrom(w.IslandId)
		case "title":
			res.Title = null.StringFrom(w.Title)
		case "title_color":
			res.TitleColor = null.StringFrom(w.TitleColor)
		case "description":
			res.Description = null.StringFrom(w.Description)
		case "category":
			res.Category = null.StringFrom(w.Category)
		case "model_type":
			res.ModelType = null.StringFrom(w.ModelType)
		case "word_count":
			res.WordCount = null.IntFrom(int64(w.WordCount))
		case "hint":
			res.Hint = null.StringFrom(w.Hint)
		case "vendor":
			res.Vendor = null.StringFrom(w.Vendor)
		case "model":
			res.Model = null.StringFrom(w.Model)
		case "style_preset":
			res.StylePreset = null.StringFrom(w.StylePreset)
		case "prompt":
			res.Prompt = null.StringFrom(w.Prompt)
		case "bg_image":
			res.BgImage = null.StringFrom(w.BgImage)
		case "bg_embedded_image":
			res.BgEmbeddedImage = null.StringFrom(w.BgEmbeddedImage)
		case "label":
			res.Label = null.StringFrom(w.Label)
		case "label_color":
			res.LabelColor = null.StringFrom(w.LabelColor)
		case "submit_btn_text":
			res.SubmitBtnText = null.StringFrom(w.SubmitBtnText)
		case "prompt_input_title":
			res.PromptInputTitle = null.StringFrom(w.PromptInputTitle)
		case "wait_seconds":
			res.WaitSeconds = null.IntFrom(int64(w.WaitSeconds))
		case "show_image_style_selector":
			res.ShowImageStyleSelector = null.IntFrom(int64(w.ShowImageStyleSelector))
		case "no_prompt":
			res.NoPrompt = null.IntFrom(int64(w.NoPrompt))
		case "version_min":
			res.VersionMin = null.StringFrom(w.VersionMin)
		case "version_max":
			res.VersionMax = null.StringFrom(w.VersionMax)
		case "status":
			res.Status = null.IntFrom(int64(w.Status))
		case "priority":
			res.Priority = null.IntFrom(int64(w.Priority))
		case "ext":
			res.Ext = null.StringFrom(w.Ext)
		case "created_at":
			res.CreatedAt = null.TimeFrom(w.CreatedAt)
		case "updated_at":
			res.UpdatedAt = null.TimeFrom(w.UpdatedAt)
		default:
		}
	}

	return res
}

// As convert object to other type
// dst must be a pointer to struct
func (w CreativeIsland) As(dst interface{}) error {
	return query.Copy(w, dst)
}

func (w *CreativeIslandN) ToCreativeIsland() CreativeIsland {
	return CreativeIsland{

		Id:                     w.Id.Int64,
		IslandId:               w.IslandId.String,
		Title:                  w.Title.String,
		TitleColor:             w.TitleColor.String,
		Description:            w.Description.String,
		Category:               w.Category.String,
		ModelType:              w.ModelType.String,
		WordCount:              w.WordCount.Int64,
		Hint:                   w.Hint.String,
		Vendor:                 w.Vendor.String,
		Model:                  w.Model.String,
		StylePreset:            w.StylePreset.String,
		Prompt:                 w.Prompt.String,
		BgImage:                w.BgImage.String,
		BgEmbeddedImage:        w.BgEmbeddedImage.String,
		Label:                  w.Label.String,
		LabelColor:             w.LabelColor.String,
		SubmitBtnText:          w.SubmitBtnText.String,
		PromptInputTitle:       w.PromptInputTitle.String,
		WaitSeconds:            w.WaitSeconds.Int64,
		ShowImageStyleSelector: w.ShowImageStyleSelector.Int64,
		NoPrompt:               w.NoPrompt.Int64,
		VersionMin:             w.VersionMin.String,
		VersionMax:             w.VersionMax.String,
		Status:                 w.Status.Int64,
		Priority:               w.Priority.Int64,
		Ext:                    w.Ext.String,
		CreatedAt:              w.CreatedAt.Time,
		UpdatedAt:              w.UpdatedAt.Time,
	}
}

// CreativeIslandModel is a model which encapsulates the operations of the object
type CreativeIslandModel struct {
	db        *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes  []string

	query query.SQLBuilder
}

var creativeIslandTableName = "creative_island"

// CreativeIslandTable return table name for CreativeIsland
func CreativeIslandTable() string {
	return creativeIslandTableName
}

const (
	FieldCreativeIslandId                     = "id"
	FieldCreativeIslandIslandId               = "island_id"
	FieldCreativeIslandTitle                  = "title"
	FieldCreativeIslandTitleColor             = "title_color"
	FieldCreativeIslandDescription            = "description"
	FieldCreativeIslandCategory               = "category"
	FieldCreativeIslandModelType              = "model_type"
	FieldCreativeIslandWordCount              = "word_count"
	FieldCreativeIslandHint                   = "hint"
	FieldCreativeIslandVendor                 = "vendor"
	FieldCreativeIslandModel                  = "model"
	FieldCreativeIslandStylePreset            = "style_preset"
	FieldCreativeIslandPrompt                 = "prompt"
	FieldCreativeIslandBgImage                = "bg_image"
	FieldCreativeIslandBgEmbeddedImage        = "bg_embedded_image"
	FieldCreativeIslandLabel                  = "label"
	FieldCreativeIslandLabelColor             = "label_color"
	FieldCreativeIslandSubmitBtnText          = "submit_btn_text"
	FieldCreativeIslandPromptInputTitle       = "prompt_input_title"
	FieldCreativeIslandWaitSeconds            = "wait_seconds"
	FieldCreativeIslandShowImageStyleSelector = "show_image_style_selector"
	FieldCreativeIslandNoPrompt               = "no_prompt"
	FieldCreativeIslandVersionMin             = "version_min"
	FieldCreativeIslandVersionMax             = "version_max"
	FieldCreativeIslandStatus                 = "status"
	FieldCreativeIslandPriority               = "priority"
	FieldCreativeIslandExt                    = "ext"
	FieldCreativeIslandCreatedAt              = "created_at"
	FieldCreativeIslandUpdatedAt              = "updated_at"
)

// CreativeIslandFields return all fields in CreativeIsland model
func CreativeIslandFields() []string {
	return []string{
		"id",
		"island_id",
		"title",
		"title_color",
		"description",
		"category",
		"model_type",
		"word_count",
		"hint",
		"vendor",
		"model",
		"style_preset",
		"prompt",
		"bg_image",
		"bg_embedded_image",
		"label",
		"label_color",
		"submit_btn_text",
		"prompt_input_title",
		"wait_seconds",
		"show_image_style_selector",
		"no_prompt",
		"version_min",
		"version_max",
		"status",
		"priority",
		"ext",
		"created_at",
		"updated_at",
	}
}

func SetCreativeIslandTable(tableName string) {
	creativeIslandTableName = tableName
}

// NewCreativeIslandModel create a CreativeIslandModel
func NewCreativeIslandModel(db query.Database) *CreativeIslandModel {
	return &CreativeIslandModel{
		db:                  query.NewDatabaseWrap(db),
		tableName:           creativeIslandTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes:  make([]string, 0),
		query:               query.Builder(),
	}
}

// GetDB return database instance
func (m *CreativeIslandModel) GetDB() query.Database {
	return m.db.GetDB()
}

func (m *CreativeIslandModel) clone() *CreativeIslandModel {
	return &CreativeIslandModel{
		db:                  m.db,
		tableName:           m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes:  append([]string{}, m.includeLocalScopes...),
		query:               m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *CreativeIslandModel) WithoutGlobalScopes(names ...string) *CreativeIslandModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *CreativeIslandModel) WithLocalScopes(names ...string) *CreativeIslandModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Condition add query builder to model
func (m *CreativeIslandModel) Condition(builder query.SQLBuilder) *CreativeIslandModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *CreativeIslandModel) Find(ctx context.Context, id int64) (*CreativeIslandN, error) {
	return m.First(ctx, m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *CreativeIslandModel) Exists(ctx context.Context, builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(ctx, builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *CreativeIslandModel) Count(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {
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

func (m *CreativeIslandModel) Paginate(ctx context.Context, page int64, perPage int64, builders ...query.SQLBuilder) ([]CreativeIslandN, query.PaginateMeta, error) {
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
func (m *CreativeIslandModel) Get(ctx context.Context, builders ...query.SQLBuilder) ([]CreativeIslandN, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"id",
			"island_id",
			"title",
			"title_color",
			"description",
			"category",
			"model_type",
			"word_count",
			"hint",
			"vendor",
			"model",
			"style_preset",
			"prompt",
			"bg_image",
			"bg_embedded_image",
			"label",
			"label_color",
			"submit_btn_text",
			"prompt_input_title",
			"wait_seconds",
			"show_image_style_selector",
			"no_prompt",
			"version_min",
			"version_max",
			"status",
			"priority",
			"ext",
			"created_at",
			"updated_at",
		)
	}

	fields := b.GetFields()
	selectFields := make([]query.Expr, 0)

	for _, f := range fields {
		switch strcase.ToSnake(f.Value) {

		case "id":
			selectFields = append(selectFields, f)
		case "island_id":
			selectFields = append(selectFields, f)
		case "title":
			selectFields = append(selectFields, f)
		case "title_color":
			selectFields = append(selectFields, f)
		case "description":
			selectFields = append(selectFields, f)
		case "category":
			selectFields = append(selectFields, f)
		case "model_type":
			selectFields = append(selectFields, f)
		case "word_count":
			selectFields = append(selectFields, f)
		case "hint":
			selectFields = append(selectFields, f)
		case "vendor":
			selectFields = append(selectFields, f)
		case "model":
			selectFields = append(selectFields, f)
		case "style_preset":
			selectFields = append(selectFields, f)
		case "prompt":
			selectFields = append(selectFields, f)
		case "bg_image":
			selectFields = append(selectFields, f)
		case "bg_embedded_image":
			selectFields = append(selectFields, f)
		case "label":
			selectFields = append(selectFields, f)
		case "label_color":
			selectFields = append(selectFields, f)
		case "submit_btn_text":
			selectFields = append(selectFields, f)
		case "prompt_input_title":
			selectFields = append(selectFields, f)
		case "wait_seconds":
			selectFields = append(selectFields, f)
		case "show_image_style_selector":
			selectFields = append(selectFields, f)
		case "no_prompt":
			selectFields = append(selectFields, f)
		case "version_min":
			selectFields = append(selectFields, f)
		case "version_max":
			selectFields = append(selectFields, f)
		case "status":
			selectFields = append(selectFields, f)
		case "priority":
			selectFields = append(selectFields, f)
		case "ext":
			selectFields = append(selectFields, f)
		case "created_at":
			selectFields = append(selectFields, f)
		case "updated_at":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*CreativeIslandN, []interface{}) {
		var creativeIslandVar CreativeIslandN
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {

			case "id":
				scanFields = append(scanFields, &creativeIslandVar.Id)
			case "island_id":
				scanFields = append(scanFields, &creativeIslandVar.IslandId)
			case "title":
				scanFields = append(scanFields, &creativeIslandVar.Title)
			case "title_color":
				scanFields = append(scanFields, &creativeIslandVar.TitleColor)
			case "description":
				scanFields = append(scanFields, &creativeIslandVar.Description)
			case "category":
				scanFields = append(scanFields, &creativeIslandVar.Category)
			case "model_type":
				scanFields = append(scanFields, &creativeIslandVar.ModelType)
			case "word_count":
				scanFields = append(scanFields, &creativeIslandVar.WordCount)
			case "hint":
				scanFields = append(scanFields, &creativeIslandVar.Hint)
			case "vendor":
				scanFields = append(scanFields, &creativeIslandVar.Vendor)
			case "model":
				scanFields = append(scanFields, &creativeIslandVar.Model)
			case "style_preset":
				scanFields = append(scanFields, &creativeIslandVar.StylePreset)
			case "prompt":
				scanFields = append(scanFields, &creativeIslandVar.Prompt)
			case "bg_image":
				scanFields = append(scanFields, &creativeIslandVar.BgImage)
			case "bg_embedded_image":
				scanFields = append(scanFields, &creativeIslandVar.BgEmbeddedImage)
			case "label":
				scanFields = append(scanFields, &creativeIslandVar.Label)
			case "label_color":
				scanFields = append(scanFields, &creativeIslandVar.LabelColor)
			case "submit_btn_text":
				scanFields = append(scanFields, &creativeIslandVar.SubmitBtnText)
			case "prompt_input_title":
				scanFields = append(scanFields, &creativeIslandVar.PromptInputTitle)
			case "wait_seconds":
				scanFields = append(scanFields, &creativeIslandVar.WaitSeconds)
			case "show_image_style_selector":
				scanFields = append(scanFields, &creativeIslandVar.ShowImageStyleSelector)
			case "no_prompt":
				scanFields = append(scanFields, &creativeIslandVar.NoPrompt)
			case "version_min":
				scanFields = append(scanFields, &creativeIslandVar.VersionMin)
			case "version_max":
				scanFields = append(scanFields, &creativeIslandVar.VersionMax)
			case "status":
				scanFields = append(scanFields, &creativeIslandVar.Status)
			case "priority":
				scanFields = append(scanFields, &creativeIslandVar.Priority)
			case "ext":
				scanFields = append(scanFields, &creativeIslandVar.Ext)
			case "created_at":
				scanFields = append(scanFields, &creativeIslandVar.CreatedAt)
			case "updated_at":
				scanFields = append(scanFields, &creativeIslandVar.UpdatedAt)
			}
		}

		return &creativeIslandVar, scanFields
	}

	sqlStr, params := b.Fields(selectFields...).ResolveQuery()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	creativeIslands := make([]CreativeIslandN, 0)
	for rows.Next() {
		creativeIslandReal, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		creativeIslandReal.original = &creativeIslandOriginal{}
		_ = query.Copy(creativeIslandReal, creativeIslandReal.original)

		creativeIslandReal.SetModel(m)
		creativeIslands = append(creativeIslands, *creativeIslandReal)
	}

	return creativeIslands, nil
}

// First return first result for given query
func (m *CreativeIslandModel) First(ctx context.Context, builders ...query.SQLBuilder) (*CreativeIslandN, error) {
	res, err := m.Get(ctx, append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, query.ErrNoResult
	}

	return &res[0], nil
}

// Create save a new creative_island to database
func (m *CreativeIslandModel) Create(ctx context.Context, kv query.KV) (int64, error) {

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

// SaveAll save all creative_islands to database
func (m *CreativeIslandModel) SaveAll(ctx context.Context, creativeIslands []CreativeIslandN) ([]int64, error) {
	ids := make([]int64, 0)
	for _, creativeIsland := range creativeIslands {
		id, err := m.Save(ctx, creativeIsland)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a creative_island to database
func (m *CreativeIslandModel) Save(ctx context.Context, creativeIsland CreativeIslandN, onlyFields ...string) (int64, error) {
	return m.Create(ctx, creativeIsland.StaledKV(onlyFields...))
}

// SaveOrUpdate save a new creative_island or update it when it has a id > 0
func (m *CreativeIslandModel) SaveOrUpdate(ctx context.Context, creativeIsland CreativeIslandN, onlyFields ...string) (id int64, updated bool, err error) {
	if creativeIsland.Id.Int64 > 0 {
		_, _err := m.UpdateById(ctx, creativeIsland.Id.Int64, creativeIsland, onlyFields...)
		return creativeIsland.Id.Int64, true, _err
	}

	_id, _err := m.Save(ctx, creativeIsland, onlyFields...)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *CreativeIslandModel) UpdateFields(ctx context.Context, kv query.KV, builders ...query.SQLBuilder) (int64, error) {
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
func (m *CreativeIslandModel) Update(ctx context.Context, builder query.SQLBuilder, creativeIsland CreativeIslandN, onlyFields ...string) (int64, error) {
	return m.UpdateFields(ctx, creativeIsland.StaledKV(onlyFields...), builder)
}

// UpdateById update a model by id
func (m *CreativeIslandModel) UpdateById(ctx context.Context, id int64, creativeIsland CreativeIslandN, onlyFields ...string) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).UpdateFields(ctx, creativeIsland.StaledKV(onlyFields...))
}

// Delete remove a model
func (m *CreativeIslandModel) Delete(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).Table(m.tableName).ResolveDelete()

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

// DeleteById remove a model by id
func (m *CreativeIslandModel) DeleteById(ctx context.Context, id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Delete(ctx)
}