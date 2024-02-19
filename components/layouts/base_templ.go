// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package layouts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html x-data=\"{\n		theme: localStorage.getItem(&#39;theme&#39;) || &#39;light&#39;,\n		}\" @set-theme=\"theme = theme === &#39;light&#39; ? &#39;dark&#39; : &#39;light&#39;; localStorage.setItem(&#39;theme&#39;, theme);\" x-bind:data-theme=\"theme\"><head><meta charset=\"utf-8\"><title>Golang fun</title><link rel=\"stylesheet\" href=\"/css/tailwind.css\"><script src=\"https://kit.fontawesome.com/8a37954e59.js\" crossorigin=\"anonymous\"></script><script src=\"https://unpkg.com/htmx.org@1.9.10\"></script><script defer src=\"https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js\"></script></head><body class=\"flex flex-col min-h-[100vh]\" hx-boost=\"true\"><div class=\"hidden\" id=\"notification\"></div><div class=\"navbar bg-base-100\"><div class=\"flex-1\"><a class=\"btn btn-ghost text-xl\">Golang Fun</a> <a href=\"/author\" class=\"btn btn-ghost text-xl\">Authors</a> <a href=\"/test\" class=\"btn btn-ghost text-xl\">Tests</a> <a href=\"/question\" class=\"btn btn-ghost text-xl\">Questions</a></div><div class=\"flex-none\"><ul class=\"menu menu-horizontal px-1\"><li><button @click=\"$dispatch(&#39;set-theme&#39;)\" x-bind:data-theme=\"$root.dataset.theme\" x-bind:class=\"{&#39;bg-sun&#39;: theme === &#39;light&#39;, &#39;bg-moon&#39;: theme === &#39;dark&#39;}\" class=\"size-[24px] bg-transparent w-full h-full bg-no-repeat bg-center btn-ghost btn btn-sm\"></button></li><li><a>Link</a></li><li><details><summary>Parent</summary><ul class=\"p-2 bg-base-100 rounded-t-none\"><li><a>Link 1</a></li><li><a>Link 2</a></li></ul></details></li></ul></div></div><div hx-boost=\"true\" class=\"container mx-auto px-6 mt-3 min-h-[90vh] mb-10\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><footer class=\"footer p-10 bg-base-300 text-base-content mt-auto\"><nav><h6 class=\"footer-title\">Services</h6><a class=\"link link-hover\">Branding</a> <a class=\"link link-hover\">Design</a> <a class=\"link link-hover\">Marketing</a> <a class=\"link link-hover\">Advertisement</a></nav><nav><h6 class=\"footer-title\">Company</h6><a class=\"link link-hover\">About us</a> <a class=\"link link-hover\">Contact</a> <a class=\"link link-hover\">Jobs</a> <a class=\"link link-hover\">Press kit</a></nav><nav><h6 class=\"footer-title\">Social</h6><div class=\"grid grid-flow-col gap-4\"><a><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" class=\"fill-current\"><path d=\"M24 4.557c-.883.392-1.832.656-2.828.775 1.017-.609 1.798-1.574 2.165-2.724-.951.564-2.005.974-3.127 1.195-.897-.957-2.178-1.555-3.594-1.555-3.179 0-5.515 2.966-4.797 6.045-4.091-.205-7.719-2.165-10.148-5.144-1.29 2.213-.669 5.108 1.523 6.574-.806-.026-1.566-.247-2.229-.616-.054 2.281 1.581 4.415 3.949 4.89-.693.188-1.452.232-2.224.084.626 1.956 2.444 3.379 4.6 3.419-2.07 1.623-4.678 2.348-7.29 2.04 2.179 1.397 4.768 2.212 7.548 2.212 9.142 0 14.307-7.721 13.995-14.646.962-.695 1.797-1.562 2.457-2.549z\"></path></svg></a> <a><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" class=\"fill-current\"><path d=\"M19.615 3.184c-3.604-.246-11.631-.245-15.23 0-3.897.266-4.356 2.62-4.385 8.816.029 6.185.484 8.549 4.385 8.816 3.6.245 11.626.246 15.23 0 3.897-.266 4.356-2.62 4.385-8.816-.029-6.185-.484-8.549-4.385-8.816zm-10.615 12.816v-8l8 3.993-8 4.007z\"></path></svg></a> <a><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" class=\"fill-current\"><path d=\"M9 8h-3v4h3v12h5v-12h3.642l.358-4h-4v-1.667c0-.955.192-1.333 1.115-1.333h2.885v-5h-3.808c-3.596 0-5.192 1.583-5.192 4.615v3.385z\"></path></svg></a></div></nav></footer></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
