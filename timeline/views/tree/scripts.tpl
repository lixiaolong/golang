<script type="text/javascript" src="/static/js/jstree.min.js"></script>
<script>

$(document).ready(function(){
  var curmenu = $("#menulist > ul > li").eq(1);
  var curtext = curmenu.children('a').text() + ' <span class="sr-only">(current2)</span>';
  curmenu.attr("class","active");
  curmenu.children('a').attr("href","#");
  curmenu.children('a').html(curtext);

  $(function() {
  $("#jsTree").jstree({
    "core": {
      "check_callback" : true
    },
    "checkbox": {
      "keep_selected_style": false
    },
    "state" : { "key" : "demo2" },
    "plugins": ["dnd", "contextmenu", "state", "unique"]
  });
});
});
</script>