package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
)

type LogURL struct {
	gorm.Model
	BlockNumber       uint64    `gorm:"column:block_number;index;" json:"blockNumber"`
	BlockHash         string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash            string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex           uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex          uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed           bool      `gorm:"column:removed;" json:"removed"`
	Date              time.Time `gorm:"column:date;" json:"date"`
	QueryId           string    `gorm:"column:query_id" json:"queryId"`
	Timeout           string    `gorm:"column:time_out" json:"timeOut"`
	DataSource        string    `gorm:"column:data_source" json:"dataSource"`
	Selector          string    `gorm:"column:selector" json:"selector"`
	Randomness        string    `gorm:"column:randomness" json:"randomness"`
	DispatchedGroupId string    `gorm:"column:dispatched_groupid" json:"dispatchedGroupId"`
}

// Set User's table name to be `logurl`
func (LogURL) TableName() string {
	return "logurl"
}

type LogRequestUserRandom struct {
	gorm.Model
	BlockNumber          uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash            string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash               string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex              uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex             uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed              bool      `gorm:"column:removed;" json:"removed"`
	Date                 time.Time `gorm:"column:date;" json:"date"`
	RequestId            string    `gorm:"column:request_id" json:"requestId"`
	LastSystemRandomness string    `gorm:"column:last_system_randomness" json:"lastSystemRandomness"`
	UserSeed             string    `gorm:"column:user_seed" json:"userSeed"`
	DispatchedGroupId    string    `gorm:"column:dispatched_groupid" json:"dispatchedGroupId"`
}

// Set User's table name to be `profiles`
func (LogRequestUserRandom) TableName() string {
	return "logrequestuserrandom"
}

type LogNonSupportedType struct {
	gorm.Model
	BlockNumber     uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash       string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash          string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex         uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex        uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed         bool      `gorm:"column:removed;" json:"removed"`
	Date            time.Time `gorm:"column:date;" json:"date"`
	InvalidSelector string    `gorm:"column:invalid_selector" json:"invalidSelector"`
}

// Set User's table name to be `profiles`
func (LogNonSupportedType) TableName() string {
	return "lognonsupportedtype"
}

type LogNonContractCall struct {
	gorm.Model
	BlockNumber uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool      `gorm:"column:removed;" json:"removed"`
	Date        time.Time `gorm:"column:date;" json:"date"`
	From        string    `gorm:"column:from" json:"from"`
}

// Set User's table name to be `profiles`
func (LogNonContractCall) TableName() string {
	return "lognoncontractcall"
}

type LogCallbackTriggeredFor struct {
	gorm.Model
	BlockNumber  uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash    string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash       string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex      uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex     uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed      bool      `gorm:"column:removed;" json:"removed"`
	Date         time.Time `gorm:"column:date;" json:"date"`
	CallbackAddr string    `gorm:"column:call_back_addr" json:"callbackAddr"`
}

// Set User's table name to be `profiles`
func (LogCallbackTriggeredFor) TableName() string {
	return "logcallbacktriggeredfor"
}

type LogRequestFromNonExistentUC struct {
	gorm.Model
	BlockNumber uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool      `gorm:"column:removed;" json:"removed"`
	Date        time.Time `gorm:"column:date;" json:"date"`
}

// Set User's table name to be `profiles`
func (LogRequestFromNonExistentUC) TableName() string {
	return "logrequestfromnonexistentuc"
}

type LogUpdateRandom struct {
	gorm.Model
	BlockNumber       uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash         string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash            string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex           uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex          uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed           bool      `gorm:"column:removed;" json:"removed"`
	Date              time.Time `gorm:"column:date;" json:"date"`
	LastRandomness    string    `gorm:"column:last_randomness" json:"lastRandomness"`
	DispatchedGroupId string    `gorm:"column:dispatched_groupid" json:"dispatchedGroupId"`
}

// Set User's table name to be `profiles`
func (LogUpdateRandom) TableName() string {
	return "logupdaterandom"
}

