$(document).ready( function(){
  //declare variables
  var title = $("#title");
  var description = $("#description");
  var add = $("#add");
  var postURL = "add/";
  //event handlers
  $(add).click( function(e){
    var data = "title=" + title.val() + "&description=" + description.val();
    $.ajax({
          type: "POST",
          url: postURL,
          data: data,
        success: function(responseText) {
          if(responseText == "0"){
            $(divError).show();
          }else{
            $(divSuccess).show();
            $(divReset).show()
             //$(submit).attr("disabled", true);
          }
        }
    });
    e.preventDefault();
  });
});