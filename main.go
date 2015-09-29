package main

import (
	"bytes"
	"github.com/Unknwon/com"
	"github.com/Unknwon/macaron"
	"github.com/lonnc/golang-nw"
	"github.com/zhuharev/vigenere"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Setup our handler
	m := macaron.Classic()
	m.Get("/", hello)
	m.Get("/enc", Enc)
	m.Get("/dec", Dec)

	m.Get("/vienc", ViEnc)
	m.Get("/videc", ViDec)

	str, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	e = ioutil.WriteFile("logg", []byte(str), 0777)
	if e != nil {
		panic(e)
	}
	// Create a link back to node-webkit using the environment variable
	// populated by golang-nw's node-webkit code
	nodeWebkit, err := nw.New()
	if err != nil {
		panic(err)
	}

	// Pick a random localhost port, start listening for http requests using default handler
	// and send a message back to node-webkit to redirect
	if err := nodeWebkit.ListenAndServe(m); err != nil {
		panic(err)
	}
}

func hello(ctx *macaron.Context) string {
	data, e := Asset("templates/hello.tmpl")
	if e != nil {
		return ""
	}
	t := template.Must(template.New("name").Parse(string(data)))
	buf := []byte{}
	wr := bytes.NewBuffer(buf)
	e = t.Execute(wr, map[string]interface{}{"Name": "Kirill"})
	if e != nil {
		return ""
	}
	return wr.String()
}

func caesar(r rune, shift int) rune {
	// Shift character by specified number of places.
	// ... If beyond range, shift backward or forward.

	if r == ' ' || r == '.' || r == ',' || r == '?' {
		return r
	}
	s := int(r) + shift
	if s > 'z' && s < 255 {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	} else if s > 'я' {
		return rune(s - (int('я') - int('а')))
	} else if s < 'а' && s > 256 {
		return rune(s + (int('я') - int('а')))
	}
	return rune(s)
}

func enc(in string, shift int) string {
	return strings.Map(func(r rune) rune {
		return caesar(r, shift)
	}, strings.Join(strings.Fields(strings.ToLower(in)), " "))
}

func dec(in string, shift int) string {
	return strings.Map(func(r rune) rune {
		return caesar(r, -shift)
	}, strings.Join(strings.Fields(strings.ToLower(in)), " "))
}

func Enc(ctx *macaron.Context) string {
	str := ctx.Req.FormValue("s")
	offset := ctx.Req.FormValue("offset")
	return enc(str, com.StrTo(offset).MustInt())
}

func Dec(ctx *macaron.Context) string {
	str := ctx.Req.FormValue("s")
	offset := ctx.Req.FormValue("offset")
	return dec(str, com.StrTo(offset).MustInt())
}

func ViEnc(ctx *macaron.Context) string {
	str := ctx.Req.FormValue("s")
	key := ctx.Req.FormValue("key")
	vc := vigenere.New(vigenere.EN_WITH_SPACE+vigenere.RU, key)
	return vc.Encrypt(str)
}

func ViDec(ctx *macaron.Context) string {
	str := ctx.Req.FormValue("s")
	key := ctx.Req.FormValue("key")
	vc := vigenere.New(vigenere.EN_WITH_SPACE+vigenere.RU, key)
	return vc.Decrypt(str)
}
