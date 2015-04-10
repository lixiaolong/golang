<script type="text/javascript" src="/static/js/bootstrap-datetimepicker.min.js"></script>
  <script type="text/javascript" src="/static/js/locales/bootstrap-datetimepicker.zh-CN.js" charset="UTF-8"></script>
  <script src="/static/js/timeline.js"></script>
  <script>
    $(document).ready(function(){
      var curmenu = $("#menulist > ul > li").eq(0);
      var curtext = curmenu.children('a').text() + ' <span class="sr-only">(current2)</span>';
      curmenu.attr("class","active");
      curmenu.children('a').attr("href","#");
      curmenu.children('a').html(curtext);
    });
  </script>