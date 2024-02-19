// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package testComponents

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "echo-test/components"
import "echo-test/context"

type CreateTestStepOneProps struct {
	Erors map[string]string
}

func CreateTestStepOne(props CreateTestStepOneProps) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-init=\"currentStep = 1\" class=\"flex flex-col justify-between flex-1 h-full\"><div class=\"grid grid-cols-2 lg:grid-cols-5 gap-5 w-full\"><div id=\"test-name\" class=\"col-span-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.TextInput(components.TextInputProps{Name: "test-name", Label: "Test name", Placeholder: "Enter test name", ID: "test-name-input", Required: true, Model: "testName", Error: props.Erors["TestName"]}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div id=\"test-duration\" class=\"col-span-3\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.RangeInput(components.RangeInputProps{Name: "test-duration", Label: "Test duration", ID: "test-duration-input", Min: 15, Max: 180, Step: 15, Value: 60, Unit: "minutes", Model: "testDuration"}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div id=\"teacher-name\" class=\"col-span-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.SelectInput(components.SelectInputProps{Name: "teacher-id", Label: "Test type", ID: "test-type-input", Description: "Supervising teacher", ContextKey: contextValues.USER_SUGGESTIONS, Model: "teacherId"}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div class=\"ml-auto mt-auto mb-4\"><button hx-post=\"/test/new/step2\" hx-push-url=\"true\" hx-target=\"#create-test-form-step-2\" type=\"button\" class=\"btn btn-primary self-end\">Next Step</button></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
