<!DOCTYPE html>
<html>

<head>
  <title>{{.ver.Title}}</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">

  <!-- 引入 Bootstrap -->
  <link href="/static/css/bootstrap.min.css" rel="stylesheet">

  <!-- HTML5 Shim 和 Respond.js 用于让 IE8 支持 HTML5元素和媒体查询 -->
  <!-- 注意： 如果通过 file://  引入 Respond.js 文件，则该文件无法起效果 -->
  <!--[if lt IE 9]>
         <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
         <script src="https://oss.maxcdn.com/libs/respond.js/1.3.0/respond.min.js"></script>
      <![endif]-->

  <link rel="stylesheet" href="/static/css/timeline.css">
</head>

<body>

  <div class="container">
    <ul class="timeline">
      <h1>{{.ver.Title}}</h1> {{range $key, $val := .ver.Contents}}
      <li>
        <span class="time">{{$val.Time}}</span>
        <span class="title">{{$val.Title}}</span>
        <pre class="content">{{$val.Details}}</pre>
      </li>{{end}}
      <li>
        <botton class="btn btn-primary btn-lg" data-toggle="modal" data-target="#verSubmit">
          新增记录
        </botton>
        <div class="modal fade" tabindex="-1" id="verSubmit" role="dialog" aria-hidden="true">
          <div class="modal-dialog">
            <div class="modal-content">
              <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span>
                </button>
                <h4 class="modal-title">新增记录</h4>
              </div>
              <div class="modal-body">
                <form action="/" method="post">
                  <div class="form-group">
                    <label for="verTitle">名称</label>
                    <input type="text" name="title" class="form-control" id="newTitle" placeholder="请输入名称">
                  </div>
                  <div class="form-group">
                    <label for="exampleInputFile">发布时间</label>
                    <input type="text" name="time" class="form-control" id="newTime" placeholder="请输入时间">
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

  <!-- jQuery (Bootstrap 的 JavaScript 插件需要引入 jQuery) -->
  <script src="/static/js/jquery-1.11.2.min.js"></script>
  <!-- 包括所有已编译的插件 -->
  <script src="/static/js/bootstrap.min.js"></script>

  <script src="/static/js/timeline.js"></script>
</body>

</html>
