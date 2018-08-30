package models

import "fmt"

// =====================
// User Operation Errors
// =====================

// ErrUsernameExists represents a "UsernameAlreadyExists" kind of error.
type ErrUsernameExists struct {
	UserID   int64
	Username string
}

// IsErrUsernameExists checks if an error is a ErrUsernameExists.
func IsErrUsernameExists(err error) bool {
	_, ok := err.(ErrUsernameExists)
	return ok
}

func (err ErrUsernameExists) Error() string {
	return fmt.Sprintf("a user with this username does already exist [user id: %d, username: %s]", err.UserID, err.Username)
}

// ErrUserEmailExists represents a "UserEmailExists" kind of error.
type ErrUserEmailExists struct {
	UserID int64
	Email  string
}

// IsErrUserEmailExists checks if an error is a ErrUserEmailExists.
func IsErrUserEmailExists(err error) bool {
	_, ok := err.(ErrUserEmailExists)
	return ok
}

func (err ErrUserEmailExists) Error() string {
	return fmt.Sprintf("a user with this email does already exist [user id: %d, email: %s]", err.UserID, err.Email)
}

// ErrNoUsername represents a "UsernameAlreadyExists" kind of error.
type ErrNoUsername struct {
	UserID int64
}

// IsErrNoUsername checks if an error is a ErrUsernameExists.
func IsErrNoUsername(err error) bool {
	_, ok := err.(ErrNoUsername)
	return ok
}

func (err ErrNoUsername) Error() string {
	return fmt.Sprintf("you need to specify a username [user id: %d]", err.UserID)
}

// ErrNoUsernamePassword represents a "NoUsernamePassword" kind of error.
type ErrNoUsernamePassword struct{}

// IsErrNoUsernamePassword checks if an error is a ErrNoUsernamePassword.
func IsErrNoUsernamePassword(err error) bool {
	_, ok := err.(ErrNoUsernamePassword)
	return ok
}

func (err ErrNoUsernamePassword) Error() string {
	return fmt.Sprintf("you need to specify a username and a password")
}

// ErrUserDoesNotExist represents a "UserDoesNotExist" kind of error.
type ErrUserDoesNotExist struct {
	UserID int64
}

// IsErrUserDoesNotExist checks if an error is a ErrUserDoesNotExist.
func IsErrUserDoesNotExist(err error) bool {
	_, ok := err.(ErrUserDoesNotExist)
	return ok
}

func (err ErrUserDoesNotExist) Error() string {
	return fmt.Sprintf("this user does not exist [user id: %d]", err.UserID)
}

// ErrCouldNotGetUserID represents a "ErrCouldNotGetUserID" kind of error.
type ErrCouldNotGetUserID struct{}

// IsErrCouldNotGetUserID checks if an error is a ErrCouldNotGetUserID.
func IsErrCouldNotGetUserID(err error) bool {
	_, ok := err.(ErrCouldNotGetUserID)
	return ok
}

func (err ErrCouldNotGetUserID) Error() string {
	return fmt.Sprintf("could not get user ID")
}

// ErrCannotDeleteLastUser represents a "ErrCannotDeleteLastUser" kind of error.
type ErrCannotDeleteLastUser struct{}

// IsErrCannotDeleteLastUser checks if an error is a ErrCannotDeleteLastUser.
func IsErrCannotDeleteLastUser(err error) bool {
	_, ok := err.(ErrCannotDeleteLastUser)
	return ok
}

func (err ErrCannotDeleteLastUser) Error() string {
	return fmt.Sprintf("cannot delete last user")
}

// ===================
// Empty things errors
// ===================

// ErrIDCannotBeZero represents a "IDCannotBeZero" kind of error. Used if an ID (of something, not defined) is 0 where it should not.
type ErrIDCannotBeZero struct{}

// IsErrIDCannotBeZero checks if an error is a ErrIDCannotBeZero.
func IsErrIDCannotBeZero(err error) bool {
	_, ok := err.(ErrIDCannotBeZero)
	return ok
}

func (err ErrIDCannotBeZero) Error() string {
	return fmt.Sprintf("ID cannot be 0")
}

