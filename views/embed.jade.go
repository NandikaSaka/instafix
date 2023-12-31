// Code generated by "jade.go"; DO NOT EDIT.

package views

import (
	pool "github.com/valyala/bytebufferpool"
)

const (
	embed__0  = `<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"/><meta name="theme-color" content="#CE0071"/>`
	embed__1  = `<meta property="og:url" content="`
	embed__2  = `"/><meta property="og:description" content="`
	embed__3  = `"/>`
	embed__4  = `<meta http-equiv="refresh" content="`
	embed__5  = `"/></head><body>Redirecting you to the post in a moment. <a href="`
	embed__6  = `">Or click here.</a></body></html>`
	embed__7  = `<meta name="twitter:card" content="`
	embed__9  = `<meta name="twitter:title" content="`
	embed__11 = `<meta name="twitter:image" content="`
	embed__13 = `<meta name="twitter:player:width" content="`
	embed__14 = `"/><meta name="twitter:player:height" content="`
	embed__15 = `"/><meta name="twitter:player:stream" content="`
	embed__16 = `"/><meta name="twitter:player:stream:content_type" content="video/mp4"/>`
	embed__17 = `<meta property="og:site_name" content="InstaFix"/>`
	embed__18 = `<meta property="og:image" content="`
	embed__20 = `<meta property="og:video" content="`
	embed__21 = `"/><meta property="og:video:secure_url" content="`
	embed__22 = `"/><meta property="og:video:type" content="video/mp4"/><meta property="og:video:width" content="`
	embed__23 = `"/><meta property="og:video:height" content="`
	embed__25 = `<link rel="alternate" href="`
	embed__26 = `" type="application/json+oembed" title="`
)

func Embed(v *ViewsData, buffer *pool.ByteBuffer) {

	buffer.WriteString(embed__0)

	if v.Card != "" {
		buffer.WriteString(embed__7)
		WriteAll(v.Card, true, buffer)
		buffer.WriteString(embed__3)
	}
	if v.Title != "" {
		buffer.WriteString(embed__9)
		WriteAll(v.Title, true, buffer)
		buffer.WriteString(embed__3)
	}
	if v.ImageURL != "" {
		buffer.WriteString(embed__11)
		WriteAll(v.ImageURL, true, buffer)
		buffer.WriteString(embed__3)
	}
	if v.VideoURL != "" {
		buffer.WriteString(embed__13)
		WriteAll(v.Width, true, buffer)
		buffer.WriteString(embed__14)
		WriteAll(v.Height, true, buffer)
		buffer.WriteString(embed__15)
		WriteAll(v.VideoURL, true, buffer)
		buffer.WriteString(embed__16)

	}
	if v.VideoURL != "" || v.ImageURL != "" {
		buffer.WriteString(embed__17)
	}
	buffer.WriteString(embed__1)
	WriteAll(v.URL, true, buffer)
	buffer.WriteString(embed__2)
	WriteAll(v.Description, true, buffer)
	buffer.WriteString(embed__3)
	if v.ImageURL != "" {
		buffer.WriteString(embed__18)
		WriteAll(v.ImageURL, true, buffer)
		buffer.WriteString(embed__3)
	}
	if v.VideoURL != "" {
		buffer.WriteString(embed__20)
		WriteAll(v.VideoURL, true, buffer)
		buffer.WriteString(embed__21)
		WriteAll(v.VideoURL, true, buffer)
		buffer.WriteString(embed__22)
		WriteAll(v.Width, true, buffer)
		buffer.WriteString(embed__23)
		WriteAll(v.Height, true, buffer)
		buffer.WriteString(embed__3)
	}
	if v.OEmbedURL != "" {
		buffer.WriteString(embed__25)
		WriteAll(v.OEmbedURL, false, buffer)
		buffer.WriteString(embed__26)
		WriteAll(v.Title, true, buffer)
		buffer.WriteString(embed__3)
	}
	buffer.WriteString(embed__4)
	WriteAll(`0; url = `+v.URL+``, true, buffer)
	buffer.WriteString(embed__5)
	WriteAll(v.URL, true, buffer)
	buffer.WriteString(embed__6)

}
