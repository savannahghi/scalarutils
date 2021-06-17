package go_utils_test

import (
	"bytes"
	"os"
	"strconv"
	"testing"

	base "github.com/savannahghi/go_utils"
	"github.com/stretchr/testify/assert"
)

func TestGender_String(t *testing.T) {
	tests := []struct {
		name string
		e    base.Gender
		want string
	}{
		{
			name: "male",
			e:    base.GenderMale,
			want: "male",
		},
		{
			name: "female",
			e:    base.GenderFemale,
			want: "female",
		},
		{
			name: "unknown",
			e:    base.GenderUnknown,
			want: "unknown",
		},
		{
			name: "other",
			e:    base.GenderOther,
			want: "other",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Gender.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGender_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    base.Gender
		want bool
	}{
		{
			name: "valid male",
			e:    base.GenderMale,
			want: true,
		},
		{
			name: "invalid gender",
			e:    base.Gender("this is not a real gender"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("Gender.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGender_UnmarshalGQL(t *testing.T) {
	female := base.GenderFemale
	invalid := base.Gender("")
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *base.Gender
		args    args
		wantErr bool
	}{
		{
			name: "valid female gender",
			e:    &female,
			args: args{
				v: "female",
			},
			wantErr: false,
		},
		{
			name: "invalid gender",
			e:    &invalid,
			args: args{
				v: "this is not a real gender",
			},
			wantErr: true,
		},
		{
			name: "non string gender",
			e:    &invalid,
			args: args{
				v: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.UnmarshalGQL(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Gender.UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGender_MarshalGQL(t *testing.T) {
	tests := []struct {
		name  string
		e     base.Gender
		wantW string
	}{
		{
			name:  "valid unknown gender enum",
			e:     base.GenderUnknown,
			wantW: strconv.Quote("unknown"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.e.MarshalGQL(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Gender.MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestContentType_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    base.ContentType
		want bool
	}{
		{
			name: "good case",
			e:    base.ContentTypeJpg,
			want: true,
		},
		{
			name: "bad case",
			e:    base.ContentType("not a real content type"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("ContentType.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContentType_String(t *testing.T) {
	tests := []struct {
		name string
		e    base.ContentType
		want string
	}{
		{
			name: "default case",
			e:    base.ContentTypePdf,
			want: "PDF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("ContentType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContentType_UnmarshalGQL(t *testing.T) {
	var sc base.ContentType
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *base.ContentType
		args    args
		wantErr bool
	}{
		{
			name: "valid unmarshal",
			e:    &sc,
			args: args{
				v: "PDF",
			},
			wantErr: false,
		},
		{
			name: "invalid unmarshal",
			e:    &sc,
			args: args{
				v: "this is not a valid scalar value",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.UnmarshalGQL(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ContentType.UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContentType_MarshalGQL(t *testing.T) {
	tests := []struct {
		name  string
		e     base.ContentType
		wantW string
	}{
		{
			name:  "default case",
			e:     base.ContentTypePdf,
			wantW: strconv.Quote("PDF"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.e.MarshalGQL(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("ContentType.MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestContentType(t *testing.T) {
	type expects struct {
		isValid      bool
		canUnmarshal bool
	}

	cases := []struct {
		name        string
		args        base.ContentType
		convert     interface{}
		expectation expects
	}{
		{
			name:    "invalid_string",
			args:    "testcontent",
			convert: "testcontent",
			expectation: expects{
				isValid:      false,
				canUnmarshal: false,
			},
		},
		{
			name:    "invalid_int_convert",
			args:    "testaddres",
			convert: 101,
			expectation: expects{
				isValid:      false,
				canUnmarshal: false,
			},
		},
		{
			name:    "valid",
			args:    base.ContentTypePng,
			convert: base.ContentTypePng,
			expectation: expects{
				isValid:      true,
				canUnmarshal: true,
			},
		},
		{
			name:    "valid_no_convert",
			args:    base.ContentTypePdf,
			convert: "testaddress",
			expectation: expects{
				isValid:      true,
				canUnmarshal: true,
			},
		},
		{
			name:    "valid_can_convert",
			args:    base.ContentTypePdf,
			convert: base.ContentTypePdf,
			expectation: expects{
				isValid:      true,
				canUnmarshal: true,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation.isValid, tt.args.IsValid())
			assert.NotEmpty(t, tt.args.String())
			err := tt.args.UnmarshalGQL(tt.convert)
			assert.NotNil(t, err)
			tt.args.MarshalGQL(os.Stdout)

		})
	}

}