// ===========
// List errors
// ===========

// ErrListDoesNotExist represents a "ErrListDoesNotExist" kind of error. Used if the list does not exist.
type ErrListDoesNotExist struct {
	ID int64
}

// IsErrListDoesNotExist checks if an error is a ErrListDoesNotExist.
func IsErrListDoesNotExist(err error) bool {
	_, ok := err.(ErrListDoesNotExist)
	return ok
}

func (err ErrListDoesNotExist) Error() string {
	return fmt.Sprintf("List does not exist [ID: %d]", err.ID)
}

// ErrNeedToBeListAdmin represents an error, where the user is not the owner of that list (used i.e. when deleting a list)
type ErrNeedToBeListAdmin struct {
	ListID int64
	UserID int64
}

// IsErrNeedToBeListAdmin checks if an error is a ErrListDoesNotExist.
func IsErrNeedToBeListAdmin(err error) bool {
	_, ok := err.(ErrNeedToBeListAdmin)
	return ok
}

func (err ErrNeedToBeListAdmin) Error() string {
	return fmt.Sprintf("You need to be list owner to do that [ListID: %d, UserID: %d]", err.ListID, err.UserID)
}

// ErrNeedToBeListWriter represents an error, where the user is not the owner of that list (used i.e. when deleting a list)
type ErrNeedToBeListWriter struct {
	ListID int64
	UserID int64
}

// IsErrNeedToBeListWriter checks if an error is a ErrListDoesNotExist.
func IsErrNeedToBeListWriter(err error) bool {
	_, ok := err.(ErrNeedToBeListWriter)
	return ok
}

func (err ErrNeedToBeListWriter) Error() string {
	return fmt.Sprintf("You need to have write acces to the list to do that [ListID: %d, UserID: %d]", err.ListID, err.UserID)
}

// ErrNeedToHaveListReadAccess represents an error, where the user dont has read access to that List
type ErrNeedToHaveListReadAccess struct {
	ListID int64
	UserID int64
}

// IsErrNeedToHaveListReadAccess checks if an error is a ErrListDoesNotExist.
func IsErrNeedToHaveListReadAccess(err error) bool {
	_, ok := err.(ErrNeedToHaveListReadAccess)
	return ok
}

func (err ErrNeedToHaveListReadAccess) Error() string {
	return fmt.Sprintf("You need to be List owner to do that [ListID: %d, UserID: %d]", err.ListID, err.UserID)
}

// ErrListTitleCannotBeEmpty represents a "ErrListTitleCannotBeEmpty" kind of error. Used if the list does not exist.
type ErrListTitleCannotBeEmpty struct{}

// IsErrListTitleCannotBeEmpty checks if an error is a ErrListTitleCannotBeEmpty.
func IsErrListTitleCannotBeEmpty(err error) bool {
	_, ok := err.(ErrListTitleCannotBeEmpty)
	return ok
}

func (err ErrListTitleCannotBeEmpty) Error() string {
	return fmt.Sprintf("List task text cannot be empty.")
}

// ================
// List task errors
// ================

// ErrListTaskCannotBeEmpty represents a "ErrListDoesNotExist" kind of error. Used if the list does not exist.
type ErrListTaskCannotBeEmpty struct{}

// IsErrListTaskCannotBeEmpty checks if an error is a ErrListDoesNotExist.
func IsErrListTaskCannotBeEmpty(err error) bool {
	_, ok := err.(ErrListTaskCannotBeEmpty)
	return ok
}

func (err ErrListTaskCannotBeEmpty) Error() string {
	return fmt.Sprintf("List task text cannot be empty.")
}

// ErrListTaskDoesNotExist represents a "ErrListDoesNotExist" kind of error. Used if the list does not exist.
type ErrListTaskDoesNotExist struct {
	ID int64
}

// IsErrListTaskDoesNotExist checks if an error is a ErrListDoesNotExist.
func IsErrListTaskDoesNotExist(err error) bool {
	_, ok := err.(ErrListTaskDoesNotExist)
	return ok
}

