// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package partial

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"github.com/andrei0427/lifeofmarrow/view/helpers"
)

type RecipeEntity struct {
	Title        string                        `json:"Title"`
	MealType     MealType                      `json:"meal_type"`
	Cuisine      Cuisine                       `json:"cuisine"`
	Duration     Duration                      `json:"duration"`
	Requirements Requirements                  `json:"requirements"`
	Images       helpers.StrapiMediaCollection `json:"Images"`
	Method       string                        `json:"Method"`
	Foreword     string                        `json:"Foreword"`
	Ingredients  string                        `json:"Ingredients"`
}

type MealType struct {
	Data struct {
		Attributes struct {
			Type string `json:"Type"`
		} `json:"attributes"`
	} `json:"Data"`
}

type Cuisine struct {
	Data struct {
		Attributes struct {
			Cuisine string `json:"Cuisine"`
		} `json:"attributes"`
	} `json:"Data"`
}

type Duration struct {
	Data struct {
		Attributes struct {
			Time string `json:"Time"`
		} `json:"attributes"`
	} `json:"Data"`
}

type Requirements struct {
	Data []struct {
		id         int
		Attributes struct {
			Requirement string `json:"Requirement"`
		} `json:"attributes"`
	} `json:"Data"`
}

func RecipeTile(i int, r RecipeEntity) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-white relative hover:text-emerald-600\" _=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf(`
					on mouseover add .scale-90 to #img-%d add .text-emerald-600 to #link-%d end
					on mouseout remove .scale-90 from #img-%d remove .text-emerald-600 from #link-%d end`,
			i, i, i, i)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.SafeURL = templ.SafeURL("/")
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"opacity-0 transition-all duration-500\" _=\"init remove .opacity-0 from me\"><div id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("img-%d", i)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"transition-all ease-in-out duration-150\"><div class=\"absolute top-3 left-3 text-white font-medium text-lg z-[5]\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(r.MealType.Data.Attributes.Type) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"inline-flex uppercase items-baseline rounded-full mr-0.5 px-2.5 py-0.5 md:mt-2 lg:mt-0 bg-emerald-600\"><span class=\"sr-only\">Type</span> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(r.MealType.Data.Attributes.Type)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/partial/recipe_tile.templ`, Line: 66, Col: 40}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if len(r.Cuisine.Data.Attributes.Cuisine) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"inline-flex uppercase items-baseline rounded-full mr-0.5 px-2.5 py-0.5 md:mt-2 lg:mt-0 bg-emerald-600\"><span class=\"sr-only\">Cuisine</span> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(r.Cuisine.Data.Attributes.Cuisine)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/partial/recipe_tile.templ`, Line: 72, Col: 42}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if len(r.Duration.Data.Attributes.Time) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"inline-flex uppercase items-baseline rounded-full mr-0.5 px-2.5 py-0.5 md:mt-2 lg:mt-0 bg-emerald-600\"><span class=\"sr-only\">Duration</span> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(r.Duration.Data.Attributes.Time)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/partial/recipe_tile.templ`, Line: 78, Col: 40}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if len(r.Requirements.Data) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"inline-flex uppercase items-baseline rounded-full mr-0.5 px-2.5 py-0.5 md:mt-2 lg:mt-0 bg-emerald-600\"><span class=\"sr-only\">Requirements</span> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for i, req := range r.Requirements.Data {
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(req.Attributes.Requirement)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/partial/recipe_tile.templ`, Line: 85, Col: 36}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if len(r.Requirements.Data)-1 > i {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(",")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(r.Images.Data) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(helpers.GetStrapiMediaUrl(r.Images.Data[0].Attributes.Url)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"object-cover aspect-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img src=\"/static/img/logo-responsive.png\" class=\"object-cover aspect-1 transition-all ease-in-out duration-150\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"py-4 text-gray-900\"><h3 id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("link-%d", i)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"text-center text-xl font-serif font-bold leading-6\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(r.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/partial/recipe_tile.templ`, Line: 107, Col: 14}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3></div></a></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
