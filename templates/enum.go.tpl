{{/*************************************************************************/}}
{{/* Variables */}}
{{/*************************************************************************/}}

{{ $en := .Enum.Name }}
{{ $m := $en | singular }}
{{ $mv := $en | singular | lower }}

{{/*************************************************************************/}}
{{/* Our Package */}}
{{/*************************************************************************/}}

package enum

{{/*************************************************************************/}}
{{/* Pogo marker */}}
{{/*************************************************************************/}}

// GENERATED BY POGO. DO NOT EDIT.

{{/*************************************************************************/}}
{{/* Our enum type */}}
{{/*************************************************************************/}}

// {{ $m }} is the `{{ $en }}` enum type from `{{ .Schema.Name }}`.
type {{ $m }} string

{{/*************************************************************************/}}
{{/* Our enum type */}}
{{/*************************************************************************/}}

const (
  {{ range .Enum.Values }}
	// {{ $m }}{{ .Label | capitalize }} is the '{{ .Label }}' {{ $m }}.
	{{ $m }}{{ .Label | capitalize }} = {{ $m }}("{{ .Label }}")
  {{ end }}
)

{{/*************************************************************************/}}
{{/* Satisfy the sql/driver.Valuer interface */}}
{{/*************************************************************************/}}

// Value satisfies the sql/driver.Valuer interface for {{ $m }}.
func ({{ $mv }} {{ $m }}) Value() (driver.Value, error) {
	return string({{ $mv }}), nil
}