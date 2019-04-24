// This file is automatically generated by qtc from "integration.qtpl".
// See https://github.com/valyala/quicktemplate for details.

// This is a template for integration test.
// It should contains all the quicktemplate stuff.
//

//line testdata/templates/integration.qtpl:4
package templates

//line testdata/templates/integration.qtpl:4
import "fmt"

//line testdata/templates/integration.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line testdata/templates/integration.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line testdata/templates/integration.qtpl:6
func StreamIntegration(qw422016 *qt422016.Writer) {
//line testdata/templates/integration.qtpl:6
	qw422016.N().S(`
	Output tags`)
//line testdata/templates/integration.qtpl:6
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:6
	qw422016.N().S(` verification.

	`)
//line testdata/templates/integration.qtpl:10
	p := &integrationPage{
		S: "foobar",
	}

//line testdata/templates/integration.qtpl:13
	qw422016.N().S(`
	Embedded func template:
		plain: `)
//line testdata/templates/integration.qtpl:15
	streamembeddedFunc(qw422016, p)
//line testdata/templates/integration.qtpl:15
	qw422016.N().S(`
		html-escaped: `)
//line testdata/templates/integration.qtpl:16
	{
//line testdata/templates/integration.qtpl:16
		qb422016 := qt422016.AcquireByteBuffer()
//line testdata/templates/integration.qtpl:16
		writeembeddedFunc(qb422016, p)
//line testdata/templates/integration.qtpl:16
		qw422016.E().Z(qb422016.B)
//line testdata/templates/integration.qtpl:16
		qt422016.ReleaseByteBuffer(qb422016)
//line testdata/templates/integration.qtpl:16
	}
//line testdata/templates/integration.qtpl:16
	qw422016.N().S(`
		url-escaped: `)
//line testdata/templates/integration.qtpl:17
	{
//line testdata/templates/integration.qtpl:17
		qb422016 := qt422016.AcquireByteBuffer()
//line testdata/templates/integration.qtpl:17
		writeembeddedFunc(qb422016, p)
//line testdata/templates/integration.qtpl:17
		qw422016.N().UZ(qb422016.B)
//line testdata/templates/integration.qtpl:17
		qt422016.ReleaseByteBuffer(qb422016)
//line testdata/templates/integration.qtpl:17
	}
//line testdata/templates/integration.qtpl:17
	qw422016.N().S(`
		quoted json string: `)
//line testdata/templates/integration.qtpl:18
	{
//line testdata/templates/integration.qtpl:18
		qb422016 := qt422016.AcquireByteBuffer()
//line testdata/templates/integration.qtpl:18
		writeembeddedFunc(qb422016, p)
//line testdata/templates/integration.qtpl:18
		qw422016.N().QZ(qb422016.B)
//line testdata/templates/integration.qtpl:18
		qt422016.ReleaseByteBuffer(qb422016)
//line testdata/templates/integration.qtpl:18
	}
//line testdata/templates/integration.qtpl:18
	qw422016.N().S(`
		unquoted json string: `)
//line testdata/templates/integration.qtpl:19
	{
//line testdata/templates/integration.qtpl:19
		qb422016 := qt422016.AcquireByteBuffer()
//line testdata/templates/integration.qtpl:19
		writeembeddedFunc(qb422016, p)
//line testdata/templates/integration.qtpl:19
		qw422016.N().JZ(qb422016.B)
//line testdata/templates/integration.qtpl:19
		qt422016.ReleaseByteBuffer(qb422016)
//line testdata/templates/integration.qtpl:19
	}
//line testdata/templates/integration.qtpl:19
	qw422016.N().S(`
		html-escaped url-escaped: `)
//line testdata/templates/integration.qtpl:20
	{
//line testdata/templates/integration.qtpl:20
		qb422016 := qt422016.AcquireByteBuffer()
//line testdata/templates/integration.qtpl:20
		writeembeddedFunc(qb422016, p)
//line testdata/templates/integration.qtpl:20
		qw422016.N().UZ(qb422016.B)
//line testdata/templates/integration.qtpl:20
		qt422016.ReleaseByteBuffer(qb422016)
//line testdata/templates/integration.qtpl:20
	}
//line testdata/templates/integration.qtpl:20
	qw422016.N().S(`
		html-escaped quoted json string: `)
//line testdata/templates/integration.qtpl:21
	{
//line testdata/templates/integration.qtpl:21
		qb422016 := qt422016.AcquireByteBuffer()
//line testdata/templates/integration.qtpl:21
		writeembeddedFunc(qb422016, p)
//line testdata/templates/integration.qtpl:21
		qw422016.E().QZ(qb422016.B)
//line testdata/templates/integration.qtpl:21
		qt422016.ReleaseByteBuffer(qb422016)
//line testdata/templates/integration.qtpl:21
	}
//line testdata/templates/integration.qtpl:21
	qw422016.N().S(`
		html-escaped unquoted json string: `)
//line testdata/templates/integration.qtpl:22
	{
//line testdata/templates/integration.qtpl:22
		qb422016 := qt422016.AcquireByteBuffer()
//line testdata/templates/integration.qtpl:22
		writeembeddedFunc(qb422016, p)
//line testdata/templates/integration.qtpl:22
		qw422016.E().JZ(qb422016.B)
//line testdata/templates/integration.qtpl:22
		qt422016.ReleaseByteBuffer(qb422016)
//line testdata/templates/integration.qtpl:22
	}
//line testdata/templates/integration.qtpl:22
	qw422016.N().S(`

	Html-escaped output tags:
	<ul>
		<li>`)
//line testdata/templates/integration.qtpl:26
	qw422016.E().S("<b>html-escaped `string</b>")
//line testdata/templates/integration.qtpl:26
	qw422016.N().S(`</li>
		<li>`)
//line testdata/templates/integration.qtpl:27
	qw422016.E().Z([]byte("<b>html-escaped `byte slice</b>"))
//line testdata/templates/integration.qtpl:27
	qw422016.N().S(`</li>
		<li>Int: `)
//line testdata/templates/integration.qtpl:28
	qw422016.N().D(42)
//line testdata/templates/integration.qtpl:28
	qw422016.N().S(`</li>
		<li>Float: `)
//line testdata/templates/integration.qtpl:29
	qw422016.N().F(3.14)
//line testdata/templates/integration.qtpl:29
	qw422016.N().S(`</li>
		<li>`)
//line testdata/templates/integration.qtpl:30
	qw422016.E().Q(`<quoted> "json"
				string`)
//line testdata/templates/integration.qtpl:31
	qw422016.N().S(`</li>
		<li>alert("foo `)
//line testdata/templates/integration.qtpl:32
	qw422016.E().J(`"json"-safe
				<string>`)
//line testdata/templates/integration.qtpl:33
	qw422016.N().S(` aa" + 'bar `)
//line testdata/templates/integration.qtpl:33
	qw422016.E().J(`';alert("evil")</script>`)
//line testdata/templates/integration.qtpl:33
	qw422016.N().S(`')</li>
		<li><a href="?`)
//line testdata/templates/integration.qtpl:34
	qw422016.N().U("ключ")
//line testdata/templates/integration.qtpl:34
	qw422016.N().S(`=`)
//line testdata/templates/integration.qtpl:34
	qw422016.N().U("значение&=?123")
//line testdata/templates/integration.qtpl:34
	qw422016.N().S(`">test</a></li>
		<li>`)
//line testdata/templates/integration.qtpl:35
	qw422016.E().V(struct{ A string }{A: "<b>foobar`</b>"})
//line testdata/templates/integration.qtpl:35
	qw422016.N().S(`</li>
	</ul>

	Output tags without html escaping
	<ul>
		<li>`)
//line testdata/templates/integration.qtpl:40
	qw422016.N().S("<b>html-escaped `string</b>")
//line testdata/templates/integration.qtpl:40
	qw422016.N().S(`</li>
		<li>`)
//line testdata/templates/integration.qtpl:41
	qw422016.N().Z([]byte("<b>html-escaped `byte slice</b>"))
//line testdata/templates/integration.qtpl:41
	qw422016.N().S(`</li>
		<li>Int: `)
//line testdata/templates/integration.qtpl:42
	qw422016.N().D(42)
//line testdata/templates/integration.qtpl:42
	qw422016.N().S(`</li>
		<li>Float: `)
//line testdata/templates/integration.qtpl:43
	qw422016.N().F(3.14)
//line testdata/templates/integration.qtpl:43
	qw422016.N().S(`</li>
		<li>`)
//line testdata/templates/integration.qtpl:44
	qw422016.N().Q(`<quoted> "json"
				string`)
//line testdata/templates/integration.qtpl:45
	qw422016.N().S(`</li>
		<li>alert("foo `)
//line testdata/templates/integration.qtpl:46
	qw422016.N().J(`"json"-safe
				<string>`)
//line testdata/templates/integration.qtpl:47
	qw422016.N().S(` aa" + 'bar `)
//line testdata/templates/integration.qtpl:47
	qw422016.N().J(`';alert("evil")</script>`)
//line testdata/templates/integration.qtpl:47
	qw422016.N().S(`')</li>
		<li><a href="?`)
//line testdata/templates/integration.qtpl:48
	qw422016.N().U("ключ")
//line testdata/templates/integration.qtpl:48
	qw422016.N().S(`=`)
//line testdata/templates/integration.qtpl:48
	qw422016.N().U("значение&=?123")
//line testdata/templates/integration.qtpl:48
	qw422016.N().S(`">test</a></li>
		<li>`)
//line testdata/templates/integration.qtpl:49
	qw422016.N().V(struct{ A string }{A: "<b>foobar`</b>"})
//line testdata/templates/integration.qtpl:49
	qw422016.N().S(`</li>
	</ul>

	`)
//line testdata/templates/integration.qtpl:52
	qw422016.N().S(`Strip space`)
//line testdata/templates/integration.qtpl:53
	qw422016.N().S(` `)
//line testdata/templates/integration.qtpl:53
	qw422016.N().S(`between lines and tags`)
//line testdata/templates/integration.qtpl:55
	qw422016.N().S(`
			Tags aren't parsed {%inside %}
			plain
		`)
//line testdata/templates/integration.qtpl:59
	// one-liner comment

//line testdata/templates/integration.qtpl:61
	// multi-line
	// comment

//line testdata/templates/integration.qtpl:65
	/*
	  yet another
	  multi-line comment
	*/

//line testdata/templates/integration.qtpl:70
	qw422016.N().S(`

	`)
//line testdata/templates/integration.qtpl:72
	qw422016.N().S(` Collapse space `)
//line testdata/templates/integration.qtpl:73
	qw422016.N().S(` `)
//line testdata/templates/integration.qtpl:73
	qw422016.N().S(` between `)
//line testdata/templates/integration.qtpl:74
	qw422016.N().S(`
`)
//line testdata/templates/integration.qtpl:74
	qw422016.N().S(` lines and tags `)
//line testdata/templates/integration.qtpl:78
	qw422016.N().S(` `)
//line testdata/templates/integration.qtpl:80
	for _, s := range []string{"foo", "bar", "baz"} {
//line testdata/templates/integration.qtpl:80
		qw422016.N().S(` `)
//line testdata/templates/integration.qtpl:81
		if s == "bar" {
//line testdata/templates/integration.qtpl:81
			qw422016.N().S(` Bar `)
//line testdata/templates/integration.qtpl:83
		} else if s == "baz" {
//line testdata/templates/integration.qtpl:83
			qw422016.N().S(` Baz `)
//line testdata/templates/integration.qtpl:85
			break
//line testdata/templates/integration.qtpl:86
		} else {
//line testdata/templates/integration.qtpl:86
			qw422016.N().S(` `)
//line testdata/templates/integration.qtpl:87
			if s == "never" {
//line testdata/templates/integration.qtpl:87
				qw422016.N().S(` `)
//line testdata/templates/integration.qtpl:88
				return
//line testdata/templates/integration.qtpl:89
			}
//line testdata/templates/integration.qtpl:89
			qw422016.N().S(` `)
//line testdata/templates/integration.qtpl:91
			switch s {
//line testdata/templates/integration.qtpl:92
			case "foobar":
//line testdata/templates/integration.qtpl:92
				qw422016.N().S(` s = foobar `)
//line testdata/templates/integration.qtpl:94
			case "barbaz":
//line testdata/templates/integration.qtpl:94
				qw422016.N().S(` s = barbaz `)
//line testdata/templates/integration.qtpl:96
			default:
//line testdata/templates/integration.qtpl:96
				qw422016.N().S(` s = `)
//line testdata/templates/integration.qtpl:97
				qw422016.E().S(s)
//line testdata/templates/integration.qtpl:97
				qw422016.N().S(` `)
//line testdata/templates/integration.qtpl:98
			}
//line testdata/templates/integration.qtpl:98
			qw422016.N().S(` `)
//line testdata/templates/integration.qtpl:100
			continue
//line testdata/templates/integration.qtpl:101
		}
//line testdata/templates/integration.qtpl:101
		qw422016.N().S(` `)
//line testdata/templates/integration.qtpl:102
	}
//line testdata/templates/integration.qtpl:102
	qw422016.N().S(` `)
//line testdata/templates/integration.qtpl:103
	qw422016.N().S(`

	`)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(`This is a template for integration test.
It should contains all the quicktemplate stuff.

{% import "fmt" %}

{% func Integration() %}
	Output tags`)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(` verification.

	{% code
		p := &integrationPage{
			S: "foobar",
		}
	%}
	Embedded func template:
		plain: {%= embeddedFunc(p) %}
		html-escaped: {%=h embeddedFunc(p) %}
		url-escaped: {%=u embeddedFunc(p) %}
		quoted json string: {%=q embeddedFunc(p) %}
		unquoted json string: {%=j embeddedFunc(p) %}
		html-escaped url-escaped: {%=uh embeddedFunc(p) %}
		html-escaped quoted json string: {%=qh embeddedFunc(p) %}
		html-escaped unquoted json string: {%=jh embeddedFunc(p) %}

	Html-escaped output tags:
	<ul>
		<li>{%s "<b>html-escaped `)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(`string</b>" %}</li>
		<li>{%z []byte("<b>html-escaped `)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(`byte slice</b>") %}</li>
		<li>Int: {%d 42 %}</li>
		<li>Float: {%f 3.14 %}</li>
		<li>{%q `)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(`<quoted> "json"
				string`)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(` %}</li>
		<li>alert("foo {%j `)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(`"json"-safe
				<string>`)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(` %} aa" + 'bar {%j `)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(`';alert("evil")</script>`)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(` %}')</li>
		<li><a href="?{%u "ключ" %}={%u "значение&=?123" %}">test</a></li>
		<li>{%v struct{ A string }{A: "<b>foobar`)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(`</b>"} %}</li>
	</ul>

	Output tags without html escaping
	<ul>
		<li>{%s= "<b>html-escaped `)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(`string</b>" %}</li>
		<li>{%z= []byte("<b>html-escaped `)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(`byte slice</b>") %}</li>
		<li>Int: {%d= 42 %}</li>
		<li>Float: {%f= 3.14 %}</li>
		<li>{%q= `)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(`<quoted> "json"
				string`)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(` %}</li>
		<li>alert("foo {%j= `)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(`"json"-safe
				<string>`)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(` %} aa" + 'bar {%j= `)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(`';alert("evil")</script>`)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(` %}')</li>
		<li><a href="?{%u= "ключ" %}={%u= "значение&=?123" %}">test</a></li>
		<li>{%v= struct{ A string }{A: "<b>foobar`)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S("`")
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(`</b>"} %}</li>
	</ul>

	{% stripspace %}
		Strip space {%space%}
		between lines and tags
		{%plain%}
			Tags aren't parsed {%inside %}
			plain
		{%endplain%}
		{% code // one-liner comment %}
		{% code
		// multi-line
		// comment
		%}
		{% code
		/*
		  yet another
		  multi-line comment
		*/
		%}
	{% endstripspace %}

	{% collapsespace %}
		Collapse space {%space %}
		between {%newline%} lines and tags
		{%comment%}
			Comments {%are%}
			ignored
		{%endcomment%}

		{% for _, s := range []string{"foo","bar","baz"} %}
			{% if s == "bar" %}
				Bar
			{% elseif s == "baz" %}
				Baz
				{% break %}
			{% else %}
				{% if s == "never" %}
					{% return %}
				{% endif %}

				{% switch s %}
				{% case "foobar" %}
					s = foobar
				{% case "barbaz" %}
					s = barbaz
				{% default %}
					s = {%s s %}
				{% endswitch %}

				{% continue %}
			{% endif %}
		{% endfor %}
	{% endcollapsespace %}

	{% cat "integration.qtpl" %}

	tail of the func
{% endfunc %}

{%
interface Page {
	Header()
	Body()
}
%}

{% func embeddedFunc(p Page) %}
	Page's header: {%= p.Header() %}
	Body: {%s= fmt.Sprintf("<b>%s</b>", p.Body()) %}
{% endfunc %}

{% code
type integrationPage struct {
	S string
}
%}

{% func (p *integrationPage) Header() %}Header{% endfunc %}

{% func (p *integrationPage) Body() %}
	S={%q p.S %}
{% endfunc %}
`)
//line testdata/templates/integration.qtpl:105
	qw422016.N().S(`

	tail of the func
`)
//line testdata/templates/integration.qtpl:108
}

