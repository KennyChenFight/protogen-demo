package module

import (
	"text/template"

	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	pgs "github.com/lyft/protoc-gen-star"
)

type JSONifyModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

func NewJSONifyModule() *JSONifyModule { return &JSONifyModule{ModuleBase: &pgs.ModuleBase{}} }

func (p *JSONifyModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("jsonify").Funcs(map[string]interface{}{
		"package": p.ctx.PackageName,
		"name":    p.ctx.Name,
	})

	p.tpl = template.Must(tpl.Parse(jsonifyTpl))
}

func (p *JSONifyModule) Name() string { return "jsonify" }

func (p *JSONifyModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {

	for _, t := range targets {
		p.generate(t)
	}

	return p.Artifacts()
}

func (p *JSONifyModule) generate(f pgs.File) {
	if len(f.Messages()) == 0 {
		return
	}

	name := p.ctx.OutputPath(f).SetBase("pb").SetExt("_demo2.json.go")
	p.AddGeneratorTemplateFile(name.String(), p.tpl, f)
}

const jsonifyTpl = `package {{ package . }}

import "google.golang.org/protobuf/encoding/protojson"

{{ range .AllMessages }}

func (m *{{ name . }}) MarshalJSON() ([]byte, error) {
	b, err := protojson.Marshal(m)
	if err != nil {
		return nil, err	
	}
	return b, nil
}

var _ json.Marshaler = (*{{ name . }})(nil)

func (m *{{ name . }}) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, m)
}

var _ json.Unmarshaler = (*{{ name . }})(nil)

{{ end }}
`
