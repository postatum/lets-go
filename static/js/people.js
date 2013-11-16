(function($) {
  var $json_container = $('.json-container');
  var $error = $('.error');

  function updatePage() {
      $json_container.empty();

      $.getJSON("/api/people", function(data) {
      if (data['people']) {
        var elements = [];
        $.each(data['people'], function(ind, el) {
          var rating = el.rating || 0;
          var text = el.name + ' ' + el.email + ' (' + rating + ') ';
          var $html_el = $('<li/>').html(text);
          var $like = $('<a/>').attr({
              'data-pid': el._id,
              'href': "#"
            }).addClass('js-like').text('like');
          $html_el.append($like);
          elements.push($html_el);
        });
        $json_container.append(elements);
      } else {
        $json_container.html('Ewps c:');
      }
    });

    setTimeout(function(){ updatePage(); }, 5000);
  }

  updatePage();

  $(document).on('click', '.js-like', function(e) {
    var $this = $(this);
    e.preventDefault();
    $.ajax({
      type: "POST",
      dataType: "json",
      url: "/api/people/like",
      data: {'pid': $this.data('pid')},
      success: function(data) {
        if (!data['success']) {
          $error.html(data['err']);
        }
      }
    });
  });

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
