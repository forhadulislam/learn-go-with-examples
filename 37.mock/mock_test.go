package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestMockAPIMockRecorder_CreateItem(t *testing.T) {
	tests := []struct {
		name         string
		inputData    map[string]any
		expectedData string
	}{
		{
			name:         "test1",
			inputData:    map[string]any{"name": "test1", "id": 1},
			expectedData: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := NewMockAPI(gomock.NewController(t))
			mr.EXPECT().CreateItem(tt.inputData).Return(tt.expectedData, nil)
			got, err := mr.CreateItem(tt.inputData)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedData, got)
		})
	}
}

func TestMockAPIMockRecorder_DeleteItem(t *testing.T) {
	tests := []struct {
		name          string
		deleteID      string
		expectedError error
	}{
		{
			name:     "test1",
			deleteID: "1",
		},
		{
			name:          "error",
			deleteID:      "2",
			expectedError: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := NewMockAPI(gomock.NewController(t))
			mr.EXPECT().DeleteItem(tt.deleteID).Return(tt.expectedError)
			err := mr.DeleteItem(tt.deleteID)

			if tt.expectedError != nil {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestMockAPIMockRecorder_GetItem(t *testing.T) {
	tests := []struct {
		name          string
		inputID       string
		outputData    map[string]any
		expectedError error
	}{
		{
			name:       "test1",
			inputID:    "1",
			outputData: map[string]any{"name": "test1", "id": 1},
		},
		{
			name:          "test1",
			inputID:       "2",
			outputData:    map[string]any{},
			expectedError: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := NewMockAPI(gomock.NewController(t))
			mr.EXPECT().GetItem(tt.inputID).Return(tt.outputData, tt.expectedError)
			got, err := mr.GetItem(tt.inputID)
			if tt.expectedError != nil {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.outputData, got)
		})
	}
}
