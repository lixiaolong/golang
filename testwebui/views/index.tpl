<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Window Layout - jQuery EasyUI Demo</title>
  <link rel="stylesheet" type="text/css" href="/static/bower_components/jquery-easyui/themes/gray/easyui.css">
  <link rel="stylesheet" type="text/css" href="/static/bower_components/jquery-easyui/themes/icon.css">
  <script type="text/javascript" src="/static/bower_components/jquery-easyui/jquery-1.7.2.min.js"></script>
  <script type="text/javascript" src="/static/bower_components/jquery-easyui/jquery.easyui.min.js"></script>
</head>
<body>
  <h2>Window Layout</h2>
  <p>Using layout on window.</p>
  <div style="margin:20px 0;">
    <a href="javascript:void(0)" class="easyui-linkbutton" onclick="$('#w').window('open')">Open</a>
    <a href="javascript:void(0)" class="easyui-linkbutton" onclick="$('#w').window('close')">Close</a>
  </div>
  <div id="w" class="easyui-window" title="Window Layout" style="width:800px;height:600px;">
    <div class="easyui-layout" data-options="fit:true">
      <div data-options="region:'center'">
        <div class="easyui-tabs" fit="true">
          <div title="基本配置" style="padding:10px">
            <form action="">
              <input type="text" />
              <input type="text" />
              <input type="text" />
            </form>
          </div>
          <div title="高级配置" style="padding:10px">
            <ul class="easyui-tree" data-options="url:'tree_data1.json',method:'get',animate:true"></ul>
          </div>
          <div title="端口配置" data-options="closable:true" style="padding:10px">
          </div>
        </div>
      </div>
      <div data-options="region:'east',split:true" style="width:200px;">
      </div>
      <div data-options="region:'south',border:false" style="text-align:right;padding:5px 0 0;">
        <a class="easyui-linkbutton" data-options="iconCls:'icon-ok'" href="javascript:void(0)" onclick="javascript:alert('ok')" style="width:80px">Ok</a>
        <a class="easyui-linkbutton" data-options="iconCls:'icon-cancel'" href="javascript:void(0)" onclick="javascript:alert('cancel')" style="width:80px">Cancel</a>
      </div>
    </div>
  </div>

</body>
</html>