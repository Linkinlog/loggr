// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package web

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"
import "strings"

func containerSecondary() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--color-secondary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid var(--color-accent-green);`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-primary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--border-radius);`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:var(--box-shadow);`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:1rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column;`)
	templ_7745c5c3_CSSID := templ.CSSID(`containerSecondary`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func containerPrimary() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--color-primary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid var(--color-accent-red);`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-secondary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--border-radius);`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:var(--box-shadow);`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:1rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column;`)
	templ_7745c5c3_CSSID := templ.CSSID(`containerPrimary`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func input() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--border-radius);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid var(--color-accent-red);`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--color-primary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-secondary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`height:2rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:0.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin-bottom:.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`width:-webkit-fill-available;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:block;`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:var(--box-shadow);`)
	templ_7745c5c3_CSSID := templ.CSSID(`input`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func inputSecondary() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--border-radius);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid var(--color-accent-green);`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--color-secondary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-primary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`height:2rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:0.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin-bottom:.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`width:-webkit-fill-available;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:block;`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:var(--box-shadow);`)
	templ_7745c5c3_CSSID := templ.CSSID(`inputSecondary`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func inputLabel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-self:flex-start;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:0.25rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`inputLabel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func btn() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`align-self:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`min-width:7rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`min-height:3rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:2rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:var(--font-size);`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--color-primary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-secondary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--border-radius);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid var(--color-accent-red);`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:var(--box-shadow);`)
	templ_7745c5c3_CSSBuilder.WriteString(`cursor:pointer;`)
	templ_7745c5c3_CSSBuilder.WriteString(`transition:all 1s;`)
	templ_7745c5c3_CSSID := templ.CSSID(`btn`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func globalStyles() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style type=\"text/css\">\n        :root {\n            --color-primary: #002A32;\n            --color-secondary: #C9C19F;\n            --color-accent-green: #4B5842;\n            --color-accent-red: #764248;\n            --color-accent-grey: #888DA7;\n            --font-family: \"Outfit\", sans-serif;\n            --border-radius: 0.5rem;\n            --box-shadow: 4.0px 8.0px 8.0px hsl(0deg 0% 0% / 0.38);\n            --font-size: 1rem;\n            --grid-cols: 1fr;\n\n            color: var(--color-secondary);\n            background-color: var(--color-primary);\n            font-size: var(--font-size);\n\n            color-scheme: dark;\n            font-synthesis: none;\n            font-optical-sizing: auto;\n            text-rendering: optimizeLegibility;\n            -webkit-font-smoothing: antialiased;\n            -moz-osx-font-smoothing: grayscale;\n        }\n\n        * {\n            font-family: var(--font-family);\n        }\n\n        body {\n          margin: 0;\n          display: flex;\n          place-items: center;\n        }\n\n        @media (min-width: 768px) {\n            :root {\n                --font-size: 1.2rem;\n                --grid-cols: 1fr 1fr;\n            }\n\n            body {\n                padding: 0 25%;\n            }\n        }\n\n        #app {\n          width: 90vw;\n          height: 100%;\n          margin: 0 auto;\n          padding: 0.75rem;\n          text-align: center;\n        }\n\n        .hover-secondary:hover{\n          background-color: var(--color-secondary);\n          border: 1px solid var(--color-accent-green);\n          color: var(--color-primary);\n          transition: all .5s;\n          transform : translateY(-0.2rem);\n        }\n\n        .hover-primary:hover{\n          background-color: var(--color-primary);\n          border: 1px solid var(--color-accent-red);\n          color: var(--color-secondary);\n          transition: all .5s;\n          transform : translateY(-0.2rem);\n        }\n\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
