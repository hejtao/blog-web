/**

 @Name：layui.blog 闲言轻博客模块
 @Author：徐志文
 @License：MIT
 @Site：http://www.layui.com/template/xianyan/
    
 */
layui.define(['element', 'form','laypage','jquery','laytpl'],function(exports){
  var element = layui.element
  ,form = layui.form
  ,laypage = layui.laypage
  ,$ = layui.jquery
  ,laytpl = layui.laytpl;




  //留言板提交按钮START-------------------------------------------------------
  $('#item-btn').on('click', function(){
    var elemCont = $('#LAY-msg-content') ,content = elemCont.val();
    if(content.replace(/\s/g, '') == ""){
      layer.msg('请输入留言', { icon: 7});
      return elemCont.focus();
    }

    $.post("/comment_config/save", {content:content}, function(bee_data){  //ajax 请求
          if (bee_data.code==3333) {

            var html = DrawMessage(bee_data.message);
            var $html = $(html);
            $html.find(".like").on("click", praiseStyle);

            // $('#LAY-msg-box').prepend($html); //附加到 message.html 相应的位置
            // elemCont.val('');  // 清空 textarea
            layer.msg('留言成功', { icon: 1})  
            if ($("#test1").size()>0) {
              MessagePageination() 
            }                    
                                      
              
              
          } else{
              layer.msg("留言失败: "+bee_data.msg, { icon: 5});  //json的msg
          } 

        }, "json").error(function(){
          layer.msg("网络异常", { icon: 2})
        }); 

        return false
  });
  //留言板提交按钮START-------------------------------------------------------

  //留言板分页 START----------------------------------------------------------------------------------------
  if ($("#test1").size()>0) {
    MessagePageination() 
  }

  function MessagePageination(){
    //从 beego 获取留言总数的函数的路由/comment_config/count
    var count = 0;
    $.ajax({
      url:"/comment_config/count",
      type: "GET",
      async: false, //同步调用
      success:function(bee_data){
        count = bee_data.count;
      },

      error:function(){
        layer.msg("网络异常", { icon: 2})
      }
    });

    laypage.render({

      elem: 'test1' //注意，这里的 test1 是 ID，不用加 # 号
      ,count: count //数据总数，从服务端得到
      ,theme: '#1e9fff'
      ,limit: 5

      ,jump: function(obj, first){
      //obj包含了当前分页的所有参数，比如：
      console.log(obj.curr); //得到当前页，以便向服务端请求对应页的数据。
      console.log(obj.limit); //得到每页显示的条数
      
      // 查询留言的函数的路由 /comment_config/query
      $.get("/comment_config/query", {page:obj.curr, limit:obj.limit}, function(bee_data){
        if(bee_data.code == 5555){

          var comments = bee_data.comments;
          var html = "";
          for (var i = 0; i < comments.length; i++) {
            html += DrawMessage(comments[i]);
          }
          var $html = $(html);
          $html.find(".like").on("click", praiseStyle);

          $("#LAY-msg-box").html($html);

        }else{

          layer.msg(bee_data.msg, { icon: 5});

        }

        }).error(function(){
          layer.msg("网络异常", { icon: 2});
        });
      }

    });
  }

  function DrawMessage(x){
    var view = $('#LAY-msg-tpl').html()  //获取 message.html 中相应的 html 代码

      //模版数据
      ,data = {
        name: x.Author.Name
        ,avatar: x.Author.Avatar || '/static/images/info-img.png'
        ,likes: x.Likes
        ,content: x.Content
        ,key: x.CommentKey
      };

      return laytpl(view).render(data);
      //模板渲染
      // laytpl(view).render(data, function(html){  //得到 html 代码
      //   $('#LAY-msg-box').prepend(html); //附加到 message.html 相应的位置
      //       elemCont.val('');  // 清空 textarea
      //       layer.msg('留言成功', { 
      //         icon: 1
      //       })   
      // });
  }
  //留言板分页 END---------------------------------------------------------------------------------------








  //评论提交按钮START-------------------------------------------------------
  $('#item-btn-cmt').on('click', function(){
     var elemCont = $('#LAY-msg-content') ,content = elemCont.val();
    if(content.replace(/\s/g, '') == ""){
      layer.msg('请输入评论', { icon: 7});
      return elemCont.focus();
    }

    var key = $("#noteKey").val()
    $.post("/comment_config/save/" + key, {content:content}, function(bee_data){  //ajax 请求
          if (bee_data.code == 4444) {

            var html = DrawMessage(bee_data.comment);
            var $html = $(html);
            $html.find(".like").on("click", praiseStyle);

            // $('#LAY-msg-box').prepend($html); //附加到 message.html 相应的位置
            // elemCont.val('');  // 清空 textarea
            layer.msg('评论成功', { icon: 1})   

            if ($("#test3").size()>0) {
               CommentPageination()
            }                   
                                
          } else{
              layer.msg("评论失败: "+bee_data.msg, { icon: 5});  //json的msg
          } 

        }, "json").error(function(){
          layer.msg("网络异常", { icon: 2})
        }); 
  });
  //评论提交按钮START-------------------------------------------------------



  //文章评论分页 START----------------------------------------------------------------------------------------
  if ($("#test3").size()>0) {
    CommentPageination()
  }

  function CommentPageination(){
    var count = 0;
    var key = $("#noteKey").val()  //从 note.html 获取文章的 key

    $.ajax({
      url:"/comment_config/count/" + key,
      type: "GET",
      async: false, //同步调用
      success:function(bee_data){
        count = bee_data.count;
      },

      error:function(){
        layer.msg("网络异常", { icon: 2})
      }
    });

    laypage.render({

      elem: 'test3' //注意，这里的 test1 是 ID，不用加 # 号
      ,count: count //数据总数，从服务端得到
      ,theme: '#1e9fff'
      ,limit: 3

      ,jump: function(obj, first){
      //obj包含了当前分页的所有参数，比如：
      console.log(obj.curr); //得到当前页，以便向服务端请求对应页的数据。
      console.log(obj.limit); //得到每页显示的条数
      
      $.get("/comment_config/query/" + key, {page:obj.curr, limit:obj.limit}, function(bee_data){
        if(bee_data.code == 5555){

          var comments = bee_data.comments;
          var html = "";
          for (var i = 0; i < comments.length; i++) {
            html += DrawMessage(comments[i]);
          }
          var $html = $(html);
          $html.find(".like").on("click", praiseStyle);

          $("#LAY-msg-box").html($html);

        }else{

          layer.msg(bee_data.msg, { icon: 5});

        }

        }).error(function(){
          layer.msg("网络异常", { icon: 2});
        });
      }

    });
  }

  
  //文章评论分页 END---------------------------------------------------------------------------------------



  //主页分页 START----------------------------------------------------------------------------------------
  if ($("#test2").size()>0) {

    //从 beego 获取文章总数的函数的路由/count
    var count = 0;
    var search = "";
    $.ajax({
      url:"/count",
      type: "GET",
      async: false, //同步调用
      success:function(bee_data){
        count = bee_data.count;
        search = bee_data.search;
      },

      error:function(){
        layer.msg("网络异常", { icon: 2})
      }
    });

    laypage.render({

      elem: 'test2' //注意，这里的 test2 是 ID，不用加 # 号
      ,count: count //数据总数，从服务端得到
      ,theme: '#1e9fff'
      ,limit: 5

      ,jump: function(obj, first){
      //obj包含了当前分页的所有参数，比如：
      console.log(obj.curr); //得到当前页，以便向服务端请求对应页的数据。
      console.log(obj.limit); //得到每页显示的条数
      
      
      $.get("/query", {page:obj.curr, limit:obj.limit, search:search}, function(bee_data){
        if(bee_data.code == 9999){

          var notes = bee_data.notes;
          var dates = bee_data.dates;
          var html = "";
          var j = 0;
          for (var i = 0; i < notes.length; i++) {
            j = i;
            html += DrawHome(notes[j], dates[j]);
          }

          var $html = $(html);
          $("#LAY-msg-box").html($html);

        }else{

          layer.msg(bee_data.msg, { icon: 5});

        }

        }).error(function(){
          layer.msg("网络异常", { icon: 2});
        });
      }

    });
  }

  function DrawHome(note, date){
    var view = $('#LAY-note-tpl').html()  //获取 message.html 中相应的 html 代码

     ,data ={ 
        key : note.Key,
        title : note.Title,
        summary: note.Summary,
        update_time : date
     };
      return laytpl(view).render(data);
  }
  //主页分页 END ----------------------------------------------------------------------------------------


  // start 导航显示隐藏
  
  $("#mobile-nav").on('click', function(){
    $("#pop-nav").toggle();
  });

  // end 导航显示隐藏




  //start 评论的特效
  
  (function ($) {
    $.extend({
        tipsBox: function (options) {
          options = $.extend({
            obj: null,  //jq对象，要在那个html标签上显示
            str: "+1",  //字符串，要显示的内容;也可以传一段html，如: "<b style='font-family:Microsoft YaHei;'>+1</b>"
            startSize: "12px",  //动画开始的文字大小
            endSize: "30px",    //动画结束的文字大小
            interval: 600,  //动画时间间隔
            color: "red",    //文字颜色
            callback: function () { }    //回调函数
          }, options);

          $("body").append("<span class='num'>" + options.str + "</span>");

          var box = $(".num");
          var left = options.obj.offset().left + options.obj.width() / 2;
          var top = options.obj.offset().top - 10;
          box.css({
            "position": "absolute",
            "left": left + "px",
            "top": top + "px",
            "z-index": 9999,
            "font-size": options.startSize,
            "line-height": options.endSize,
            "color": options.color
          });
          box.animate({
            "font-size": options.endSize,
            "opacity": "0",
            "top": top - parseInt(options.endSize) + "px"
          }, options.interval, function () {
            box.remove();
            options.callback();
          });
        }
      });
  })($); 

  function niceIn(prop){
    prop.find('i').addClass('niceIn');
    setTimeout(function(){
      prop.find('i').removeClass('niceIn'); 
    },1000);    
  }

  //----------------------------------------------------------------------------
  $(function () {
    $(".like").on('click', praiseStyle);
  });
  
  function praiseStyle() {
      if(!($(this).hasClass("layblog-this"))){
          var type = $(this).data("type");
          var key = $(this).data("key");
          var that = this
        $.post("/likes/" + type + "/" + key, function(bee_data){

          if (bee_data.code == 1111) {
              that.text = '已赞';
              $(that).addClass('layblog-this');
              $.tipsBox({
                  obj: $(that),
                  str: "+1",
                  callback: function () {
                  }
              });
              niceIn($(that));
              layer.msg('点赞成功', {icon: 6, time: 1000})
              $(that).find(".value").text(bee_data.likes_value)

          } else{
             if (bee_data.code = 2222) {
                $(that).addClass('layblog-this');
                layer.msg(bee_data.msg, { icon: 5});
             } else {
                layer.msg(bee_data.msg, { icon: 5});
             }
          }

        }).error(function(){
           layer.msg("网络异常", { icon: 2})
        });

        
      } 
    }


  //end 评论的特效


  // start点赞图标变身
  $('#LAY-msg-box').on('click', '.info-img', function(){
    $(this).addClass('layblog-this');
  })
  // end点赞图标变身


  // start  图片遮罩
  var layerphotos = document.getElementsByClassName('layer-photos-demo');
  for(var i = 1;i <= layerphotos.length;i++){
    layer.photos({
      photos: ".layer-photos-demo"+i+""
      ,anim: 0
    }); 
  }
  // end 图片遮罩


  //输出test接口
  exports('blog', {}); 
});  