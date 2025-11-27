package kaptinlin_jsonschema_test

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/kaptinlin/jsonschema"

type Level1Struct struct {
	Name   string       `json:"name" jsonschema:"required"`
	Level2 Level2Struct `json:"level2" jsonschema:"required"`
}

type Level2Struct struct {
	Level3 Level3Struct `json:"level3" jsonschema:"required"`
}

type Level3Struct struct {
	Description string `json:"description" jsonschema:"required"`
}

func Test1(t *testing.T) {
	structTagOptions := jsonschema.DefaultStructTagOptions()
	schema, err := jsonschema.FromStructWithOptions[Level1Struct](structTagOptions)

	if assert.NoError(t, err) {
		{
			v := Level1Struct{
				Name: "Bob",
				Level2: Level2Struct{
					Level3: Level3Struct{
						Description: "This is a description",
					},
				},
			}
			assert.True(t, validateStruct(schema, v))
		}
		{
			v := Level1Struct{
				Level2: Level2Struct{
					Level3: Level3Struct{
						Description: "This is a description",
					},
				},
			}
			assert.False(t, validateStruct(schema, v))
		}
		{
			v := Level1Struct{
				Name: "Bob",
				Level2: Level2Struct{
					Level3: Level3Struct{},
				},
			}
			assert.False(t, validateStruct(schema, v))
		}
		{
			v := Level1Struct{
				Level2: Level2Struct{
					Level3: Level3Struct{},
				},
			}
			assert.False(t, validateStruct(schema, v))
		}
	}
}

func validateStruct(schema *jsonschema.Schema, v any) bool {
	result := schema.ValidateStruct(v)
	return result.IsValid()
}
