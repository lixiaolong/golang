<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>测试系统</title>
  <link rel="stylesheet" type="text/css" href="/static/bower_components/jquery-easyui/themes/default/easyui.css">
  <link rel="stylesheet" type="text/css" href="/static/bower_components/jquery-easyui/themes/icon.css">
  <style>
    #topMenu {
      float: right;
      margin-right: 10px;
      /*border: solid 1px black;*/
    }
    #productName {
      margin-bottom: 2px;
     /* border: solid 1px black;*/
    }

    div#statusLine td {
      min-width: 120px;
      text-align: center;
    }
  </style>
  <script type="text/javascript" src="/static/bower_components/jquery-easyui/jquery.min.js"></script>
  <script type="text/javascript" src="/static/bower_components/jquery-easyui/jquery.easyui.min.js"></script>
</head>
<body class="easyui-layout">

  <div data-options="region:'north',title:'欢迎使用测试系统'" style="height:110px;background:rgb(81,168,231);">
    <h1 id="productName">自动化测试系统</h1>
    <div id="topMenu">
      <table>
        <tr>
          <td><span>管理员：【admin】</span></td>
          <td>
            <a href="javascript:void(0)" class="easyui-linkbutton" data-options="plain:true,iconCls:'icon-back'" onclick="alert('首页');">首页</a>
          </td>
          <td>
            <a href="javascript:void(0)" class="easyui-linkbutton" data-options="plain:true,iconCls:'icon-reload'" onclick="alert('重启');">重启</a>
          </td>
          <td>
            <a href="javascript:void(0)" class="easyui-linkbutton" data-options="plain:true,iconCls:'icon-help'" onclick="alert('帮助');">帮助</a>
          </td>
          <td>
            <a href="javascript:void(0)" class="easyui-linkbutton" data-options="plain:true,iconCls:'icon-redo'" onclick="alert('退出');">退出</a>
          </td>
        </tr>
      </table>
    </div>
  </div>


  <div data-options="region:'south'" style="height:36px;">
    <div id="statusLine">
      <table>
        <tr>
          <td><span>测试任务</span><span> 211111 </span><span>个</span></td>
          <td><span>测试用例</span><span> 11168 </span><span>条</span></td>
          <td><span>CPU利用率</span></td>
          <td><div class="easyui-progressbar" data-options="value:80" style="width:100px;"></div></td>
          <td><span>内存利用率</span></td>
          <td><div class="easyui-progressbar" data-options="value:70" style="width:100px;"></div></td>
          <td>
            <a href="javascript:void(0)" class="easyui-linkbutton" data-options="plain:true,iconCls:'icon-reload'" onclick="alert('刷新');">刷新</a>
          </td>
        </tr>
      </table>
    </div>
  </div>


  <div data-options="region:'west',split:true" title="West" style="width:200px;">
    <div class="easyui-accordion" data-options="fit:true,border:false">
      <div title="Title1" style="padding:10px;">
        content1
      </div>
      <div title="Title2" data-options="selected:true" style="padding:10px;">
        content2
      </div>
      <div title="Title3" style="padding:10px">
        content3
      </div>
    </div>
  </div>


  <div data-options="region:'center',buttons:$('#buttons')" style="overflow:auto;">
    <div class="easyui-tabs" data-options="fit:true,border:false,plain:false">
      <div title="DataGrid" style="padding:5px">
        <table id="dg" class="easyui-datagrid" title="Custom DataGrid Pager" data-options="
          rownumbers:true,
          singleSelect:true,
          pagination:true,
          pageList: [10,20,30,40,50,100,200,400],
          url:'/static/data/datagrid_data1.json',method:'get',
          fit:true,fitColumns:true">
          <thead>
            <tr>
              <th data-options="field:'itemid',width:80">Item ID</th>
              <th data-options="field:'productid',width:100">Product</th>
              <th data-options="field:'listprice',width:80,align:'right'">List Price</th>
              <th data-options="field:'unitcost',width:80,align:'right'">Unit Cost</th>
              <th data-options="field:'attr1',width:240">Attribute</th>
              <th data-options="field:'status',width:60,align:'center'">Status</th>
            </tr>
          </thead>
        </table>
      </div>
      <div title="About" data-options="href:'/static/data/_content.html'" style="padding:10px"></div>
    </div>
  </div>

  <script>
    $(document).ready(function () {

    });
  </script>
</body>
</html>