package main

import (
	"fmt"
	"strings"

	"github.com/conductorone/baton-sdk/pkg/field"
	"github.com/spf13/viper"
)

var (
	UsernameField = field.StringField(
		"username",
		field.WithDescription("The Avalara username used to connect to the Avalara API"),
		field.WithRequired(true),
	)
	PasswordField = field.StringField(
		"password",
		field.WithDescription("The Avalara password used to connect to the Avalara API"),
		field.WithRequired(true),
	)
	EnvironmentField = field.StringField(
		"environment",
		field.WithDescription("The Avalara environment to connect to (production or sandbox)"),
		field.WithDefaultValue("production"),
	)

	ConfigurationFields = []field.SchemaField{
		UsernameField,
		PasswordField,
		EnvironmentField,
	}

	FieldRelationships = []field.SchemaFieldRelationship{
		field.FieldsRequiredTogether(
			UsernameField,
			PasswordField,
		),
	}
)

// ValidateConfig checks the configuration values and returns an error if they are invalid
func ValidateConfig(v *viper.Viper) error {
	username := v.GetString(UsernameField.FieldName)
	password := v.GetString(PasswordField.FieldName)
	environment := v.GetString(EnvironmentField.FieldName)

	if username == "" || password == "" {
		return fmt.Errorf("both username and password are required")
	}

	environment = strings.ToLower(environment)
	if environment != "production" && environment != "sandbox" {
		return fmt.Errorf("invalid environment: must be either 'production' or 'sandbox'")
	}

	return nil
}