//line testdata/templates/integration.qtpl:108
func WriteIntegration(qq422016 qtio422016.Writer) {
//line testdata/templates/integration.qtpl:108
	qw422016 := qt422016.AcquireWriter(qq422016)
//line testdata/templates/integration.qtpl:108
	StreamIntegration(qw422016)
//line testdata/templates/integration.qtpl:108
	qt422016.ReleaseWriter(qw422016)
//line testdata/templates/integration.qtpl:108
}

//line testdata/templates/integration.qtpl:108
func Integration() string {
//line testdata/templates/integration.qtpl:108
	qb422016 := qt422016.AcquireByteBuffer()
//line testdata/templates/integration.qtpl:108
	WriteIntegration(qb422016)
//line testdata/templates/integration.qtpl:108
	qs422016 := string(qb422016.B)
//line testdata/templates/integration.qtpl:108
	qt422016.ReleaseByteBuffer(qb422016)
//line testdata/templates/integration.qtpl:108
	return qs422016
//line testdata/templates/integration.qtpl:108
}

//line testdata/templates/integration.qtpl:111
type Page interface {
//line testdata/templates/integration.qtpl:111
	Header() string
//line testdata/templates/integration.qtpl:111
	StreamHeader(qw422016 *qt422016.Writer)
//line testdata/templates/integration.qtpl:111
	WriteHeader(qq422016 qtio422016.Writer)
//line testdata/templates/integration.qtpl:111
	Body() string
//line testdata/templates/integration.qtpl:111
	StreamBody(qw422016 *qt422016.Writer)
//line testdata/templates/integration.qtpl:111
	WriteBody(qq422016 qtio422016.Writer)
//line testdata/templates/integration.qtpl:111
}

