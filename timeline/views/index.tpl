<!DOCTYPE html>
<html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">

  <title>{{.ver.Title}}</title>

  <!-- 引入 Bootstrap -->
  <link rel="stylesheet" href="/static/css/bootstrap.min.css">
  <link rel="stylesheet" href="/static/css/bootstrap-datetimepicker.min.css">
  <link rel="stylesheet" href="/static/css/timeline.css">

  {{str2html "<!--[if lt IE 9]>"}}
      <script src="/static/js/html5shiv.min.js"></script>
      <script src="/static/js/respond.min.js"></script>
  {{str2html "<![endif]-->"}}

  <!-- jQuery (Bootstrap 的 JavaScript 插件需要引入 jQuery) -->
  <script type="text/javascript" src="/static/js/jquery-1.11.2.min.js"></script>
  <!-- 包括所有已编译的插件 -->
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
</head>

<body>

  <div class="container">
    <nav class="navbar navbar-default">
      <div class="container-fluid">
        <div class="navbar-header">

        </div>
        <div class="collapse navbar-collapse">
          <ul class="nav navbar-nav">
            <li class="active dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">修改记录 <span class="caret"></span></a>
              <ul class="dropdown-menu" role="menu">
                <li>
                  <a data-toggle="modal" href="#verSubmit" class="btn btn-default navbar-btn center">新增</a>
                </li>
              </ul>
            </li>
            <li>
              <a href="#">资产管理</a>
            </li>
            <li>
              <a href="#">统计图表</a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  </div>

  <div class="container">
    <ul class="timeline">
      <h1>{{.ver.Title}}</h1> {{range $key, $val := .ver.Contents}}
      <li>
        <span class="time">{{$val.Time}}</span>
        <span class="title">{{$val.Title}}</span>
        <pre class="content">{{$val.Details}}</pre>
      </li>{{end}}
      <li>
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
      </li>
    </ul>
  </div>

  <script type="text/javascript" src="/static/js/bootstrap-datetimepicker.min.js"></script>
  <script type="text/javascript" src="/static/js/locales/bootstrap-datetimepicker.zh-CN.js" charset="UTF-8"></script>
  <script src="/static/js/timeline.js"></script>
</body>

</html>