func (err ErrListTaskDoesNotExist) Error() string {
	return fmt.Sprintf("List task does not exist. [ID: %d]", err.ID)
}

// ErrNeedToBeTaskOwner represents an error, where the user is not the owner of that task (used i.e. when deleting a list)
type ErrNeedToBeTaskOwner struct {
	TaskID int64
	UserID int64
}

// IsErrNeedToBeTaskOwner checks if an error is a ErrNeedToBeTaskOwner.
func IsErrNeedToBeTaskOwner(err error) bool {
	_, ok := err.(ErrNeedToBeTaskOwner)
	return ok
}

func (err ErrNeedToBeTaskOwner) Error() string {
	return fmt.Sprintf("You need to be task owner to do that [TaskID: %d, UserID: %d]", err.TaskID, err.UserID)
}

// =================
// Namespace errors
// =================

// ErrNamespaceDoesNotExist represents a "ErrNamespaceDoesNotExist" kind of error. Used if the namespace does not exist.
type ErrNamespaceDoesNotExist struct {
	ID int64
}

// IsErrNamespaceDoesNotExist checks if an error is a ErrNamespaceDoesNotExist.
func IsErrNamespaceDoesNotExist(err error) bool {
	_, ok := err.(ErrNamespaceDoesNotExist)
	return ok
}

func (err ErrNamespaceDoesNotExist) Error() string {
	return fmt.Sprintf("Namespace does not exist [ID: %d]", err.ID)
}

// ErrNeedToBeNamespaceOwner represents an error, where the user is not the owner of that namespace (used i.e. when deleting a namespace)
type ErrNeedToBeNamespaceOwner struct {
	NamespaceID int64
	UserID      int64
}

// IsErrNeedToBeNamespaceOwner checks if an error is a ErrNamespaceDoesNotExist.
func IsErrNeedToBeNamespaceOwner(err error) bool {
	_, ok := err.(ErrNeedToBeNamespaceOwner)
	return ok
}

func (err ErrNeedToBeNamespaceOwner) Error() string {
	return fmt.Sprintf("You need to be namespace owner to do that [NamespaceID: %d, UserID: %d]", err.NamespaceID, err.UserID)
}

// ErrUserDoesNotHaveAccessToNamespace represents an error, where the user is not the owner of that namespace (used i.e. when deleting a namespace)
type ErrUserDoesNotHaveAccessToNamespace struct {
	NamespaceID int64
	UserID      int64
}

// IsErrUserDoesNotHaveAccessToNamespace checks if an error is a ErrNamespaceDoesNotExist.
func IsErrUserDoesNotHaveAccessToNamespace(err error) bool {
	_, ok := err.(ErrUserDoesNotHaveAccessToNamespace)
	return ok
}

func (err ErrUserDoesNotHaveAccessToNamespace) Error() string {
	return fmt.Sprintf("You need to have access to this namespace to do that [NamespaceID: %d, UserID: %d]", err.NamespaceID, err.UserID)
}

// ErrUserNeedsToBeNamespaceAdmin represents an error, where the user is not the owner of that namespace (used i.e. when deleting a namespace)
type ErrUserNeedsToBeNamespaceAdmin struct {
	NamespaceID int64
	UserID      int64
}

// IsErrUserNeedsToBeNamespaceAdmin checks if an error is a ErrNamespaceDoesNotExist.
func IsErrUserNeedsToBeNamespaceAdmin(err error) bool {
	_, ok := err.(ErrUserNeedsToBeNamespaceAdmin)
	return ok
}

func (err ErrUserNeedsToBeNamespaceAdmin) Error() string {
	return fmt.Sprintf("You need to be namespace admin to do that [NamespaceID: %d, UserID: %d]", err.NamespaceID, err.UserID)
}

// ErrUserDoesNotHaveWriteAccessToNamespace represents an error, where the user is not the owner of that namespace (used i.e. when deleting a namespace)
type ErrUserDoesNotHaveWriteAccessToNamespace struct {
	NamespaceID int64
	UserID      int64
}

