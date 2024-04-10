// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package web

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"
import "strings"

func InventoryListing() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`display:grid;`)
	templ_7745c5c3_CSSBuilder.WriteString(`grid-template-columns:1fr;`)
	templ_7745c5c3_CSSBuilder.WriteString(`grid-gap:1rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin-top:1rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`InventoryListing`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Listing() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`display:grid;`)
	templ_7745c5c3_CSSBuilder.WriteString(`grid-template-columns:var(--grid-cols);`)
	templ_7745c5c3_CSSBuilder.WriteString(`grid-gap:1rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin-top:1rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Listing`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func NoOverFlow() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`overflow:hidden;`)
	templ_7745c5c3_CSSBuilder.WriteString(`text-overflow:ellipsis;`)
	templ_7745c5c3_CSSBuilder.WriteString(`white-space:nowrap;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:block;`)
	templ_7745c5c3_CSSBuilder.WriteString(`max-width:10rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`NoOverFlow`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Card() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`max-height:5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`text-align:left;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:1rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Card`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func CardImg() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`width:4rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`height:4rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`objject-fit:cover;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--border-radius);`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:var(--box-shadow);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid var(--color-accent-red);`)
	templ_7745c5c3_CSSID := templ.CSSID(`CardImg`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func CardTitle() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:.75rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:0 0 0.5rem 0;`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-primary);`)
	templ_7745c5c3_CSSID := templ.CSSID(`CardTitle`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func CardLocation() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:0.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:0 0 0.5rem 0;`)
	templ_7745c5c3_CSSID := templ.CSSID(`CardLocation`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func CardDescription() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:0.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:0 0 0.5rem 0;`)
	templ_7745c5c3_CSSID := templ.CSSID(`CardDescription`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func CardDetailsSection() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`width:50%;`)
	templ_7745c5c3_CSSID := templ.CSSID(`CardDetailsSection`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func CardImgSection() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`margin-left:0.5rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`CardImgSection`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Nav() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:row;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:space-between;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin-bottom:1rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Nav`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Row() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:row;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:space-between;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Row`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Column() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Column`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func ContainerSecondary() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--color-secondary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid var(--color-accent-green);`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-primary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--border-radius);`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:var(--box-shadow);`)
	templ_7745c5c3_CSSID := templ.CSSID(`ContainerSecondary`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func ContainerPrimary() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--color-primary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid var(--color-accent-red);`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-secondary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--border-radius);`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:var(--box-shadow);`)
	templ_7745c5c3_CSSID := templ.CSSID(`ContainerPrimary`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func ContainerHeader() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:space-between;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSID := templ.CSSID(`ContainerHeader`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func LandingContainer() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`min-height:70vh;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:var(--landing-font);`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:1rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`LandingContainer`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Input() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--border-radius);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid var(--color-accent-red);`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--color-primary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-secondary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:0.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin-bottom:.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`width:-webkit-fill-available;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:block;`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:var(--box-shadow);`)
	templ_7745c5c3_CSSID := templ.CSSID(`Input`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func InputSecondary() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--border-radius);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid var(--color-accent-green);`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--color-secondary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-primary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:0.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin-bottom:.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`width:-webkit-fill-available;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:block;`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:var(--box-shadow);`)
	templ_7745c5c3_CSSID := templ.CSSID(`InputSecondary`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func InputLabel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-self:flex-start;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:0.25rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`InputLabel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func InputSmall() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`height:2rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`InputSmall`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func InputBig() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`height:5rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`InputBig`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func FieldDescriptor() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-self:flex-start;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding-left:0.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:0.25rem 0;`)
	templ_7745c5c3_CSSID := templ.CSSID(`FieldDescriptor`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func FieldLabel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`margin-top:0;`)
	templ_7745c5c3_CSSID := templ.CSSID(`FieldLabel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func CreateBtn() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`width:7rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`min-height:3rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:2rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:1.25rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`CreateBtn`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func EditBtn() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`padding-top:0.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`min-height:2rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`min-width:4rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:1rem 0;`)
	templ_7745c5c3_CSSID := templ.CSSID(`EditBtn`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func DeleteBtn() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`padding-top:0.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`min-height:2rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`min-width:5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:1rem 0;`)
	templ_7745c5c3_CSSID := templ.CSSID(`DeleteBtn`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Btn() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`align-self:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--color-primary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-secondary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--border-radius);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid var(--color-accent-red);`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:var(--box-shadow);`)
	templ_7745c5c3_CSSBuilder.WriteString(`cursor:pointer;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Btn`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func BtnSecondary() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`align-self:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--color-secondary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-primary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--border-radius);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid var(--color-accent-green);`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:var(--box-shadow);`)
	templ_7745c5c3_CSSBuilder.WriteString(`cursor:pointer;`)
	templ_7745c5c3_CSSID := templ.CSSID(`BtnSecondary`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func BtnWarn() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`align-self:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--color-accent-red);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:1px solid var(--color-accent-green);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--border-radius);`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-secondary);`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:var(--box-shadow);`)
	templ_7745c5c3_CSSBuilder.WriteString(`cursor:pointer;`)
	templ_7745c5c3_CSSID := templ.CSSID(`BtnWarn`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func SpaceBetween() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:space-between;`)
	templ_7745c5c3_CSSID := templ.CSSID(`SpaceBetween`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func InventoryItemName() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`height:1rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`InventoryItemName`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Underline() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`text-decoration:underline;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Underline`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Error() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-accent-red);`)
	templ_7745c5c3_CSSBuilder.WriteString(`text-decoration:underline;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Error`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Info() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-accent-green);`)
	templ_7745c5c3_CSSBuilder.WriteString(`text-decoration:underline;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Info`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Link() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--color-accent-red);`)
	templ_7745c5c3_CSSBuilder.WriteString(`text-decoration:underline;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Link`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Padded() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`padding:1rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Padded`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func FullSpan() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`grid-column:span 2;`)
	templ_7745c5c3_CSSID := templ.CSSID(`FullSpan`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func NoMargin() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`margin:0;`)
	templ_7745c5c3_CSSID := templ.CSSID(`NoMargin`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func AlignLeft() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`text-align:left;`)
	templ_7745c5c3_CSSID := templ.CSSID(`AlignLeft`, templ_7745c5c3_CSSBuilder.String())
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style type=\"text/css\">\n        :root {\n            --color-primary: #002A32;\n            --color-secondary: #C9C19F;\n            --color-accent-green: #4B5842;\n            --color-accent-red: #764248;\n            --color-accent-grey: #888DA7;\n            --font-family: \"Outfit\", sans-serif;\n            --border-radius: 0.5rem;\n            --box-shadow: 4.0px 8.0px 8.0px hsl(0deg 0% 0% / 0.38);\n            --font-size: 1rem;\n            --landing-font: 0.8rem;\n            --grid-cols: 1fr;\n            --overflow-width: 5rem;\n\n            color: var(--color-secondary);\n            background-color: var(--color-primary);\n            font-size: var(--font-size);\n\n            color-scheme: dark;\n            font-synthesis: none;\n            font-optical-sizing: auto;\n            text-rendering: optimizeLegibility;\n            -webkit-font-smoothing: antialiased;\n            -moz-osx-font-smoothing: grayscale;\n        }\n\n        * {\n            font-family: var(--font-family);\n        }\n\n        body {\n          margin: 0;\n          display: flex;\n          place-items: center;\n        }\n\n        @media (min-width: 768px) {\n            :root {\n                --font-size: 1.2rem;\n                --landing-font: 1.3rem;\n                --grid-cols: 1fr 1fr;\n            }\n\n            body {\n                padding: 0 25%;\n            }\n        }\n\n        #app {\n          width: 90vw;\n          height: 100%;\n          margin: 0 auto;\n          padding: 0.75rem;\n          text-align: center;\n        }\n\n        .hover{\n          transition: all .5s;\n        }\n\n        .hover:hover{\n          background-color: var(--color-accent-grey);\n          border: 1px solid var(--color-accent-red);\n          color: var(--color-primary);\n          transition: all .5s;\n          transform : translateY(-0.2rem);\n        }\n\n        .htmx-indicator {\n          display: none;\n        }\n        .htmx-request .htmx-indicator {\n          display: block;\n        }\n        .htmx-request.htmx-indicator {\n          display: block;\n        }\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func stylesMap() map[templ.CSSClass]struct{} {
	return map[templ.CSSClass]struct{}{
		// styles.templ
		InventoryListing():   {},
		Listing():            {},
		NoOverFlow():         {},
		Card():               {},
		CardImg():            {},
		CardTitle():          {},
		CardLocation():       {},
		CardDescription():    {},
		CardDetailsSection(): {},
		CardImgSection():     {},
		Nav():                {},
		Row():                {},
		Column():             {},
		ContainerSecondary(): {},
		ContainerPrimary():   {},
		ContainerHeader():    {},
		LandingContainer():   {},
		Input():              {},
		InputSecondary():     {},
		InputLabel():         {},
		InputBig():           {},
		InputSmall():         {},
		FieldDescriptor():    {},
		FieldLabel():         {},
		CreateBtn():          {},
		EditBtn():            {},
		DeleteBtn():          {},
		Btn():                {},
		BtnSecondary():       {},
		SpaceBetween():       {},
		InventoryItemName():  {},
		Underline():          {},
		Error():              {},
		Info():               {},
		Link():               {},
		Padded():             {},
		FullSpan():           {},
		NoMargin():           {},
		AlignLeft():          {},
		// base.templ
		HeaderLink():  {},
		HeaderBrand(): {},
		// garden.templ
		FieldsContainer():  {},
		Fields():           {},
		LocationField():    {},
		DescriptionField(): {},
		GardenImg():        {},
		// gardenListing.templ
		GardenCreateLink(): {},
		// gardenInventoryItem.templ
		ItemAndType():    {},
		InventoryField(): {},
		ItemImg():        {},
		ItemLabel():      {},
		// imageUploader.templ
		ImageUploaderLabel(): {},
		UploaderContainer():  {},
		GardenPic():          {},
		// landing.templ
		FeatureDescriptor(): {},
		Features():          {},
		Cta():               {},
		// search.templ
		SearchInput(): {},
		SearchForm():  {},
		// sign-in.templ
		SignInContainer(): {},
		SignInForm():      {},
		SignInBtn():       {},
		SignUpLink():      {},
		PasswordLink():    {},
		// sign-up.templ
		SignUpContainer(): {},
		SignUpForm():      {},
		SignUpBtn():       {},
		// forgotPassword.templ
		ForgotPasswordBtn(): {},
	}
}

func Styles() []templ.CSSClass {
	m := stylesMap()
	s := make([]templ.CSSClass, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}
