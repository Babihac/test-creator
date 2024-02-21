// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package questionComponents

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "echo-test/components"

type AnswerDefinitionRowProps struct {
	AnswerText string
	IsCorrect  bool
}

func AnswerDefinitionRow() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center gap-4\"><div class=\"lg:min-w-[350px]\"><label class=\"form-control w-full max-w-sm\"><div class=\"label\"><span class=\"label-text font-bold\">Answer</span> <span x-show=\"answer.answerText===&#39;&#39;\" class=\"label-text-alt text-red-500\">Please, fill in this field</span></div><div class=\"flex input input-bordered items-center\"><input name=\"answer[text]\" type=\"text\" placeholder=\"Enter answer\" class=\"grow bg-transparent %s\" x-model=\"answer.answerText\"> <span class=\"badge badge-warning\">Required</span></div><div class=\"label\"></div></label></div><div class=\"pt-5 flex items-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Toggle(
			components.ToggleProps{
				Label: "Correct answer",
				Name:  "answer[correct]",
				Model: "answer.isCorrect",
			}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div @click=\"removeAnswer(index)\" class=\"btn btn-circle btn-error bg-delete bg-no-repeat bg-center btn-xs\"></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