// IsErrUserDoesNotHaveWriteAccessToNamespace checks if an error is a ErrNamespaceDoesNotExist.
func IsErrUserDoesNotHaveWriteAccessToNamespace(err error) bool {
	_, ok := err.(ErrUserDoesNotHaveWriteAccessToNamespace)
	return ok
}

func (err ErrUserDoesNotHaveWriteAccessToNamespace) Error() string {
	return fmt.Sprintf("You need to have write access to this namespace to do that [NamespaceID: %d, UserID: %d]", err.NamespaceID, err.UserID)
}

// ErrNamespaceNameCannotBeEmpty represents an error, where a namespace name is empty.
type ErrNamespaceNameCannotBeEmpty struct {
	NamespaceID int64
	UserID      int64
}

// IsErrNamespaceNameCannotBeEmpty checks if an error is a ErrNamespaceDoesNotExist.
func IsErrNamespaceNameCannotBeEmpty(err error) bool {
	_, ok := err.(ErrNamespaceNameCannotBeEmpty)
	return ok
}

func (err ErrNamespaceNameCannotBeEmpty) Error() string {
	return fmt.Sprintf("Namespace name cannot be emtpy [NamespaceID: %d, UserID: %d]", err.NamespaceID, err.UserID)
}

// ErrNamespaceOwnerCannotBeEmpty represents an error, where a namespace owner is empty.
type ErrNamespaceOwnerCannotBeEmpty struct {
	NamespaceID int64
	UserID      int64
}

// IsErrNamespaceOwnerCannotBeEmpty checks if an error is a ErrNamespaceDoesNotExist.
func IsErrNamespaceOwnerCannotBeEmpty(err error) bool {
	_, ok := err.(ErrNamespaceOwnerCannotBeEmpty)
	return ok
}

func (err ErrNamespaceOwnerCannotBeEmpty) Error() string {
	return fmt.Sprintf("Namespace owner cannot be emtpy [NamespaceID: %d, UserID: %d]", err.NamespaceID, err.UserID)
}

// ErrNeedToBeNamespaceAdmin represents an error, where the user is not the admin of that namespace (used i.e. when deleting a namespace)
type ErrNeedToBeNamespaceAdmin struct {
	NamespaceID int64
	UserID      int64
}

// IsErrNeedToBeNamespaceAdmin checks if an error is a ErrNamespaceDoesNotExist.
func IsErrNeedToBeNamespaceAdmin(err error) bool {
	_, ok := err.(ErrNeedToBeNamespaceAdmin)
	return ok
}

func (err ErrNeedToBeNamespaceAdmin) Error() string {
	return fmt.Sprintf("You need to be namespace owner to do that [NamespaceID: %d, UserID: %d]", err.NamespaceID, err.UserID)
}

// ErrNeedToHaveNamespaceReadAccess represents an error, where the user is not the owner of that namespace (used i.e. when deleting a namespace)
type ErrNeedToHaveNamespaceReadAccess struct {
	NamespaceID int64
	UserID      int64
}

// IsErrNeedToHaveNamespaceReadAccess checks if an error is a ErrNamespaceDoesNotExist.
func IsErrNeedToHaveNamespaceReadAccess(err error) bool {
	_, ok := err.(ErrNeedToHaveNamespaceReadAccess)
	return ok
}

func (err ErrNeedToHaveNamespaceReadAccess) Error() string {
	return fmt.Sprintf("You need to be namespace owner to do that [NamespaceID: %d, UserID: %d]", err.NamespaceID, err.UserID)
}

// ============
// Team errors
// ============

// ErrTeamNameCannotBeEmpty represents an error where a team name is empty.
type ErrTeamNameCannotBeEmpty struct {
	TeamID int64
}

// IsErrTeamNameCannotBeEmpty checks if an error is a ErrNamespaceDoesNotExist.
func IsErrTeamNameCannotBeEmpty(err error) bool {
	_, ok := err.(ErrTeamNameCannotBeEmpty)
	return ok
}

func (err ErrTeamNameCannotBeEmpty) Error() string {
	return fmt.Sprintf("Team name cannot be empty [Team ID: %d]", err.TeamID)
}

