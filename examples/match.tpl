
Environment variables starting with PROD_:
{{ range $key, $value := match "PROD_"  }}{{ $key }}="{{ $value }}"
{{ end -}}
