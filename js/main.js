$(document).ready( function(){
  //declare variables
  var title = $("#title");
  var description = $("#description");
  var add = $("#add");
  var postURL = "add/";
  var divError = $("#error");
  var divSuccess = $("#success");
  //event handlers
  $(add).click( function(e){
    var data = "title=" + title.val() + "&description=" + description.val();
    $.ajax({
          type: "POST",
          url: postURL,
          data: data,
        success: function(responseText) {
          if(responseText = "success"){
            $(divSuccess).show();            
          }else{
           $(divError).show();
             //$(submit).attr("disabled", true);
          }
        }
    });
    e.preventDefault();
  });
});