type LogValidationResult struct {
	gorm.Model
	BlockNumber uint64         `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string         `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string         `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint           `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint           `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool           `gorm:"column:removed;" json:"removed"`
	Date        time.Time      `gorm:"column:date;" json:"date"`
	TrafficType uint8          `gorm:"column:traffic_type" json:"trafficType"`
	TrafficId   string         `gorm:"column:traffic_id" json:"trafficId"`
	Message     []byte         `gorm:"column:message;type:bytea" json:"message"`
	Signature   pq.StringArray `gorm:"column:signature;type:varchar(100)[]" json:"signature"`
	PubKey      pq.StringArray `gorm:"column:pub_key;type:varchar(100)[]" json:"pubKey"`
	Pass        bool           `gorm:"column:pass;" json:"pass"`
}

// Set User's table name to be `profiles`
func (LogValidationResult) TableName() string {
	return "logvalidationresult"
}

type LogInsufficientPendingNode struct {
	gorm.Model
	BlockNumber     uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash       string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash          string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex         uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex        uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed         bool      `gorm:"column:removed;" json:"removed"`
	Date            time.Time `gorm:"column:date;" json:"date"`
	NumPendingNodes uint64    `gorm:"column:num_pending_nodes" json:"numPendingNodes"`
}

// Set User's table name to be `profiles`
func (LogInsufficientPendingNode) TableName() string {
	return "loginsufficientpendingnode"
}

type LogInsufficientWorkingGroup struct {
	gorm.Model
	BlockNumber      uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash        string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash           string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex          uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex         uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed          bool      `gorm:"column:removed;" json:"removed"`
	Date             time.Time `gorm:"column:date;" json:"date"`
	NumWorkingGroups uint64    `gorm:"column:num_working_groups" json:"numWorkingGroups"`
	NumPendingGroups uint64    `gorm:"column:num_pending_groups" json:"numPendingGroups"`
}

// Set User's table name to be `profiles`
func (LogInsufficientWorkingGroup) TableName() string {
	return "loginsufficientworkinggroup"
}

type LogGrouping struct {
	gorm.Model
	BlockNumber uint64         `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string         `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string         `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint           `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint           `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool           `gorm:"column:removed;" json:"removed"`
	Date        time.Time      `gorm:"column:date;" json:"date"`
	GroupId     string         `gorm:"column:group_id" json:"groupId"`
	NodeId      pq.StringArray `gorm:"column:node_id;type:varchar(100)[]" json:"nodeId"`
}

// Set User's table name to be `profiles`
func (LogGrouping) TableName() string {
	return "loggrouping"
}

