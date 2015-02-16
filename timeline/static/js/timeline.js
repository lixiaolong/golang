$(document).ready(function() {
  var tl_title = $("ul.timeline li span.title");

  $("ul.timeline h1").click(function() {
    tl_title.next().toggle();
  });

  tl_title.click(function() {
    $(this).next().toggle(400);
  });

  // 要先引用datapicker
  $("#newTime").datetimepicker({
    format: 'yyyy-mm-dd',
    autoclose: true,
    minView: 2,
    todayBtn: "linked",
    todayHighlight: true,
    language: "zh-CN",
  });
});
