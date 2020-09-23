## List of types
{{range $t, $type:= .installedTypes}}
### <a name="{{$t}}"></a>{{$t}}
> {{$type.typespec.description | replace "\n" "\n> " | noescape}}

{{range $fieldname, $field:= $type.fields}}
**{{$fieldname}}**
<small>[{{$field.type}}](#{{$field.type}})</small>
<br>{{$field.description | replace "\n" "<br> " | noescape}}
{{end}}
{{end}}

{{range $t, $type:= .installedTypes}}
{{$t}}
{{end}}