// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/view/partial"
)

type RecipeFilters struct {
	MealTypes    internal.StrapiCollectionResponse[partial.MealType]
	Cuisines     internal.StrapiCollectionResponse[partial.Cuisine]
	Durations    internal.StrapiCollectionResponse[partial.Duration]
	Requirements internal.StrapiCollectionResponse[partial.Requirements]
}

func Recipes(props []partial.RecipeEntity, filters RecipeFilters) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"pb-16\"><div class=\"relative z-40 sm:hidden\" role=\"dialog\" aria-modal=\"true\"><!--\n      Off-canvas menu backdrop, show/hide based on off-canvas menu state.\n\n      Entering: \"transition-opacity ease-linear duration-300\"\n        From: \"opacity-0\"\n        To: \"opacity-100\"\n      Leaving: \"transition-opacity ease-linear duration-300\"\n        From: \"opacity-100\"\n        To: \"opacity-0\"\n    --><div class=\"fixed inset-0 bg-black bg-opacity-25\"></div><div class=\"fixed inset-0 z-40 flex\"><!--\n        Off-canvas menu, show/hide based on off-canvas menu state.\n\n        Entering: \"transition ease-in-out duration-300 transform\"\n          From: \"translate-x-full\"\n          To: \"translate-x-0\"\n        Leaving: \"transition ease-in-out duration-300 transform\"\n          From: \"translate-x-0\"\n          To: \"translate-x-full\"\n      --><div class=\"relative ml-auto flex h-full w-full max-w-xs flex-col overflow-y-auto bg-white py-4 pb-12 shadow-xl\"><div class=\"flex items-center justify-between px-4\"><h2 class=\"text-lg font-medium text-gray-900\">Filters</h2><button type=\"button\" class=\"-mr-2 flex h-10 w-10 items-center justify-center rounded-md bg-white p-2 text-gray-400\"><span class=\"sr-only\">Close menu</span> <svg class=\"h-6 w-6\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" aria-hidden=\"true\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M6 18L18 6M6 6l12 12\"></path></svg></button></div><!-- Filters --><form class=\"mt-4\"><div class=\"border-t border-gray-200 px-4 py-6\"><h3 class=\"-mx-2 -my-3 flow-root\"><!-- Expand/collapse section button --><button type=\"button\" class=\"flex w-full items-center justify-between bg-white px-2 py-3 text-sm text-gray-400\" aria-controls=\"filter-section-0\" aria-expanded=\"false\"><span class=\"font-medium text-gray-900\">Category</span> <span class=\"ml-6 flex items-center\"><!--\n                    Expand/collapse icon, toggle classes based on section open state.\n\n                    Open: \"-rotate-180\", Closed: \"rotate-0\"\n                  --><svg class=\"rotate-0 h-5 w-5 transform\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z\" clip-rule=\"evenodd\"></path></svg></span></button></h3><!-- Filter section, show/hide based on section state. --><div class=\"pt-6\" id=\"filter-section-0\"><div class=\"space-y-6\"><div class=\"flex items-center\"><input id=\"filter-mobile-category-0\" name=\"category[]\" value=\"new-arrivals\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500\"> <label for=\"filter-mobile-category-0\" class=\"ml-3 text-sm text-gray-500\">All New Arrivals</label></div><div class=\"flex items-center\"><input id=\"filter-mobile-category-1\" name=\"category[]\" value=\"tees\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500\"> <label for=\"filter-mobile-category-1\" class=\"ml-3 text-sm text-gray-500\">Tees</label></div><div class=\"flex items-center\"><input id=\"filter-mobile-category-2\" name=\"category[]\" value=\"objects\" type=\"checkbox\" checked class=\"h-4 w-4 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500\"> <label for=\"filter-mobile-category-2\" class=\"ml-3 text-sm text-gray-500\">Objects</label></div></div></div></div><div class=\"border-t border-gray-200 px-4 py-6\"><h3 class=\"-mx-2 -my-3 flow-root\"><!-- Expand/collapse section button --><button type=\"button\" class=\"flex w-full items-center justify-between bg-white px-2 py-3 text-sm text-gray-400\" aria-controls=\"filter-section-1\" aria-expanded=\"false\"><span class=\"font-medium text-gray-900\">Color</span> <span class=\"ml-6 flex items-center\"><!--\n                    Expand/collapse icon, toggle classes based on section open state.\n\n                    Open: \"-rotate-180\", Closed: \"rotate-0\"\n                  --><svg class=\"rotate-0 h-5 w-5 transform\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z\" clip-rule=\"evenodd\"></path></svg></span></button></h3><!-- Filter section, show/hide based on section state. --><div class=\"pt-6\" id=\"filter-section-1\"><div class=\"space-y-6\"><div class=\"flex items-center\"><input id=\"filter-mobile-color-0\" name=\"color[]\" value=\"white\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500\"> <label for=\"filter-mobile-color-0\" class=\"ml-3 text-sm text-gray-500\">White</label></div><div class=\"flex items-center\"><input id=\"filter-mobile-color-1\" name=\"color[]\" value=\"beige\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500\"> <label for=\"filter-mobile-color-1\" class=\"ml-3 text-sm text-gray-500\">Beige</label></div><div class=\"flex items-center\"><input id=\"filter-mobile-color-2\" name=\"color[]\" value=\"blue\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500\"> <label for=\"filter-mobile-color-2\" class=\"ml-3 text-sm text-gray-500\">Blue</label></div></div></div></div><div class=\"border-t border-gray-200 px-4 py-6\"><h3 class=\"-mx-2 -my-3 flow-root\"><!-- Expand/collapse section button --><button type=\"button\" class=\"flex w-full items-center justify-between bg-white px-2 py-3 text-sm text-gray-400\" aria-controls=\"filter-section-2\" aria-expanded=\"false\"><span class=\"font-medium text-gray-900\">Sizes</span> <span class=\"ml-6 flex items-center\"><!--\n                    Expand/collapse icon, toggle classes based on section open state.\n\n                    Open: \"-rotate-180\", Closed: \"rotate-0\"\n                  --><svg class=\"rotate-0 h-5 w-5 transform\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z\" clip-rule=\"evenodd\"></path></svg></span></button></h3><!-- Filter section, show/hide based on section state. --><div class=\"pt-6\" id=\"filter-section-2\"><div class=\"space-y-6\"><div class=\"flex items-center\"><input id=\"filter-mobile-sizes-0\" name=\"sizes[]\" value=\"s\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500\"> <label for=\"filter-mobile-sizes-0\" class=\"ml-3 text-sm text-gray-500\">S</label></div><div class=\"flex items-center\"><input id=\"filter-mobile-sizes-1\" name=\"sizes[]\" value=\"m\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500\"> <label for=\"filter-mobile-sizes-1\" class=\"ml-3 text-sm text-gray-500\">M</label></div><div class=\"flex items-center\"><input id=\"filter-mobile-sizes-2\" name=\"sizes[]\" value=\"l\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500\"> <label for=\"filter-mobile-sizes-2\" class=\"ml-3 text-sm text-gray-500\">L</label></div></div></div></div></form></div></div></div><div class=\"mx-auto max-w-7xl px-4 py-16 sm:px-6 lg:px-8\"><h1 class=\"text-3xl font-bold tracking-tight text-gray-900\">Recipes</h1><p class=\"mt-4 max-w-xl text-sm text-gray-700\">Browse through our handpicked favourites that are both healthy and flavorful.</p></div><!-- Filters --><section aria-labelledby=\"filter-heading\"><h2 id=\"filter-heading\" class=\"sr-only\">Filters</h2><div class=\"border-b border-gray-200 bg-white pb-4\"><div class=\"mx-auto flex max-w-7xl items-center justify-between px-4 sm:px-6 lg:px-8\"><!-- Mobile filter dialog toggle, controls the 'mobileFiltersOpen' state. --><button type=\"button\" class=\"inline-block text-sm font-medium text-gray-700 hover:text-gray-900 sm:hidden\">Filters</button><div class=\"hidden sm:block\"><div class=\"flow-root\"><div class=\"-mx-4 flex items-center divide-x divide-gray-200\"><div class=\"relative inline-block px-4 text-left\"><button type=\"button\" class=\"group inline-flex justify-center text-sm font-medium text-gray-700 hover:text-gray-900\" aria-expanded=\"false\" _=\"on click show #ddl-mealtype with *ease\n									    on click elsewhere hide #ddl-mealtype with *ease\"><span>Meal Type</span> <span _=\"\n										on updateMealtypeFilter\n											if $mealtypecounts &gt; 0 then\n											   remove .hidden from me\n											   put $mealtypecounts into me\n											else\n											   add .hidden to me\n										end\n										\" id=\"mealtype-count\" class=\"hidden ml-1.5 rounded bg-emerald-600 px-1.5 py-0.5 text-xs font-semibold tabular-nums text-white\">0</span> <svg class=\"-mr-1 ml-1 h-5 w-5 flex-shrink-0 text-gray-400 group-hover:text-gray-500\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z\" clip-rule=\"evenodd\"></path></svg></button><div id=\"ddl-mealtype\" class=\"hidden absolute right-0 z-10 mt-2 origin-top-right rounded-md bg-white p-4 shadow-2xl ring-1 ring-black ring-opacity-5 focus:outline-none\"><form class=\"space-y-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, mt := range filters.MealTypes.Data {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center\"><input id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("mealtype-%d", mt.Id)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" name=\"mealtype[]\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(mt.Attributes.Type))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500\" _=\"on change if me.checked then increment $mealtypecounts else decrement $mealtypecounts end trigger updateMealtypeFilter on #mealtype-count\"> <label for=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("mealtype-%d", mt.Id)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"ml-3 whitespace-nowrap pr-6 text-sm font-medium text-gray-900\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(mt.Attributes.Type)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/pages/recipes.templ`, Line: 213, Col: 152}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</form></div></div><div class=\"relative inline-block px-4 text-left\"><button type=\"button\" class=\"group inline-flex justify-center text-sm font-medium text-gray-700 hover:text-gray-900\" aria-expanded=\"false\" _=\"on click show #ddl-cuisine with *ease\n									on click elsewhere hide #ddl-cuisine with *ease\"><span>Cuisine</span> <span _=\"\n										on updateCuisineFilter\n											if $cuisinecounts &gt; 0 then\n											   remove .hidden from me\n											   put $cuisinecounts into me\n											else\n											   add .hidden to me\n										end\n										\" id=\"cuisine-count\" class=\"hidden ml-1.5 rounded bg-emerald-600 px-1.5 py-0.5 text-xs font-semibold tabular-nums text-white\">0</span> <svg class=\"-mr-1 ml-1 h-5 w-5 flex-shrink-0 text-gray-400 group-hover:text-gray-500\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z\" clip-rule=\"evenodd\"></path></svg></button><div id=\"ddl-cuisine\" class=\"hidden absolute right-0 z-10 mt-2 origin-top-right rounded-md bg-white p-4 shadow-2xl ring-1 ring-black ring-opacity-5 focus:outline-none\"><form class=\"space-y-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, c := range filters.Cuisines.Data {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center\"><input id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("cuisine-%d", c.Id)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" name=\"cuisine[]\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(c.Attributes.Cuisine))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500\" _=\"on change if me.checked then increment $cuisinecounts else decrement $cuisinecounts end trigger updateCuisineFilter on #cuisine-count\"> <label for=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("cuisine-%d", c.Id)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"ml-3 whitespace-nowrap pr-6 text-sm font-medium text-gray-900\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(c.Attributes.Cuisine)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/pages/recipes.templ`, Line: 260, Col: 36}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</form></div></div><div class=\"relative inline-block px-4 text-left\"><button type=\"button\" class=\"group inline-flex justify-center text-sm font-medium text-gray-700 hover:text-gray-900\" aria-expanded=\"false\" _=\"on click show #ddl-duration with *ease\n									on click elsewhere hide #ddl-duration with *ease\"><span>Duration</span> <span _=\"\n										on updateDurationFilter\n											if $durationcounts &gt; 0 then\n											   remove .hidden from me\n											   put $durationcounts into me\n											else\n											   add .hidden to me\n										end\n										\" id=\"duration-count\" class=\"hidden ml-1.5 rounded bg-emerald-600 px-1.5 py-0.5 text-xs font-semibold tabular-nums text-white\">0</span> <svg class=\"-mr-1 ml-1 h-5 w-5 flex-shrink-0 text-gray-400 group-hover:text-gray-500\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z\" clip-rule=\"evenodd\"></path></svg></button><div id=\"ddl-duration\" class=\"hidden absolute right-0 z-10 mt-2 origin-top-right rounded-md bg-white p-4 shadow-2xl ring-1 ring-black ring-opacity-5 focus:outline-none\"><form class=\"space-y-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, d := range filters.Durations.Data {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center\"><input id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("duration-%d", d.Id)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" name=\"duration[]\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(d.Attributes.Time))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500\" _=\"on change if me.checked then increment $durationcounts else decrement $durationcounts end trigger updateDurationFilter on #duration-count\"> <label for=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("duration-%d", d.Id)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"ml-3 whitespace-nowrap pr-6 text-sm font-medium text-gray-900\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(d.Attributes.Time)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/pages/recipes.templ`, Line: 304, Col: 150}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</form></div></div><div class=\"relative inline-block px-4 text-left\"><button type=\"button\" class=\"group inline-flex justify-center text-sm font-medium text-gray-700 hover:text-gray-900\" aria-expanded=\"false\" _=\"on click show #ddl-requirements with *ease\n									on click elsewhere hide #ddl-requirements with *ease\"><span>Requirements</span> <span _=\"\n										on updateRequirementsFilter\n											if $requirementscount &gt; 0 then\n											   remove .hidden from me\n											   put $requirementscount into me\n											else\n											   add .hidden to me\n										end\n										\" id=\"requirements-count\" class=\"hidden ml-1.5 rounded bg-emerald-600 px-1.5 py-0.5 text-xs font-semibold tabular-nums text-white\">0</span> <svg class=\"-mr-1 ml-1 h-5 w-5 flex-shrink-0 text-gray-400 group-hover:text-gray-500\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z\" clip-rule=\"evenodd\"></path></svg></button><div id=\"ddl-requirements\" class=\"hidden absolute right-0 z-10 mt-2 origin-top-right rounded-md bg-white p-4 shadow-2xl ring-1 ring-black ring-opacity-5 focus:outline-none\"><form class=\"space-y-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, r := range filters.Requirements.Data {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center\"><input id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("req-%d", r.Id)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" name=\"requirement[]\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(r.Attributes.Requirement))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500\" _=\"on change if me.checked then increment $requirementscount else decrement $requirementscount end trigger updateRequirementsFilter on #requirements-count\"> <label for=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("req-%d", r.Id)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"ml-3 whitespace-nowrap pr-6 text-sm font-medium text-gray-900\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(r.Attributes.Requirement)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/pages/recipes.templ`, Line: 348, Col: 152}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</form></div></div></div></div></div><div class=\"relative inline-block text-left\"><div class=\"relative mt-2 rounded-md shadow-sm\"><div class=\"pointer-events-none absolute inset-y-0 left-0 flex items-center pl-2\"><span class=\"text-gray-500 sm:text-sm\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-4 h-4\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z\"></path></svg></span></div><input type=\"text\" name=\"ingredient\" id=\"ingredient\" class=\"block w-full rounded-md border-0 py-1.5 pl-7 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-emerald-600 sm:text-sm sm:leading-6\" placeholder=\"Aubergine\"></div></div></div></div><!-- Active filters --><div class=\"bg-gray-100\"><div class=\"mx-auto max-w-7xl px-4 py-3 sm:flex sm:items-center sm:px-6 lg:px-8\"><h3 class=\"text-sm font-medium text-gray-500\">Filters <span class=\"sr-only\">, active</span></h3><div aria-hidden=\"true\" class=\"hidden h-5 w-px bg-gray-300 sm:ml-4 sm:block\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, mt := range filters.MealTypes.Data {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("filter-mt-%d", mt.Id)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"mt-2 sm:ml-4 sm:mt-0\"><div class=\"-m-1 flex flex-wrap items-center\"><span class=\"m-1 inline-flex items-center rounded-full border border-gray-200 bg-white py-1.5 pl-3 pr-2 text-sm font-medium text-gray-900\"><span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(mt.Attributes.Type)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/pages/recipes.templ`, Line: 383, Col: 35}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <button type=\"button\" class=\"ml-1 inline-flex h-4 w-4 flex-shrink-0 rounded-full p-1 text-gray-400 hover:bg-gray-200 hover:text-gray-500\"><span class=\"sr-only\">Remove filter for ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(mt.Attributes.Type)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/pages/recipes.templ`, Line: 385, Col: 70}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <svg class=\"h-2 w-2\" stroke=\"currentColor\" fill=\"none\" viewBox=\"0 0 8 8\"><path stroke-linecap=\"round\" stroke-width=\"1.5\" d=\"M1 1l6 6m0-6L1 7\"></path></svg></button></span></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		for _, c := range filters.Cuisines.Data {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("filter-c-%d", c.Id)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"mt-2 sm:ml-4 sm:mt-0\"><div class=\"-m-1 flex flex-wrap items-center\"><span class=\"m-1 inline-flex items-center rounded-full border border-gray-200 bg-white py-1.5 pl-3 pr-2 text-sm font-medium text-gray-900\"><span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(c.Attributes.Cuisine)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/pages/recipes.templ`, Line: 398, Col: 37}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <button type=\"button\" class=\"ml-1 inline-flex h-4 w-4 flex-shrink-0 rounded-full p-1 text-gray-400 hover:bg-gray-200 hover:text-gray-500\"><span class=\"sr-only\">Remove filter for ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(c.Attributes.Cuisine)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/pages/recipes.templ`, Line: 400, Col: 72}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <svg class=\"h-2 w-2\" stroke=\"currentColor\" fill=\"none\" viewBox=\"0 0 8 8\"><path stroke-linecap=\"round\" stroke-width=\"1.5\" d=\"M1 1l6 6m0-6L1 7\"></path></svg></button></span></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		for _, d := range filters.Durations.Data {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("filter-d-%d", d.Id)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"mt-2 sm:ml-4 sm:mt-0\"><div class=\"-m-1 flex flex-wrap items-center\"><span class=\"m-1 inline-flex items-center rounded-full border border-gray-200 bg-white py-1.5 pl-3 pr-2 text-sm font-medium text-gray-900\"><span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(d.Attributes.Time)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/pages/recipes.templ`, Line: 413, Col: 34}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <button type=\"button\" class=\"ml-1 inline-flex h-4 w-4 flex-shrink-0 rounded-full p-1 text-gray-400 hover:bg-gray-200 hover:text-gray-500\"><span class=\"sr-only\">Remove filter for ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(d.Attributes.Time)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/pages/recipes.templ`, Line: 415, Col: 69}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <svg class=\"h-2 w-2\" stroke=\"currentColor\" fill=\"none\" viewBox=\"0 0 8 8\"><path stroke-linecap=\"round\" stroke-width=\"1.5\" d=\"M1 1l6 6m0-6L1 7\"></path></svg></button></span></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		for _, r := range filters.Requirements.Data {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("filter-r-%d", r.Id)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"mt-2 sm:ml-4 sm:mt-0\"><div class=\"-m-1 flex flex-wrap items-center\"><span class=\"m-1 inline-flex items-center rounded-full border border-gray-200 bg-white py-1.5 pl-3 pr-2 text-sm font-medium text-gray-900\"><span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(r.Attributes.Requirement)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/pages/recipes.templ`, Line: 428, Col: 41}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <button type=\"button\" class=\"ml-1 inline-flex h-4 w-4 flex-shrink-0 rounded-full p-1 text-gray-400 hover:bg-gray-200 hover:text-gray-500\"><span class=\"sr-only\">Remove filter for ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(r.Attributes.Requirement)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/pages/recipes.templ`, Line: 430, Col: 76}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <svg class=\"h-2 w-2\" stroke=\"currentColor\" fill=\"none\" viewBox=\"0 0 8 8\"><path stroke-linecap=\"round\" stroke-width=\"1.5\" d=\"M1 1l6 6m0-6L1 7\"></path></svg></button></span></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></section><div class=\"inline-grid border bg-emerald-600 gap-[1px] py-[1px] grid-cols-1 lg:grid-cols-2 xl:grid-cols-3\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i, r := range props {
			templ_7745c5c3_Err = partial.RecipeTile(r, 1, len(props) == i+1).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