//line testdata/templates/integration.qtpl:117
func streamembeddedFunc(qw422016 *qt422016.Writer, p Page) {
//line testdata/templates/integration.qtpl:117
	qw422016.N().S(`
	Page's header: `)
//line testdata/templates/integration.qtpl:118
	p.StreamHeader(qw422016)
//line testdata/templates/integration.qtpl:118
	qw422016.N().S(`
	Body: `)
//line testdata/templates/integration.qtpl:119
	qw422016.N().S(fmt.Sprintf("<b>%s</b>", p.Body()))
//line testdata/templates/integration.qtpl:119
	qw422016.N().S(`
`)
//line testdata/templates/integration.qtpl:120
}

//line testdata/templates/integration.qtpl:120
func writeembeddedFunc(qq422016 qtio422016.Writer, p Page) {
//line testdata/templates/integration.qtpl:120
	qw422016 := qt422016.AcquireWriter(qq422016)
//line testdata/templates/integration.qtpl:120
	streamembeddedFunc(qw422016, p)
//line testdata/templates/integration.qtpl:120
	qt422016.ReleaseWriter(qw422016)
//line testdata/templates/integration.qtpl:120
}

//line testdata/templates/integration.qtpl:120
func embeddedFunc(p Page) string {
//line testdata/templates/integration.qtpl:120
	qb422016 := qt422016.AcquireByteBuffer()
//line testdata/templates/integration.qtpl:120
	writeembeddedFunc(qb422016, p)
//line testdata/templates/integration.qtpl:120
	qs422016 := string(qb422016.B)
//line testdata/templates/integration.qtpl:120
	qt422016.ReleaseByteBuffer(qb422016)
//line testdata/templates/integration.qtpl:120
	return qs422016
//line testdata/templates/integration.qtpl:120
}

