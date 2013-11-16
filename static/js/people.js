(function($) {
  var $json_container = $('.json-container');
  var $error = $('.error');

  function updatePage() {
      $json_container.empty();

      $.getJSON("/api/people", function(data) {
      if (data['people']) {
        var elemelts = [];
        $.each(data['people'], function(ind, el) {
          var text = el.name + ' ' + el.email;
          var $html_el = $('<li/>').html(text);
          elemelts.push($html_el);
        });
        $json_container.append(elemelts);
      } else {
        $json_container.html('Ewps c:');
      }
    });

    setTimeout(function(){ updatePage(); }, 3000);
  }

  updatePage();

  $('#js-add-person').submit(function(e) {
    e.preventDefault();
    $error.empty();
    var $this = $(this);
    var url = $this.attr('action');
    $.ajax({
      type: "POST",
      dataType: "json",
      url: url,
      data: $this.serialize(),
      success: function(data) {
        if (!data['success']) {
          $error.html(data['err']);
        } else {
          $this.find("input").val('');
        }
      }
    });
  });

})(jQuery);
