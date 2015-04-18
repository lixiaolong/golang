<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>测试系统</title>
  <link rel="stylesheet" type="text/css" href="/static/bower_components/jquery-easyui/themes/default/easyui.css">
  <link rel="stylesheet" type="text/css" href="/static/bower_components/jquery-easyui/themes/icon.css">
  <style>
    div.easyui-tab {
      padding:20px;overflow:auto;
    }

    div#topMenu {
      left: 200px;
      top: 0px;
      height: 28px;
      position: absolute;
      /*border: solid 1px black;*/
      z-index: 10;
      overflow: hidden;
    }

    div#statusLine>* {
      display: inline-block;
      margin-right: 10px;
    }
  </style>
  <script type="text/javascript" src="/static/bower_components/jquery-easyui/jquery-1.7.2.min.js"></script>
  <script type="text/javascript" src="/static/bower_components/jquery-easyui/jquery.easyui.min.js"></script>
</head>
<body class="easyui-layout">
  <div id="topMenu">
    <span>管理员：admin</span>
    <a href="javascript:void(0)" class="easyui-linkbutton" data-options="plain:true,iconCls:'icon-back'" onclick="alert('首页');">首页</a>
    <a href="javascript:void(0)" class="easyui-linkbutton" data-options="plain:true,iconCls:'icon-reload'" onclick="alert('重启');">重启</a>
    <a href="javascript:void(0)" class="easyui-linkbutton" data-options="plain:true,iconCls:'icon-help'" onclick="alert('帮助');">帮助</a>
    <a href="javascript:void(0)" class="easyui-linkbutton" data-options="plain:true,iconCls:'icon-redo'" onclick="alert('退出');">退出</a>
  </div>
  <div data-options="region:'north'" title="欢迎使用测试系统" style="height:100px;background:rgb(81,168,231);">
    <h1>自动化测试系统</h1>
  </div>
  <div data-options="region:'west',split:true" title="菜单" style="width:200px;padding:1px;overflow:hidden;">
    <div class="easyui-accordion" data-options="fit:true,border:false">
      <div title="基本配置">
        <p>菜单1</p>
        <p>菜单2</p>
      </div>
      <div title="高级配置">
        <p>菜单1</p>
        <p>菜单2</p>
      </div>
    </div>
  </div>
  <div data-options="region:'center'" style="overflow:auto;">
    <div class="easyui-tabs" data-options="fit:true,border:false">
      <div title="标签页1" class="easyui-tab">
      </div>
      <div title="标签页2" class="easyui-tab">
      </div>
      <div title="标签页3" class="easyui-tab">
      </div>
    </div>
  </div>
  <div data-options="region:'south'" style="height:28px;overflow:hidden;">
    <div id="statusLine">
      <span>CPU利用率</span><div id="cpuUsage" class="easyui-progressbar" style="width:100px;"></div>
      <span>内存利用率</span><div id="memUsage" class="easyui-progressbar" style="width:100px;"></div>
      <a href="javascript:void(0)" class="easyui-linkbutton" data-options="plain:true,iconCls:'icon-reload'" onclick="alert('刷新');">刷新</a>
    </div>
  </div>
  <script>
    $(document).ready(function () {
      $('#memUsage').progressbar('setValue', 80);
      $('#cpuUsage').progressbar('setValue', 80);
    });
  </script>
</body>
</html>