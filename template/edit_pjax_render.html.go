// Code generated by hero.
// source: /Users/chenhg5/go/src/goAdmin/resources/pages/edit_pjax_render.html
// DO NOT EDIT!
package template

import (
	"bytes"
	"goAdmin/models"

	"github.com/shiyanhui/hero"
)

func EditPanelPjax(formData []models.FormStruct, url string, previous string, id string, title string, description string, buffer *bytes.Buffer) {
	buffer.WriteString(`<section class="content-header">
    `)
	buffer.WriteString(`
<h1>
    `)
	hero.EscapeHTML(title, buffer)
	buffer.WriteString(`
    <small>`)
	hero.EscapeHTML(description, buffer)
	buffer.WriteString(`</small>
</h1>
`)

	buffer.WriteString(`
</section>
<section class="content">
    <div class="row">
        <div class="col-md-12">
            <div class="box box-info">
                <div class="box-header with-border">
                    <h3 class="box-title">Edit</h3>
                    <div class="box-tools">
                        <div class="btn-group pull-right" style="margin-right: 10px">
                            <a href="" class="btn btn-sm btn-default"><i class="fa fa-list"></i> List</a>
                        </div> <div class="btn-group pull-right" style="margin-right: 10px">
                        <a class="btn btn-sm btn-default form-history-back"><i class="fa fa-arrow-left"></i> Back</a>
                    </div>
                    </div>
                </div>
                <!-- /.box-header -->
                <!-- form start -->
                `)
	buffer.WriteString(`
<form action='`)
	hero.EscapeHTML(url, buffer)
	buffer.WriteString(`' method="post" accept-charset="UTF-8" class="form-horizontal" pjax-container>
`)

	buffer.WriteString(`
                    <div class="box-body">
                        <div class="fields-group">
                            `)
	for _, data := range formData {
		buffer.WriteString(`
<div class="form-group ">
    `)
		if data.FormType == "default" {
			buffer.WriteString(`
    <label class="col-sm-2 control-label">`)
			hero.EscapeHTML(data.Head, buffer)
			buffer.WriteString(`</label>
    <div class="col-sm-8">
        <div class="box box-solid box-default no-margin">
            <!-- /.box-header -->
            <div class="box-body">
                `)
			hero.EscapeHTML(data.Value, buffer)
			buffer.WriteString(` 
            </div>
            <!-- /.box-body -->
        </div>
    </div>
`)
		} else if data.FormType == "text" {
			buffer.WriteString(`
    <label for="`)
			hero.EscapeHTML(data.Field, buffer)
			buffer.WriteString(`" class="col-sm-2 control-label">`)
			hero.EscapeHTML(data.Head, buffer)
			buffer.WriteString(`</label>
    <div class="col-sm-8">
        <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-pencil fa-fw"></i></span>
            <input type="text" id="`)
			hero.EscapeHTML(data.Field, buffer)
			buffer.WriteString(`" name="`)
			hero.EscapeHTML(data.Field, buffer)
			buffer.WriteString(`" value='`)
			hero.EscapeHTML(data.Value, buffer)
			buffer.WriteString(`' class="form-control json" placeholder="Input `)
			hero.EscapeHTML(data.Head, buffer)
			buffer.WriteString(`">
        </div>
    </div>
`)
		} else if data.FormType == "password" {
			buffer.WriteString(`
    <label for="`)
			hero.EscapeHTML(data.Field, buffer)
			buffer.WriteString(`" class="col-sm-2 control-label">`)
			hero.EscapeHTML(data.Head, buffer)
			buffer.WriteString(`</label>
    <div class="col-sm-8">
        <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-eye-slash"></i></span>
            <input type="password" id="`)
			hero.EscapeHTML(data.Field, buffer)
			buffer.WriteString(`" name="`)
			hero.EscapeHTML(data.Field, buffer)
			buffer.WriteString(`" value="`)
			hero.EscapeHTML(data.Value, buffer)
			buffer.WriteString(`" class="form-control password" placeholder="Input `)
			hero.EscapeHTML(data.Head, buffer)
			buffer.WriteString(`">
        </div>
    </div>
`)
		} else if data.FormType == "textarea" {
			buffer.WriteString(`
    <label for="`)
			hero.EscapeHTML(data.Field, buffer)
			buffer.WriteString(`" class="col-sm-2 control-label">`)
			hero.EscapeHTML(data.Head, buffer)
			buffer.WriteString(`</label>
    <div class="col-sm-8">
        <textarea name="`)
			hero.EscapeHTML(data.Field, buffer)
			buffer.WriteString(`" class="form-control" rows="5" placeholder="Input `)
			hero.EscapeHTML(data.Head, buffer)
			buffer.WriteString(`">`)
			hero.EscapeHTML(data.Value, buffer)
			buffer.WriteString(`</textarea>
    </div>
`)
		} else if data.FormType == "select" {
			buffer.WriteString(`
    <label for="`)
			hero.EscapeHTML(data.Field, buffer)
			buffer.WriteString(`" class="col-sm-2 control-label">`)
			hero.EscapeHTML(data.Head, buffer)
			buffer.WriteString(`</label>
    <div class="col-sm-8">
        <select class="form-control http_method select2-hidden-accessible" style="width: 100%;" name="`)
			hero.EscapeHTML(data.Field, buffer)
			buffer.WriteString(`[]" multiple="" data-placeholder="Input HTTP method" tabindex="-1" aria-hidden="true">
            `)
			for _, v := range data.Options {
				buffer.WriteString(`
                <option value='`)
				hero.EscapeHTML(v["value"], buffer)
				buffer.WriteString(`'>`)
				hero.EscapeHTML(v["field"], buffer)
				buffer.WriteString(`</option>
            `)
			}
			buffer.WriteString(`
        </select>
        <input type="hidden" name="`)
			hero.EscapeHTML(data.Field, buffer)
			buffer.WriteString(`[]">
        <span class="help-block">
            <i class="fa fa-info-circle"></i>&nbsp;All methods if empty
        </span>
    </div>
    <script>
        $(".`)
			hero.EscapeHTML(data.Field, buffer)
			buffer.WriteString(`").select2({
            allowClear: true,
            placeholder: "`)
			hero.EscapeHTML(data.Head, buffer)
			buffer.WriteString(`"
        });
    </script>
`)
		}
		buffer.WriteString(`
</div>
`)
	}

	buffer.WriteString(`
                        </div>
                    </div>
                    <!-- /.box-body -->
                    <div class="box-footer">
                        <input type="hidden" name="_token" value="7TEJrUaKsAIZ0qbc03G1nmeDOfmHyhCbMHqHlnkg"><div class="col-md-2">
                    </div>
                        <div class="col-md-8">

                            <div class="btn-group pull-right">
                                <button type="submit" class="btn btn-info pull-right" data-loading-text="&lt;i class='fa fa-spinner fa-spin '&gt;&lt;/i&gt; Save">Save</button>
                            </div>

                            <div class="btn-group pull-left">
                                <button type="reset" class="btn btn-warning">Reset</button>
                            </div>

                        </div>

                    </div>

                    <input type="hidden" name="_method" value="PUT" class="_method">
                    `)
	buffer.WriteString(`
<input type="hidden" name="_previous_" value='`)
	hero.EscapeHTML(previous, buffer)
	buffer.WriteString(`' class="_previous_">
<input type="hidden" name="id" value='`)
	hero.EscapeHTML(id, buffer)
	buffer.WriteString(`' class="_previous_">
`)

	buffer.WriteString(`
                    <!-- /.box-footer -->
                </form>
            </div>
        </div>
    </div>
</section>

<script data-exec-on-popstate>
    $(function () {
        $('.form-history-back').on('click', function (event) {
            event.preventDefault();
            history.back(1);
        });
    });
</script>`)

}
