# timeline

## 目标

1. 版本说明列表，timeline
2. jstree
3. chart.js

## 用到的开源库

http://jquery.com/ 
http://getbootstrap.com/
http://www.bootcss.com/p/bootstrap-datetimepicker/
http://www.jstree.com/

## 关于IE8

### 日期控件不兼容

datetimepicker日期控件在IE8下显示不了，找到一个链接

http://stackoverflow.com/questions/1744310/how-to-fix-array-indexof-in-javascript-for-internet-explorer-browsers

添加如下代码在js前面第一个函数里面，可以去掉空格和换行。

```
if (!Array.prototype.indexOf) {
    Array.prototype.indexOf = function(obj, start) {
         for (var i = (start || 0), j = this.length; i < j; i++) {
             if (this[i] === obj) { return i; }
         }
         return -1;
    }
}
```

### go语言模板将注释去掉导致IE8的特殊处理无效

通过如下方法添加注释即可：

```
{{str2html "<!--[if lt IE 9]>"}}
    <script src="/static/js/html5shiv.min.js"></script>
    <script src="/static/js/respond.min.js"></script>
{{str2html "<![endif]-->"}}
```

