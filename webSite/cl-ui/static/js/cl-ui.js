$('.cl-menu-boot').on('click', function () {
  let $ele = $(this);
  if ($ele.hasClass('cl-menu-show')) {
    $ele.removeClass('cl-menu-show');
  } else {
    $ele.addClass('cl-menu-show');
  }
  $ele.next().animate({
    height: 'toggle'
  });
});

const $menuChild = $('.cl-menu-child-item'); // 菜单的列表
$menuChild.on('click', function () {
  let $ele = $(this);
  if (!$ele.hasClass('active')) {
    restart('.cl-menu-child-item', $ele.children().text(), () => {
      $('#view').attr('src', $ele.children('span').attr('data-href'));
      addTitleBarItem($ele);
    });
  }
});

clickTitleBar();

function addTitleBarItem (node) {
  const title = node.children().text();
  const href = node.children().attr('data-href');
  let isExit = false;
  let $titleBarItem = $('.cl-title-bar-item');
  $titleBarItem.each((item, ele) => {
    if ($(ele).children('span').text() === title) {
      isExit = true;
      return;
    }
  })
  if (!isExit) {
    $('.cl-title-bar-ul').append(`<li class="cl-title-bar-item active"><span data-href="${href}">${title}</span><i class="glyphicon glyphicon-remove"></i></li>`);
  }
  restart('.cl-title-bar-item', title);
  clickTitleBar();
}


function clickTitleBar () {
  const $titleBarItem = $('.cl-title-bar-item');
  $titleBarItem.on('click', function () {
    const $ele = $(this);
    const title = $ele.children('span').text();
    restart('.cl-menu-child-item', title, ($ele) => {
      if ($ele !== '首页') {
        if ($ele.parent().css('display') === 'none') {
          $ele.parent().animate({
            height: 'toggle'
          });
        }
      }
    });
    restart('.cl-title-bar-item', title, (ele) => {
      $('#view').attr('src', ele.children('span').attr('data-href'));
    });
  });
  const $delete = $titleBarItem.children('i');
  $delete.on('click', function (e) {
    e.preventDefault();
    e.stopPropagation();
    restart('.cl-menu-child-item');
    const $ele = $(this);
    $ele.parent().remove();
    restart('.cl-title-bar-item', '首页', () => {
      $('#view').attr('src', 'https://www.luxuewen.com');
    });
  });
}

// 重置active
function restart (classes, title, callback) {
  const $active = $(classes);
  let nowActive = '首页'; 
  $active.each((index, ele) => {
    const $ele = $(ele);
    $ele.removeClass('active');
    if (title === $ele.children('span').text()) {
      $ele.addClass('active');
      nowActive = $ele;
    }
  });
  callback && callback(nowActive);
}