// (c) 2018 Ovi Crisan
// http://ovi.crisan.ca
// https://github.com/ovicrisan/SiteGo

package template

import "bytes"

var (
	Hostname string
	Site     string = "Site Go"
	Logo     bool   = false
)

func Header(title string, buffer *bytes.Buffer) {
	buffer.WriteString(`<!doctype html>
	<html lang="en">
	  <head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">`)
	if Logo {
		buffer.WriteString(`
		<link rel="shortcut icon" href="/static/favicon.ico" type="image/x-icon">
		<link rel="icon" href="/static/favicon.ico" type="image/x-icon">`)
	}
	buffer.WriteString(`
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
    <title>`)
	buffer.WriteString(Site)
	buffer.WriteString(" - ")
	buffer.WriteString(title)
	buffer.WriteString(`</title>
</head>
<body>
<div class="container">
<nav class="navbar navbar-expand-lg bg-info">
<span class="navbar-brand mb-0 h1 text-white">`)
	if Logo {
		buffer.WriteString(" <img src='/static/logo.png' height='16' alt='logo' align='absmiddle' /> ")
	}
	buffer.WriteString(Site)
	buffer.WriteString(`</span>
  <button class="navbar-toggler text-white" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>
  <div class="collapse navbar-collapse" id="navbarNav">
	<ul class="navbar-nav">
	<li class="nav-item"><a class="nav-link text-white" href="/">Home</a></li>
	<li class="nav-item"><a class="nav-link text-white" href="/about">About</a></li>
	<li class="nav-item"><a class="nav-link text-white" href="/contact">Contact</a></li>
	</ul>
  </div>
</nav>
<main class="p-1" style="min-height:400px;">
`)
}

func Footer(buffer *bytes.Buffer) {
	buffer.WriteString(`</main>
<footer class="nav justify-content-end bg-info text-light text-right p-1"><small>`)
	buffer.WriteString(Hostname)
	buffer.WriteString(`</small></footer>
</div>
<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/js/bootstrap.min.js" integrity="sha384-ChfqqxuZUCnJSK3+MXmPNIyE6ZbWh2IMqE241rYiqJxyMiZ6OW/JmZQ5stwEULTy" crossorigin="anonymous"></script>
</body>
</html>`)
}
