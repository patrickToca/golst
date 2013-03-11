package main

// html.go contains the templates and code for producing an HTML listing
// from a Go source file.

import (
	"bytes"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"time"
)

// HtmlTemplate is an template designed to enhance readability.
var HtmlTemplate = `
<!doctype html>
<html>
<head>
<title>{{.Filename}}</title>
<meta charset="UTF-8">
<style type="text/css" media=screen>
/* Based on werc.cat-v.org, suckless.org and garbe.us */

/* General */
body {
	color: black;
	font-family: serif;
	margin: 0;
	padding: 0
}

/* Header */
.header { background-color: white; border: 0; text-align: center; }
.header a { border: 0; color: black; text-decoration: none; }
.header a:visited { color: black }
.midHeader img { border: 0; }

.headerTitle { font-weight: bold; margin: 0 0 0 0.5em; padding: 0.5em; }
.headerTitle a { border: 0; text-decoration: none; }
.headerTitle a:visited { color: black; }

.headerSubTitle { font-weight: bold; margin-left: 1em; }

/* Side */
/* modified from sw original by jrick (jrick.devio.us) */
#side-bar {
    clear: both;
    border: 1;
    padding-left: 1em;
    background-color: white;
    margin: 0 00%;
    text-align: center;
}

#side-bar a {
    display: inline;
    line-height:2.1em;
    white-space: nowrap;
    padding: 0.1ex 1ex 0.1em 1ex;
    color: black;
    background-color: transparent;
    text-decoration: none;
    font-weight: bold;
}

#side-bar a:visited {
    color: black;
}

/* Main Copy */
#main {
	max-width: 70em;
	color: black;
	/* margin: 0 auto 0 2em; */
        margin: 0 10% 0 20% ;
	padding: 1em 1em 1em 1em;
	border: 0;
}

#main a { color: black; text-decoration: none; font-weight: bold; }
#main a:hover { text-decoration: underline; }
#main h1, #main-copy h2 { color: black; }
#main ul { list-style-type: square; }

/* Footer */
#footer {
	background-color: white;
	color: black;
	padding: 2em;
	clear: both;
}

#footer .left { text-align: left; float: left; clear: left; }
#footer .right { text-align: right; }
#footer a { color: black; text-decoration: none; font-weight: bold; }
#footer a:hover { text-decoration: underline; }

abbr, acronym { border-bottom: 1px dotted #333; cursor: help; }
blockquote { border-left: 1px solid #333; font-style: italic; padding: 1em; }
hr { border-width: 0 0 0.1em 0; border-color: black; }

code, pre { 
        display: block;
        padding: 5px;
        font-size: 1.1em;
        border: solid;
        border-color: black;
        border-width: 1px;
} 
pre { margin-left: 2em; }
a:visited { color: black; }
</style>
</head>
<body>
<div id="main">
{{.Markdown}}
</div>
<div id="footer">
<div class="right">{{.Filename}} generated by <a href="https://gokyle.github.com/golst/">golst</a> on {{.Date}}.
</body>
</html>
`

// HtmlWriter renders the markdown listing to HTML, writing that to
// a file.
func HtmlWriter(markdown, filename string) (err error) {
	var page struct {
		Filename string
		Date     string
		Markdown template.HTML
	}
	rendered := string(blackfriday.MarkdownCommon([]byte(markdown)))
	page.Markdown = template.HTML(rendered)
	page.Date = time.Now().Format(DateFormat)
	page.Filename = filename

	tmpl, err := template.New(filename).Parse(HtmlTemplate)
	if err != nil {
		return
	}

	htmlBuffer := new(bytes.Buffer)
	err = tmpl.Execute(htmlBuffer, page)
	if err != nil {
		return
	}

        outFile := GetOutFile(filename + ".html")
	err = ioutil.WriteFile(outFile, htmlBuffer.Bytes(), 0644)
	return
}
