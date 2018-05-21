package config

import (
	"time"

	"github.com/spf13/viper"
)

var (
	// IsPropertySet checks if some property exists in config.toml file.
	IsPropertySet = isPropertySet
	// GetHTTPReadTimeout returns timeout duration for HTTP read requests.
	GetHTTPReadTimeout = getHTTPReadTimeout
	// GetHTTPWriteTimeout returns timeout duration for HTTP write requests.
	GetHTTPWriteTimeout = getHTTPWriteTimeout
	// GetSSLCertificatePath returns path in which certificate .crt file can be found.
	GetSSLCertificatePath = getSSLCertificatePath
	// GetSSLKeystorePath returns path in which certificate keystore file can be found.
	GetSSLKeystorePath = getSSLKeystorePath
	// GetHTTPServerAddress returns address on which server will serve requests if it is started without transport security (http).
	GetHTTPServerAddress = getHTTPServerAddress
	// GetHTTPServerAddressSecure returns address on which server will serve requests if it is started with transport security (https).
	GetHTTPServerAddressSecure = getHTTPServerAddressSecure

	// GetDatabaseConnectionString returns string which contains all the info for connection to the database.
	GetDatabaseConnectionString = getDatabaseConnectionString
	// GetDatabaseMaxIdleConnections returns connection variable max_idle_connections to be set.
	GetDatabaseMaxIdleConnections = getDatabaseMaxIdleConnections
	// GetDatabaseMaxOpenConnections returns connection variable max_open_connections to be set.
	GetDatabaseMaxOpenConnections = getDatabaseMaxOpenConnections
	// GetDatabaseConnectionMaxLifetime returns connection variable max_lifetime to be set.
	GetDatabaseConnectionMaxLifetime = getDatabaseConnectionMaxLifetime

	GetQuestionStartString   = getQuestionStartString
	GetQuestionEndString     = getQuestionEndString
	GetQuestionTypeString    = getQuestionTypeString
	GetQuestionTextString    = getQuestionTextString
	GetQuestionAnswersString = getQuestionAnswersString

	GetQuestionTypeNamesFreeText      = getQuestionTypeNamesFreeText
	GetQuestionTypeNamesFreeNumerical = getQuestionTypeNamesFreeNumerical
	GetQuestionTypeNamesRadioGroup    = getQuestionTypeNamesRadioGroup
	GetQuestionTypeNamesCheckbox      = getQuestionTypeNamesCheckbox
)

func isPropertySet(s string) bool {
	return viper.IsSet(s)
}

func getHTTPReadTimeout() time.Duration {
	return viper.GetDuration("http.http_read_timeout")
}

func getHTTPWriteTimeout() time.Duration {
	return viper.GetDuration("http.http_write_timeout")
}

func getSSLCertificatePath() string {
	return viper.GetString("http.ssl_certificate_path")
}

func getSSLKeystorePath() string {
	return viper.GetString("http.ssl_keystore_path")
}

func getHTTPServerAddress() string {
	return viper.GetString("http.http_server_address")
}

func getHTTPServerAddressSecure() string {
	return viper.GetString("http.http_server_address_secure")
}

func getDatabaseConnectionString() string {
	return viper.GetString("db.connection")
}

func getDatabaseMaxIdleConnections() int {
	return viper.GetInt("db.max_idle_connections")
}

func getDatabaseMaxOpenConnections() int {
	return viper.GetInt("db.max_open_connections")
}

func getDatabaseConnectionMaxLifetime() time.Duration {
	return viper.GetDuration("db.max_lifetime")
}

func getQuestionStartString() string {
	return viper.GetString("questions.start")
}

func getQuestionEndString() string {
	return viper.GetString("questions.end")
}

func getQuestionTypeString() string {
	return viper.GetString("questions.type")
}

func getQuestionTextString() string {
	return viper.GetString("questions.text")
}

func getQuestionAnswersString() string {
	return viper.GetString("questions.answers")
}

func getQuestionTypeNamesFreeText() string {
	return viper.GetString("question_type_names.free_text")
}

func getQuestionTypeNamesFreeNumerical() string {
	return viper.GetString("question_type_names.free_numerical")
}

func getQuestionTypeNamesRadioGroup() string {
	return viper.GetString("question_type_names.radio_group")
}

func getQuestionTypeNamesCheckbox() string {
	return viper.GetString("question_type_names.checkbox")
}
