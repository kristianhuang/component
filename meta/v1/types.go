/*
 * Copyright 2021 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package v1

import (
	"github.com/kristianhuang/go-cmp/json"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Extend map[string]interface{}

func (e Extend) String() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// Merge merge extend fields from extendShadow.
func (e Extend) Merge(extendShadow string) Extend {
	var extend Extend

	_ = json.Unmarshal([]byte(extendShadow), &extend)
	for k, v := range extend {
		if _, ok := e[k]; !ok {
			e[k] = v
		}
	}
	return e
}

type TypeMeta struct {
	Kind       string `json:"kind,omitempty"`
	APIVersion string `json:"api_version,omitempty"`
}

type ListMeta struct {
	TotalCount int64 `json:"total_count,omitempty"`
}

type ListOptions struct {
	TypeMeta       `json:",inline"`
	LabelSelector  string `json:"label_selector,omitempty" form:"label_selector"`
	FieldSelector  string `json:"field_selector,omitempty" form:"field_selector"`
	TimeoutSeconds *int64 `json:"timeout_seconds,omitempty"`
	Offset         *int64 `json:"offset,omitempty" form:"offset"`
	Limit          *int64 `json:"limit,omitempty" form:"limit"`
}

type ObjectMeta struct {
	ID uint64 `json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id"`

	// InstanceID defines a string type resource identifier,
	// use prefixed to distinguish resource types, easy to remember, Url-friendly.
	InstanceID string `json:"instance_id,omitempty" gorm:"unique;column:instance_id;type:varchar(32);not null"`

	// Required: true
	// Name must be unique. Is required when creating resources.
	// Name is primarily intended for creation idempotence and configuration
	// definition.
	// It will be generated automated only if Name is not specified.
	// Cannot be updated.
	Name string `json:"name,omitempty" gorm:"column:name;type:varchar(64);not null;comment:资源名"`

	CreatedAt int64 `json:"created_at,omitempty" gorm:"type:int(11);not null;column:created_at;comment:创建时间;"`
	UpdatedAt int64 `json:"updated_at,omitempty" gorm:"type:int(11);not null;column:updated_at;comment:更新时间;"`

	CreatedAtFormat string `json:"created_at_format,omitempty" gorm:"-"`
	UpdatedAtFormat string `json:"updated_at_format,omitempty" gorm:"-"`

	// 拓展字段
	Extend       Extend `json:"extend,omitempty" gorm:"-"`
	ExtendShadow string `json:"-" gorm:"column:extend_shadow"`
}

// DeleteMeta is delete meta data.
type DeleteMeta struct {
	DeletedAt soft_delete.DeletedAt `json:"deleted_at,omitempty" gorm:"type:int(11);column:deleted_at;comment:删除时间"`
}

// BeforeCreate run before create database record.
func (obj *ObjectMeta) BeforeCreate(tx *gorm.DB) error {
	obj.ExtendShadow = obj.Extend.String()

	return nil
}

// BeforeUpdate run before update database record.
func (obj *ObjectMeta) BeforeUpdate(tx *gorm.DB) error {
	obj.ExtendShadow = obj.Extend.String()

	return nil
}

// AfterFind run after find to unmarshal a extend shadown string into metav1.Extend struct.
func (obj *ObjectMeta) AfterFind(tx *gorm.DB) error {
	if err := json.Unmarshal([]byte(obj.ExtendShadow), &obj.Extend); err != nil {
		return err
	}

	return nil
}

type DeleteOptions struct {
	TypeMeta `json:",inline"`
	Unscoped bool `json:"unscoped"`
}

type CreateOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dry_run,omitempty"`
}

type UpdateOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dry_run,omitempty"`
}

type PatchOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dry_run,omitempty"`
	Force    bool     `json:"force,omitempty"`
}

type AuthorizeOptions struct {
	TypeMeta `json:",inline"`
}

// GetOptions is the standard query options to the standard REST get call.
type GetOptions struct {
	TypeMeta `json:",inline"`
}

type TableOptions struct {
	TypeMeta  `json:",inline"`
	NoHeaders bool `json:"-"`
}
