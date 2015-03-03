<div class="container">
    <ul class="timeline">
      <h1>{{.ver.Title}}</h1>{{range $key, $val := .ver.Contents}}
      <li>
        <span class="time">{{$val.Time}}</span>
        <span class="title">{{$val.Title}}</span>
        <div class="content">{{$val.Details | markdown | str2html}}</div>
      </li>{{end}}
    </ul>
  </div>

  <div class="modal fade" tabindex="-1" id="verSubmit" role="dialog" aria-hidden="false">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span>
          </button>
          <h4 class="modal-title">新增</h4>
        </div>
        <div class="modal-body">
          <form action="/" method="post">
            <div class="form-group">
              <label for="verTitle">名称</label>
              <input type="text" name="title" class="form-control" id="newTitle" placeholder="请输入名称">
            </div>
            <div class="form-group">
              <label for="newTime">发布时间</label>
              <input type="text" name="time" id="newTime">
            </div>
            <div class="form-group">
              <label for="verDetails">修改记录</label>
              <textarea name="details" class="form-control" id="newDetails" rows="10" placeholder="请输入内容"></textarea>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
              <button type="submit" class="btn btn-primary">提交</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>