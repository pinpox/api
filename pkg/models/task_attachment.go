// Vikunja is a to-do list application to facilitate your life.
// Copyright 2018-2020 Vikunja and contributors. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package models

import (
	"code.vikunja.io/api/pkg/files"
	"code.vikunja.io/api/pkg/user"
	"code.vikunja.io/web"
	"io"
	"time"
)

// TaskAttachment is the definition of a task attachment
type TaskAttachment struct {
	ID     int64 `xorm:"int(11) autoincr not null unique pk" json:"id" param:"attachment"`
	TaskID int64 `xorm:"int(11) not null" json:"task_id" param:"task"`
	FileID int64 `xorm:"int(11) not null" json:"-"`

	CreatedByID int64      `xorm:"int(11) not null" json:"-"`
	CreatedBy   *user.User `xorm:"-" json:"created_by"`

	File *files.File `xorm:"-" json:"file"`

	Created time.Time `xorm:"created" json:"created"`

	web.CRUDable `xorm:"-" json:"-"`
	web.Rights   `xorm:"-" json:"-"`
}

// TableName returns the table name for task attachments
func (TaskAttachment) TableName() string {
	return "task_attachments"
}

// NewAttachment creates a new task attachment
// Note: I'm not sure if only accepting an io.ReadCloser and not an afero.File or os.File instead is a good way of doing things.
func (ta *TaskAttachment) NewAttachment(f io.ReadCloser, realname string, realsize uint64, a web.Auth) error {

	// Store the file
	file, err := files.Create(f, realname, realsize, a)
	if err != nil {
		if files.IsErrFileIsTooLarge(err) {
			return ErrTaskAttachmentIsTooLarge{Size: realsize}
		}
		return err
	}
	ta.File = file

	// Add an entry to the db
	ta.FileID = file.ID
	ta.CreatedByID = a.GetID()
	_, err = x.Insert(ta)
	if err != nil {
		// remove the  uploaded file if adding it to the db fails
		if err2 := file.Delete(); err2 != nil {
			return err2
		}
		return err
	}

	return nil
}

// ReadOne returns a task attachment
func (ta *TaskAttachment) ReadOne() (err error) {
	exists, err := x.Where("id = ?", ta.ID).Get(ta)
	if err != nil {
		return
	}
	if !exists {
		return ErrTaskAttachmentDoesNotExist{
			TaskID:       ta.TaskID,
			AttachmentID: ta.ID,
		}
	}

	// Get the file
	ta.File = &files.File{ID: ta.FileID}
	err = ta.File.LoadFileMetaByID()
	return
}

// ReadAll returns a list with all attachments
// @Summary Get  all attachments for one task.
// @Description Get all task attachments for one task.
// @tags task
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param page query int false "The page number. Used for pagination. If not provided, the first page of results is returned."
// @Param per_page query int false "The maximum number of items per page. Note this parameter is limited by the configured maximum of items per page."
// @Security JWTKeyAuth
// @Success 200 {array} models.TaskAttachment "All attachments for this task"
// @Failure 403 {object} models.Message "No access to this task."
// @Failure 404 {object} models.Message "The task does not exist."
// @Failure 500 {object} models.Message "Internal error"
// @Router /tasks/{id}/attachments [get]
func (ta *TaskAttachment) ReadAll(a web.Auth, search string, page int, perPage int) (result interface{}, resultCount int, numberOfTotalItems int64, err error) {
	attachments := []*TaskAttachment{}

	limit, start := getLimitFromPageIndex(page, perPage)

	query := x.
		Where("task_id = ?", ta.TaskID)
	if limit > 0 {
		query = query.Limit(limit, start)
	}
	err = query.Find(&attachments)
	if err != nil {
		return nil, 0, 0, err
	}

	fileIDs := make([]int64, 0, len(attachments))
	userIDs := make([]int64, 0, len(attachments))
	for _, r := range attachments {
		fileIDs = append(fileIDs, r.FileID)
		userIDs = append(userIDs, r.CreatedByID)
	}

	fs := make(map[int64]*files.File)
	err = x.In("id", fileIDs).Find(&fs)
	if err != nil {
		return nil, 0, 0, err
	}

	us := make(map[int64]*user.User)
	err = x.In("id", userIDs).Find(&us)
	if err != nil {
		return nil, 0, 0, err
	}

	for _, r := range attachments {
		// If the actual file does not exist, don't try to load it as that would fail with nil panic
		if _, exists := fs[r.FileID]; !exists {
			continue
		}
		r.File = fs[r.FileID]
		r.CreatedBy = us[r.CreatedByID]
	}

	numberOfTotalItems, err = x.
		Where("task_id = ?", ta.TaskID).
		Count(&TaskAttachment{})
	return attachments, len(attachments), numberOfTotalItems, err
}

// Delete removes an attachment
// @Summary Delete an attachment
// @Description Delete an attachment.
// @tags task
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param attachmentID path int true "Attachment ID"
// @Security JWTKeyAuth
// @Success 200 {object} models.Message "The attachment was deleted successfully."
// @Failure 403 {object} models.Message "No access to this task."
// @Failure 404 {object} models.Message "The task does not exist."
// @Failure 500 {object} models.Message "Internal error"
// @Router /tasks/{id}/attachments/{attachmentID} [delete]
func (ta *TaskAttachment) Delete() error {
	// Load the attachment
	err := ta.ReadOne()
	if err != nil && !files.IsErrFileDoesNotExist(err) {
		return err
	}

	// Delete it
	_, err = x.Where("task_id = ? AND id = ?", ta.TaskID, ta.ID).Delete(ta)
	if err != nil {
		return err
	}

	// Delete the underlying file
	err = ta.File.Delete()
	// If the file does not exist, we don't want to error out
	if err != nil && files.IsErrFileDoesNotExist(err) {
		return nil
	}
	return err
}

func getTaskAttachmentsByTaskIDs(taskIDs []int64) (attachments []*TaskAttachment, err error) {
	attachments = []*TaskAttachment{}
	err = x.
		In("task_id", taskIDs).
		Find(&attachments)
	if err != nil {
		return
	}

	fileIDs := []int64{}
	userIDs := []int64{}
	for _, a := range attachments {
		userIDs = append(userIDs, a.CreatedByID)
		fileIDs = append(fileIDs, a.FileID)
	}

	// Get all files
	fs := make(map[int64]*files.File)
	err = x.In("id", fileIDs).Find(&fs)
	if err != nil {
		return
	}

	users := make(map[int64]*user.User)
	err = x.In("id", userIDs).Find(&users)
	if err != nil {
		return
	}

	// Obfuscate all user emails
	for _, u := range users {
		u.Email = ""
	}

	for _, a := range attachments {
		a.CreatedBy = users[a.CreatedByID]
		a.File = fs[a.FileID]
	}

	return
}