type LogPublicKeyAccepted struct {
	gorm.Model
	BlockNumber      uint64         `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash        string         `gorm:"column:block_hash" json:"blockHash"`
	TxHash           string         `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex          uint           `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex         uint           `gorm:"column:log_index;" json:"logIndex"`
	Removed          bool           `gorm:"column:removed;" json:"removed"`
	Date             time.Time      `gorm:"column:date;" json:"date"`
	GroupId          string         `gorm:"column:group_id" json:"groupId"`
	PubKey           pq.StringArray `gorm:"column:pub_key;type:varchar(100)[]" json:"pubKey"`
	NumWorkingGroups uint64         `gorm:"column:num_working_groups;" json:"numWorkingGroups"`
}

// Set User's table name to be `profiles`
func (LogPublicKeyAccepted) TableName() string {
	return "logpublickeyaccepted"
}

type LogPublicKeySuggested struct {
	gorm.Model
	BlockNumber uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool      `gorm:"column:removed;" json:"removed"`
	Date        time.Time `gorm:"column:date;" json:"date"`
	GroupId     string    `gorm:"column:group_id" json:"groupId"`
	PubKeyCount uint64    `gorm:"column:pub_key_count" json:"pubKeyCount"`
}

// Set User's table name to be `profiles`
func (LogPublicKeySuggested) TableName() string {
	return "logpublickeysuggested"
}

type LogGroupDissolve struct {
	gorm.Model
	BlockNumber uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool      `gorm:"column:removed;" json:"removed"`
	Date        time.Time `gorm:"column:date;" json:"date"`
	GroupId     string    `gorm:"column:group_id" json:"groupId"`
}

// Set User's table name to be `profiles`
func (LogGroupDissolve) TableName() string {
	return "loggroupdissolve"
}

type LogRegisteredNewPendingNode struct {
	gorm.Model
	BlockNumber uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool      `gorm:"column:removed;" json:"removed"`
	Date        time.Time `gorm:"column:date;" json:"date"`
	Node        string    `gorm:"column:node" json:"node"`
}

// Set User's table name to be `profiles`
func (LogRegisteredNewPendingNode) TableName() string {
	return "logregisterednewpendingnode"
}

type LogGroupingInitiated struct {
	gorm.Model
	BlockNumber       uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash         string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash            string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex           uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex          uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed           bool      `gorm:"column:removed;" json:"removed"`
	Date              time.Time `gorm:"column:date;" json:"date"`
	PendingNodePool   uint64    `gorm:"column:pending_node_pool" json:"pendingNodePool"`
	Groupsize         uint64    `gorm:"column:group_size" json:"groupSize"`
	Groupingthreshold uint64    `gorm:"column:grouping_threshold" json:"groupingThreshold"`
}

// Set User's table name to be `profiles`
func (LogGroupingInitiated) TableName() string {
	return "loggroupinginitiated"
}

type LogNoPendingGroup struct {
	gorm.Model
	BlockNumber uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool      `gorm:"column:removed;" json:"removed"`
	Date        time.Time `gorm:"column:date;" json:"date"`
	GroupId     string    `gorm:"column:group_id" json:"groupId"`
}

// Set User's table name to be `profiles`
func (LogNoPendingGroup) TableName() string {
	return "lognopendinggroup"
}

type LogPendingGroupRemoved struct {
	gorm.Model
	BlockNumber uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool      `gorm:"column:removed;" json:"removed"`
	Date        time.Time `gorm:"column:date;" json:"date"`
	GroupId     string    `gorm:"column:group_id" json:"groupId"`
}

// Set User's table name to be `profiles`
func (LogPendingGroupRemoved) TableName() string {
	return "logpendinggroupremoved"
}

type LogError struct {
	gorm.Model
	BlockNumber uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool      `gorm:"column:removed;" json:"removed"`
	Date        time.Time `gorm:"column:date;" json:"date"`
	Err         string    `gorm:"column:err" json:"err"`
}

func (LogError) TableName() string {
	return "logerror"
}

type UpdateGroupToPick struct {
	gorm.Model
	BlockNumber uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool      `gorm:"column:removed;" json:"removed"`
	Date        time.Time `gorm:"column:date;" json:"date"`
	OldNum      uint64    `gorm:"column:old_num" json:"oldNum"`
	NewNum      uint64    `gorm:"column:new_num" json:"newNum"`
}

func (UpdateGroupToPick) TableName() string {
	return "updategrouptopick"
}

type UpdateGroupSize struct {
	gorm.Model
	BlockNumber uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool      `gorm:"column:removed;" json:"removed"`
	Date        time.Time `gorm:"column:date;" json:"date"`
	OldSize     uint64    `gorm:"column:old_size" json:"oldSize"`
	NewSize     uint64    `gorm:"column:new_size" json:"newSize"`
}

func (UpdateGroupSize) TableName() string {
	return "updategroupsize"
}

type UpdateGroupingThreshold struct {
	gorm.Model
	BlockNumber  uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash    string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash       string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex      uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex     uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed      bool      `gorm:"column:removed;" json:"removed"`
	Date         time.Time `gorm:"column:date;" json:"date"`
	OldThreshold uint64    `gorm:"column:old_threshold" json:"oldThreshold"`
	NewThreshold uint64    `gorm:"column:new_threshold" json:"newThreshold"`
}

func (UpdateGroupingThreshold) TableName() string {
	return "updategroupingthreshold"
}

type UpdateGroupMaturityPeriod struct {
	gorm.Model
	BlockNumber uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool      `gorm:"column:removed;" json:"removed"`
	Date        time.Time `gorm:"column:date;" json:"date"`
	OldPeriod   uint64    `gorm:"column:old_period" json:"oldPeriod"`
	NewPeriod   uint64    `gorm:"column:new_period" json:"newPeriod"`
}

func (UpdateGroupMaturityPeriod) TableName() string {
	return "updategroupmaturityperiod"
}

type UpdateBootstrapCommitDuration struct {
	gorm.Model
	BlockNumber uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool      `gorm:"column:removed;" json:"removed"`
	Date        time.Time `gorm:"column:date;" json:"date"`
	OldDuration uint64    `gorm:"column:old_duration" json:"oldDuration"`
	NewDuration uint64    `gorm:"column:new_duration" json:"newDuration"`
}

func (UpdateBootstrapCommitDuration) TableName() string {
	return "updatebootstrapcommitduration"
}

type UpdateBootstrapRevealDuration struct {
	gorm.Model
	BlockNumber uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool      `gorm:"column:removed;" json:"removed"`
	Date        time.Time `gorm:"column:date;" json:"date"`
	OldDuration uint64    `gorm:"column:old_duration" json:"oldDuration"`
	NewDuration uint64    `gorm:"column:new_duration" json:"newDuration"`
}

func (UpdateBootstrapRevealDuration) TableName() string {
	return "updatebootstraprevealduration"
}

type UpdatebootstrapStartThreshold struct {
	gorm.Model
	BlockNumber  uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash    string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash       string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex      uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex     uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed      bool      `gorm:"column:removed;" json:"removed"`
	Date         time.Time `gorm:"column:date;" json:"date"`
	OldThreshold uint64    `gorm:"column:old_threshold" json:"oldThreshold"`
	NewThreshold uint64    `gorm:"column:new_threshold" json:"newThreshold"`
}

func (UpdatebootstrapStartThreshold) TableName() string {
	return "updatebootstrapstartthreshold"
}

type UpdatePendingGroupMaxLife struct {
	gorm.Model
	BlockNumber   uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash     string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash        string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex       uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex      uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed       bool      `gorm:"column:removed;" json:"removed"`
	Date          time.Time `gorm:"column:date;" json:"date"`
	OldLifeBlocks uint64    `gorm:"column:old_life_blocks" json:"oldLifeBlocks"`
	NewLifeBlocks uint64    `gorm:"column:new_life_blocks" json:"newLifeBlock"`
}

func (UpdatePendingGroupMaxLife) TableName() string {
	return "updatependinggroupmaxlife"
}

type GuardianReward struct {
	gorm.Model
	BlockNumber uint64    `gorm:"column:block_number;" json:"blockNumber"`
	BlockHash   string    `gorm:"column:block_hash" json:"blockHash"`
	TxHash      string    `gorm:"column:transaction_hash" json:"transactionHash"`
	TxIndex     uint      `gorm:"column:transaction_index" json:"transactionIndex"`
	LogIndex    uint      `gorm:"column:log_index;" json:"logIndex"`
	Removed     bool      `gorm:"column:removed;" json:"removed"`
	Date        time.Time `gorm:"column:date;" json:"date"`
	BlkNum      uint64    `gorm:"column:blk_num" json:"blkNum"`
	Guardian    string    `gorm:"column:guardian" json:"guardian"`
}

func (GuardianReward) TableName() string {
	return "guardianreward"
}

func Connect() *gorm.DB {
	postgres_url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	var db *gorm.DB
	db, err := gorm.Open("postgres", postgres_url)
	if err != nil {
		log.Fatal(err)
	}
	//&LogPublicKeySuggested{},

	// AutoMigrate will ONLY create tables, missing columns and missing indexes,
	// and WON’T change existing column’s type or delete unused columns to protect your data.
	db.AutoMigrate(&LogURL{}, &LogRequestUserRandom{}, &LogNonSupportedType{}, &LogNonContractCall{}, &LogCallbackTriggeredFor{}, &LogRequestFromNonExistentUC{}, &LogUpdateRandom{}, &LogValidationResult{}, &LogInsufficientPendingNode{}, &LogInsufficientWorkingGroup{}, &LogGrouping{}, &LogPublicKeyAccepted{}, &LogPublicKeySuggested{}, &LogGroupDissolve{}, &LogRegisteredNewPendingNode{}, &LogGroupingInitiated{}, &LogNoPendingGroup{}, &LogPendingGroupRemoved{}, &LogError{}, &UpdateGroupToPick{}, &UpdateGroupSize{}, &UpdateGroupingThreshold{}, &UpdateGroupMaturityPeriod{}, &UpdateBootstrapCommitDuration{}, &UpdateBootstrapRevealDuration{}, &UpdatebootstrapStartThreshold{}, &UpdatePendingGroupMaxLife{}, &GuardianReward{})
	//db.AutoMigrate(&LogURL{}, &LogRequestUserRandom{}, &LogNonSupportedType{}, &LogNonContractCall{}, &LogCallbackTriggeredFor{}, &LogRequestFromNonExistentUC{}, &LogUpdateRandom{}, &LogValidationResult{}, &LogInsufficientPendingNode{}, &LogInsufficientWorkingGroup{}, &LogGrouping{}, &LogPublicKeyAccepted{}, &LogGroupDissolve{}, &LogRegisteredNewPendingNode{}, &LogGroupingInitiated{}, &LogNoPendingGroup{}, &LogPendingGroupRemoved{}, &LogError{}, &UpdateGroupToPick{}, &UpdateGroupSize{}, &UpdateGroupingThreshold{}, &UpdateGroupMaturityPeriod{}, &UpdateBootstrapCommitDuration{}, &UpdateBootstrapRevealDuration{}, &UpdatebootstrapStartThreshold{}, &UpdatePendingGroupMaxLife{}, &GuardianReward{})

	// DB.LogMode(true)
	log.Info("DB Connected")
	return db
}
