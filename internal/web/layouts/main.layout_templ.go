// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package layouts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func MainLayout(children ...templ.Component) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\"><title>YACA</title><style>\n    [un-cloak] {\n      display: none !important;\n    }\n\n\n    .jsm {\n      font-family: \"Noto Sans\", sans-serif;\n      font-optical-sizing: auto;\n      font-weight: 500;\n      font-style: normal;\n      font-variation-settings:\n        \"wdth\" 100;\n    }\n\n    .gu-mirror {\n      position: fixed !important;\n      margin: 0 !important;\n      z-index: 9999 !important;\n      opacity: .8\n    }\n\n    .gu-hide {\n      display: none !important\n    }\n\n    .gu-unselectable {\n      -webkit-user-select: none !important;\n      -moz-user-select: none !important;\n      -ms-user-select: none !important;\n      user-select: none !important\n    }\n\n    .gu-transit {\n      opacity: .2\n    }\n  </style><link rel=\"icon\" href=\"\" type=\"image/x-icon\"><link rel=\"preconnect\" href=\"https://fonts.googleapis.com\"><link rel=\"preconnect\" href=\"https://fonts.gstatic.com\" crossorigin><link href=\"https://fonts.googleapis.com/css2?family=Noto+Sans:wght@500&amp;display=swap\" rel=\"stylesheet\"><!--    <link rel=\"manifest\" href=\"./assets/site.webmanifest\">--></head><body class=\"jsm\" un-cloak><div class=\"w-100vw min-h-100vh h-fit bg-black flex flex-row justify-center items-center align-middle\"><div class=\"w-10/12 h-full min-h-screen bg-slate-900 flex flex-col justify-center items-center align-middle space-y-10\"><div class=\"w-fit h-8/12 flex flex-col justify-center items-center align-middle space-y-10\" id=\"start\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div><script src=\"https://cdn.jsdelivr.net/npm/@unocss/runtime/preset-wind.global.js\"></script><link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/@unocss/reset/tailwind.min.css\"><script>\n    window.__unocss = {\n      presets: [\n        () => window.__unocss_runtime.presets.presetWind({\n          cdn: 'https://esm.sh/'\n        }),\n      ],\n    }\n  </script><script id=\"uno\" src=\"https://cdn.jsdelivr.net/npm/@unocss/runtime/core.global.js\"></script><script src=\"https://unpkg.com/htmx.org@2.0.0\" integrity=\"sha384-wS5l5IKJBvK6sPTKa2WZ1js3d947pvWXbPJ1OmWfEuxLgeHcEbjUUA5i9V5ZkpCw\" crossorigin=\"anonymous\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
