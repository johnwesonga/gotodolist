$(document).ready( function(){
  //declare variables
  var title = $("#title");
  var description = $("#description");
  var add = $("#add");
  var remove = $("#remove");
  var postURL = "add/";
  var removeURL = "/delete/"+encodeURIComponent(id);
  var divError = $("#error");
  var divSuccess = $("#success");
  //event handlers
  $(add).click( function(e){
     e.preventDefault();
    var data = "title=" + title.val() + "&description=" + description.val();
    $.ajax({
          type: "POST",
          url: postURL,
          data: data,
        success: function(responseText) {
          if(responseText = "success"){
            $(divSuccess).show();     
            title.val() == " ";
            description.val() == " ";       
          }else{
           $(divError).show();
             //$(submit).attr("disabled", true);
          }
        }
    });
   
  });
 
  
});