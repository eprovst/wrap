package html

/* Embedded CSS */

const pageCSS = `/*!
* Wrap CSS
* Copyright Evert Provoost.
* Licensed under MIT.
*/

html,
body,
div,
span,
p,
ins,
b,
i,
u {
    margin: 0;
    padding: 0;
    border: 0;
    font-size: 100%;
    font: inherit;
    vertical-align: baseline;
    word-wrap: break-word;
}

b {
    font-weight: bold;
}

i {
    font-style: italic;
}

u {
    text-decoration: underline;
}

body {
    background: #ccc;
    text-align: center;
}

.page {
    display: inline-block;
    margin: 5ch auto;
    background: #fff;
}

@page {
    size: letter;
    width: 8.5in;
    height: 9in;
}

@media print {
    body {
        background: #fff;
    }

    .page {
        margin: 0!important;
    }

    .play>div.page-break {
        visibility: hidden!important;
        margin: 0!important;
        page-break-before: always;
    }
}
`
const playCSS = `
.play {
    background: #fff;
    font-family: 'Courier Prime', 'Courier Final Draft', 'Courier Screenplay', Courier, monospace;
    font-size: 12pt;
    text-align: left;
    width: 60ch;
    margin: 10ch 10ch 10ch 15ch;
    position: relative;
}

.play ins {
    text-decoration: none;
    color: #ccc;
}

.play>div.dual {
    display: flex;
}

.play>div {
    width: 60ch;
}

.play>.centered {
    text-align: center;
}

.play>div.page-break {
    display: block;
    height: 1px;
    border: 0;
    border-top: 1px solid #ccc;
    margin: 5ch 0 5ch -5ch;
    padding: 0 5ch 0 0;
}

.play>div.slug>span.scnuml,
.play>div.slug>span.scnumr {
    display: block;
    float: left;
    font-weight: inherit;
    -webkit-touch-callout: none;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
}

.play>div.slug>span.scnuml {
    float: left;
}

.play>div.slug>span.scnumr {
    float: right;
}
`

var styleCSS = map[string]string{"screen": `
.screen>div.slug {
    margin-top: 2em;
    text-transform: uppercase;
}

.screen>div.act {
    padding-left: 25ch;
    width: 35ch;
    text-transform: uppercase;
    text-decoration: underline;
}

.screen>div.action {
    margin-top: 1em;
}

.screen>div.lyrics {
    margin-top: 1em;
    font-style: italic;
}

.screen>div.dialog,
.screen>div.dual {
    margin-top: 1em;
}

.screen>div.transition {
    margin-top: 1em;
    padding-left: 45ch;
    width: 15ch;
}

.screen>div.dialog>p.character {
    padding-left: 22ch;
    padding-right: 0;
}

.screen>div.dialog>p.parenthetical {
    padding-left: 17ch;
    text-indent: -1ch;
    padding-right: 19ch;
}

.screen>div.dialog>p.lyrics {
    font-style: italic;
}

.screen>div.dialog>p {
    padding-left: 10ch;
    padding-right: 10ch;
}

.screen>div.dual>div {
    float: left;
}

.screen>div.dual>div>p.character {
    padding-left: 12ch;
}

.screen>div.dual>div>p.parenthetical {
    padding-left: 3ch;
    text-indent: -1ch;
    padding-right: 5ch;
}

.screen>div.dual>div>p.lyrics {
    font-style: italic;
}

.screen>div.dual>div.left {
    padding-left: 0;
    width: 29ch;
}

.screen>div.dual>div.right {
    margin-left: 2ch;
    width: 29ch;
}

.screen>div.slug>span.scnuml {
    margin-left: -7.5ch;
}

.screen>div.slug>span.scnumr {
    margin-right: -2.5ch;
}
`,
	"stage": `
.stage>div.slug {
    margin-top: 1em;
    padding-left: 25ch;
    width: 35ch;
    text-decoration: underline;
}

.stage>div.act {
    margin-top: 0;
    padding-left: 25ch;
    width: 35ch;
    text-transform: uppercase;
    text-decoration: underline;
}

.stage>div.action {
    margin-top: 1em;
    padding-left: 12.5ch;
    width: 33ch;
}

.stage>div.lyrics {
    margin-top: 1em;
    font-style: italic;
}

.stage>div.dialog,
.stage>div.dual {
    margin-top: 1em;
}

.stage>div.transition {
    margin-top: 1em;
    padding-left: 25ch;
    width: 35ch;
}

.stage>div.dialog>p.character {
    padding-left: 25ch;
}

.stage>div.dialog>p.parenthetical {
    padding-left: 12.5ch;
    width: 33ch;
}

.stage>div.dialog>p.lyrics {
    font-style: italic;
}

.stage>div.dual>div {
    float: left;
}

.stage>div.dual>div>p.character {
    margin-left: 8ch;
}

.stage>div.dual>div>p.parenthetical {
    padding-left: 3.5ch;
    width: 18ch;
}

.stage>div.dual>div>p.lyrics {
    font-style: italic;
}

.stage>div.dual>div.left {
    width: 28ch;
}

.stage>div.dual>div.right {
    margin-left: 4ch;
    width: 28ch;
}

.stage>div.slug>span.scnuml {
    margin-left: -32.5ch;
}

.stage>div.slug>span.scnumr {
    margin-right: -2.5ch;
}
`}
