// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package web

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func About() templ.Component {
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
		var templ_7745c5c3_Var2 = []any{landing}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/about.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><h1>About</h1><p>Bacon ipsum dolor amet drumstick ribeye in, shoulder officia kielbasa pig laboris tempor culpa reprehenderit occaecat aliquip cupim. Alcatra fatback occaecat proident ullamco tempor laborum, pastrami corned beef ea sirloin doner. Biltong excepteur bresaola sed, ipsum meatball tail sint in ut. Occaecat magna pariatur tenderloin dolor ham boudin. Ea excepteur sausage, rump anim brisket aliqua jerky buffalo. Shoulder dolor aliquip in reprehenderit.</p><p>Alcatra enim magna, in ham hock brisket incididunt shoulder ex adipisicing leberkas turducken esse salami. Burgdoggen dolor picanha proident officia aliquip est prosciutto jowl pork belly landjaeger chicken. Venison tail meatball rump, tenderloin flank sed officia andouille shankle aute capicola ground round dolor elit. Capicola tongue strip steak dolore aliquip. Bresaola esse enim fatback flank, picanha meatloaf pariatur eu cupidatat pig ham hock nulla minim nisi. Exercitation cupim fatback dolor, pork belly ham hock deserunt venison swine biltong meatball picanha dolore corned beef lorem.</p><p>Short ribs ut nulla laborum salami non, magna ball tip sausage commodo in filet mignon duis turkey. Burgdoggen officia minim meatball hamburger. In culpa ground round strip steak shankle chislic beef in sausage boudin elit jowl consectetur brisket buffalo. Turducken pastrami prosciutto elit landjaeger incididunt bacon picanha.</p><p>Landjaeger cow rump, tongue leberkas kielbasa excepteur tail incididunt proident est pork belly commodo buffalo. Aliquip elit ut, andouille ad nisi boudin proident ball tip. Fatback meatloaf chuck pork chop, lorem cupim frankfurter do anim shank cow pork beef ribs. Leberkas reprehenderit tempor, aute est ad ex jerky frankfurter sunt.</p><p>Does your lorem ipsum text long for something a little meatier? Give our generator a try… it’s tasty!</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
