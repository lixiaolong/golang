$(document).ready(function() {
  var tl_title = $("ul.timeline li span.title");

  $("ul.timeline h1").click(function() {
    tl_title.next().toggle();
  });

  tl_title.click(function() {
    $(this).next().toggle(400);
  });
});