//line testdata/templates/integration.qtpl:123
type integrationPage struct {
	S string
}

//line testdata/templates/integration.qtpl:128
func (p *integrationPage) StreamHeader(qw422016 *qt422016.Writer) {
//line testdata/templates/integration.qtpl:128
	qw422016.N().S(`Header`)
//line testdata/templates/integration.qtpl:128
}

//line testdata/templates/integration.qtpl:128
func (p *integrationPage) WriteHeader(qq422016 qtio422016.Writer) {
//line testdata/templates/integration.qtpl:128
	qw422016 := qt422016.AcquireWriter(qq422016)
//line testdata/templates/integration.qtpl:128
	p.StreamHeader(qw422016)
//line testdata/templates/integration.qtpl:128
	qt422016.ReleaseWriter(qw422016)
//line testdata/templates/integration.qtpl:128
}

//line testdata/templates/integration.qtpl:128
func (p *integrationPage) Header() string {
//line testdata/templates/integration.qtpl:128
	qb422016 := qt422016.AcquireByteBuffer()
//line testdata/templates/integration.qtpl:128
	p.WriteHeader(qb422016)
//line testdata/templates/integration.qtpl:128
	qs422016 := string(qb422016.B)
//line testdata/templates/integration.qtpl:128
	qt422016.ReleaseByteBuffer(qb422016)
//line testdata/templates/integration.qtpl:128
	return qs422016
//line testdata/templates/integration.qtpl:128
}

