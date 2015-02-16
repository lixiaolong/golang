<script type="text/javascript" src="/static/js/jstree.min.js"></script>
<script>
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
</script>