// ErrTeamDoesNotExist represents an error where a team does not exist
type ErrTeamDoesNotExist struct {
	TeamID int64
}

// IsErrTeamDoesNotExist checks if an error is ErrTeamDoesNotExist.
func IsErrTeamDoesNotExist(err error) bool {
	_, ok := err.(ErrTeamDoesNotExist)
	return ok
}

func (err ErrTeamDoesNotExist) Error() string {
	return fmt.Sprintf("Team does not exist [Team ID: %d]", err.TeamID)
}

// ErrInvalidTeamRight represents an error where a team right is invalid
type ErrInvalidTeamRight struct {
	Right TeamRight
}

// IsErrInvalidTeamRight checks if an error is ErrInvalidTeamRight.
func IsErrInvalidTeamRight(err error) bool {
	_, ok := err.(ErrInvalidTeamRight)
	return ok
}

func (err ErrInvalidTeamRight) Error() string {
	return fmt.Sprintf("The right is invalid [Right: %d]", err.Right)
}

// ErrTeamAlreadyHasAccess represents an error where a team already has access to a list/namespace
type ErrTeamAlreadyHasAccess struct {
	TeamID int64
	ID     int64
}

// IsErrTeamAlreadyHasAccess checks if an error is ErrTeamAlreadyHasAccess.
func IsErrTeamAlreadyHasAccess(err error) bool {
	_, ok := err.(ErrTeamAlreadyHasAccess)
	return ok
}

func (err ErrTeamAlreadyHasAccess) Error() string {
	return fmt.Sprintf("This team already has access. [Team ID: %d, ID: %d]", err.TeamID, err.ID)
}

// ErrUserIsMemberOfTeam represents an error where a user is already member of a team.
type ErrUserIsMemberOfTeam struct {
	TeamID int64
	UserID int64
}

// IsErrUserIsMemberOfTeam checks if an error is ErrUserIsMemberOfTeam.
func IsErrUserIsMemberOfTeam(err error) bool {
	_, ok := err.(ErrUserIsMemberOfTeam)
	return ok
}

func (err ErrUserIsMemberOfTeam) Error() string {
	return fmt.Sprintf("This user is already a member of that team. [Team ID: %d, User ID: %d]", err.TeamID, err.UserID)
}

// ErrCannotDeleteLastTeamMember represents an error where a user wants to delete the last member of a team (probably himself)
type ErrCannotDeleteLastTeamMember struct {
	TeamID int64
	UserID int64
}

// IsErrCannotDeleteLastTeamMember checks if an error is ErrCannotDeleteLastTeamMember.
func IsErrCannotDeleteLastTeamMember(err error) bool {
	_, ok := err.(ErrCannotDeleteLastTeamMember)
	return ok
}

func (err ErrCannotDeleteLastTeamMember) Error() string {
	return fmt.Sprintf("This user is already a member of that team. [Team ID: %d, User ID: %d]", err.TeamID, err.UserID)
}

// ====================
// User <-> List errors
// ====================

// ErrInvalidUserRight represents an error where a user right is invalid
type ErrInvalidUserRight struct {
	Right UserRight
}

// IsErrInvalidUserRight checks if an error is ErrInvalidUserRight.
func IsErrInvalidUserRight(err error) bool {
	_, ok := err.(ErrInvalidUserRight)
	return ok
}

func (err ErrInvalidUserRight) Error() string {
	return fmt.Sprintf("The right is invalid [Right: %d]", err.Right)
}

// ErrUserAlreadyHasAccess represents an error where a user already has access to a list/namespace
type ErrUserAlreadyHasAccess struct {
	UserID int64
	ListID int64
}

// IsErrUserAlreadyHasAccess checks if an error is ErrUserAlreadyHasAccess.
func IsErrUserAlreadyHasAccess(err error) bool {
	_, ok := err.(ErrUserAlreadyHasAccess)
	return ok
}

func (err ErrUserAlreadyHasAccess) Error() string {
	return fmt.Sprintf("This user already has access to that list. [User ID: %d, List ID: %d]", err.UserID, err.ListID)
}