//line testdata/templates/integration.qtpl:130
func (p *integrationPage) StreamBody(qw422016 *qt422016.Writer) {
//line testdata/templates/integration.qtpl:130
	qw422016.N().S(`
	S=`)
//line testdata/templates/integration.qtpl:131
	qw422016.E().Q(p.S)
//line testdata/templates/integration.qtpl:131
	qw422016.N().S(`
`)
//line testdata/templates/integration.qtpl:132
}

//line testdata/templates/integration.qtpl:132
func (p *integrationPage) WriteBody(qq422016 qtio422016.Writer) {
//line testdata/templates/integration.qtpl:132
	qw422016 := qt422016.AcquireWriter(qq422016)
//line testdata/templates/integration.qtpl:132
	p.StreamBody(qw422016)
//line testdata/templates/integration.qtpl:132
	qt422016.ReleaseWriter(qw422016)
//line testdata/templates/integration.qtpl:132
}

//line testdata/templates/integration.qtpl:132
func (p *integrationPage) Body() string {
//line testdata/templates/integration.qtpl:132
	qb422016 := qt422016.AcquireByteBuffer()
//line testdata/templates/integration.qtpl:132
	p.WriteBody(qb422016)
//line testdata/templates/integration.qtpl:132
	qs422016 := string(qb422016.B)
//line testdata/templates/integration.qtpl:132
	qt422016.ReleaseByteBuffer(qb422016)
//line testdata/templates/integration.qtpl:132
	return qs422016
//line testdata/templates/integration.qtpl:132
}
