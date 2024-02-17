// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package partial

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Header(isHome bool) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header class=\"relative z-10 font-serif\"><nav class=\"mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8\" aria-label=\"Global\"><div class=\"flex lg:flex-1\"><a href=\"/\" class=\"-m-1.5 p-1.5\"><span class=\"sr-only\">LifeofMarrow</span> <img class=\"h-24 w-auto\" src=\"/static/img/logo.png\" alt=\"logo\"></a></div><div class=\"flex lg:hidden\"><button type=\"button\" class=\"-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700\"><span class=\"sr-only\">Open main menu</span> <svg class=\"h-6 w-6\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" aria-hidden=\"true\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5\"></path></svg></button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 = []any{"hidden lg:flex lg:gap-x-12", templ.KV("text-white", isHome), templ.KV("text-gray-900",
			!isHome)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var2).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><a href=\"/about\" class=\"text-lg font-medium leading-6 hover:text-emerald-600\">About</a> <a href=\"/recipes\" class=\"text-lg font-medium leading-6 hover:text-emerald-600\">Recipes</a><div class=\"relative\"><button _=\"\n          on click show #ddl-store with *ease\n          on click elsewhere hide #ddl-store with *ease\" type=\"button\" class=\"flex items-center gap-x-1 text-lg font-medium leading-6 hover:text-emerald-600\" aria-expanded=\"false\">Store & Services <svg class=\"h-5 w-5 flex-none text-gray-400\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z\" clip-rule=\"evenodd\"></path></svg></button><div id=\"ddl-store\" class=\"hidden absolute transition opacity-0 translate-y-1 -left-8 top-full z-10 mt-3 w-screen max-w-md overflow-hidden rounded-3xl bg-white shadow-lg ring-1 ring-gray-900/5\"><div class=\"p-4\"><div class=\"group relative flex gap-x-6 rounded-lg p-4 text-lg leading-6 hover:bg-gray-50\"><div class=\"mt-1 flex h-11 w-11 flex-none items-center justify-center rounded-lg bg-gray-50 group-hover:bg-white\"><svg class=\"h-6 w-6 text-gray-600 group-hover:text-emerald-600\" viewBox=\"0 0 48 48\" xml: space=\"preserve\" xmlns=\"http://www.w3.org/2000/svg\" xmlns: xlink=\"http://www.w3.org/1999/xlink\"><g id=\"Expanded\"><g><g><path d=\"M31,48c-0.553,0-1-0.447-1-1V1c0-0.553,0.447-1,1-1s1,0.447,1,1v46C32,47.553,31.553,48,31,48z\"></path></g> <g><path d=\"M37,31h-6c-0.553,0-1-0.447-1-1s0.447-1,1-1h4.996c-0.133-16.802-3.585-26.66-5.014-27.002     c-0.552,0-0.991-0.446-0.991-0.999C29.991,0.447,30.447,0,31,0c4.188,0,7,15.512,7,30C38,30.553,37.553,31,37,31z\"></path></g> <g><path d=\"M17,18c-3.859,0-7-3.141-7-7V1c0-0.553,0.447-1,1-1s1,0.447,1,1v10c0,2.757,2.243,5,5,5s5-2.243,5-5V1     c0-0.553,0.447-1,1-1s1,0.447,1,1v10C24,14.859,20.859,18,17,18z\"></path></g> <g><path d=\"M17,48c-0.553,0-1-0.447-1-1V1c0-0.553,0.447-1,1-1s1,0.447,1,1v46C18,47.553,17.553,48,17,48z\"></path></g></g></g></svg></div><div class=\"flex-auto\"><a href=\"/store/food\" class=\"block font-medium text-gray-900\">Private Dining <span class=\"absolute inset-0\"></span></a><p class=\"mt-1 text-gray-600\">Catering for all kinds of functions as well as private home-dining.</p></div></div></div><div class=\"p-4\"><div class=\"group relative flex gap-x-6 rounded-lg p-4 text-lg leading-6 hover:bg-gray-50\"><div class=\"mt-1 flex h-11 w-11 flex-none items-center justify-center rounded-lg bg-gray-50 group-hover:bg-white\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"h-6 w-6 text-gray-600 group-hover:text-emerald-600\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M13.5 21v-7.5a.75.75 0 0 1 .75-.75h3a.75.75 0 0 1 .75.75V21m-4.5 0H2.36m11.14 0H18m0 0h3.64m-1.39 0V9.349M3.75 21V9.349m0 0a3.001 3.001 0 0 0 3.75-.615A2.993 2.993 0 0 0 9.75 9.75c.896 0 1.7-.393 2.25-1.016a2.993 2.993 0 0 0 2.25 1.016c.896 0 1.7-.393 2.25-1.015a3.001 3.001 0 0 0 3.75.614m-16.5 0a3.004 3.004 0 0 1-.621-4.72l1.189-1.19A1.5 1.5 0 0 1 5.378 3h13.243a1.5 1.5 0 0 1 1.06.44l1.19 1.189a3 3 0 0 1-.621 4.72M6.75 18h3.75a.75.75 0 0 0 .75-.75V13.5a.75.75 0 0 0-.75-.75H6.75a.75.75 0 0 0-.75.75v3.75c0 .414.336.75.75.75Z\"></path></svg></div><div class=\"flex-auto\"><a href=\"/store/food\" class=\"block font-medium text-gray-900\">Food <span class=\"absolute inset-0\"></span></a><p class=\"mt-1 text-gray-600\">Artisinal & Organic food items made with love.</p></div></div></div><div class=\"p-4\"><div class=\"group relative flex gap-x-6 rounded-lg p-4 text-lg leading-6 hover:bg-gray-50\"><div class=\"mt-1 flex h-11 w-11 flex-none items-center justify-center rounded-lg bg-gray-50 group-hover:bg-white\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"h-6 w-6 text-gray-600 group-hover:text-emerald-600\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M12 6.042A8.967 8.967 0 0 0 6 3.75c-1.052 0-2.062.18-3 .512v14.25A8.987 8.987 0 0 1 6 18c2.305 0 4.408.867 6 2.292m0-14.25a8.966 8.966 0 0 1 6-2.292c1.052 0 2.062.18 3 .512v14.25A8.987 8.987 0 0 0 18 18a8.967 8.967 0 0 0-6 2.292m0-14.25v14.25\"></path></svg></div><div class=\"flex-auto\"><a href=\"/store/food\" class=\"block font-medium text-gray-900\">Books <span class=\"absolute inset-0\"></span></a><p class=\"mt-1 text-gray-600\">Explore my cullinary journey through my recipe books.</p></div></div></div></div></div></div><div class=\"hidden lg:flex lg:flex-1 lg:justify-end\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 = []any{"text-lg font-medium leading-6 hover:text-emerald-600", templ.KV("text-white",
			isHome), templ.KV("text-gray-900", !isHome)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var3...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"/contact\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var3).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Get in touch <span aria-hidden=\"true\">&rarr;</span></a></div></nav><!-- Mobile menu, show/hide based on menu open state. --><div class=\"lg:hidden\" role=\"dialog\" aria-modal=\"true\"><!-- Background backdrop, show/hide based on slide-over state. --><div class=\"fixed inset-0 z-10\"></div><div class=\"fixed inset-y-0 right-0 z-10 w-full overflow-y-auto bg-white px-6 py-6 sm:max-w-sm sm:ring-1 sm:ring-gray-900/10\"><div class=\"flex items-center justify-between\"><a href=\"/\" class=\"-m-1.5 p-1.5\"><span class=\"sr-only\">LifeofMarrow</span> <img class=\"h-24 w-auto\" src=\"/static/img/logo-responsive.png\" alt=\"side menu logo\"></a> <button type=\"button\" class=\"-m-2.5 rounded-md p-2.5 text-gray-700\"><span class=\"sr-only\">Close menu</span> <svg class=\"h-6 w-6\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" aria-hidden=\"true\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M6 18L18 6M6 6l12 12\"></path></svg></button></div><div class=\"mt-6 flow-root\"><div class=\"-my-6 divide-y divide-gray-500/10\"><div class=\"space-y-2 py-6\"><div class=\"-mx-3\"><button type=\"button\" class=\"flex w-full items-center justify-between rounded-lg py-2 pl-3 pr-3.5 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50\" aria-controls=\"disclosure-1\" aria-expanded=\"false\">Product<!--\n                  Expand/collapse icon, toggle classes based on menu open state.\n\n                  Open: \"rotate-180\", Closed: \"\"\n                --><svg class=\"h-5 w-5 flex-none\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z\" clip-rule=\"evenodd\"></path></svg></button><!-- 'Product' sub-menu, show/hide based on menu state. --><div class=\"mt-2 space-y-2\" id=\"disclosure-1\"><a href=\"#\" class=\"block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Analytics</a> <a href=\"#\" class=\"block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Engagement</a> <a href=\"#\" class=\"block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Security</a> <a href=\"#\" class=\"block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Integrations</a> <a href=\"#\" class=\"block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Automations</a> <a href=\"#\" class=\"block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Watch demo</a> <a href=\"#\" class=\"block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Contact sales</a></div></div><a href=\"#\" class=\"-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Features</a> <a href=\"#\" class=\"-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Marketplace</a><div class=\"-mx-3\"><button type=\"button\" class=\"flex w-full items-center justify-between rounded-lg py-2 pl-3 pr-3.5 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50\" aria-controls=\"disclosure-2\" aria-expanded=\"false\">Company<!--\n                  Expand/collapse icon, toggle classes based on menu open state.\n\n                  Open: \"rotate-180\", Closed: \"\"\n                --><svg class=\"h-5 w-5 flex-none\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z\" clip-rule=\"evenodd\"></path></svg></button><!-- 'Company' sub-menu, show/hide based on menu state. --><div class=\"mt-2 space-y-2\" id=\"disclosure-2\"><a href=\"#\" class=\"block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50\">About us</a> <a href=\"#\" class=\"block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Careers</a> <a href=\"#\" class=\"block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Support</a> <a href=\"#\" class=\"block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Press</a> <a href=\"#\" class=\"block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Blog</a></div></div></div><div class=\"py-6\"><a href=\"#\" class=\"-mx-3 block rounded-lg px-3 py-2.5 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Log in</a></div></div></div></div></div></header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
