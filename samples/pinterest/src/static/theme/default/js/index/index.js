$(function() {
    var $container = $('#container');
    $container.imagesLoaded(function() {
        $container.masonry({
            itemSelector: '.box',
            columnWidth: 100
        });
    });

    $container.infinitescroll({
        navSelector: '#page-nav',
        // selector for the paged navigation 
        nextSelector: '#page-nav a',
        // selector for the NEXT link (to page 2)
        itemSelector: '.box',
        // selector for all items you'll retrieve
        loading: {
            finishedMsg: '无更多的页面被加载',
            msgText: "<em>加载中,请稍等...</em>",
            img: '/static/theme/default/img/global/loading.gif'
        }
    },
    // trigger Masonry as a callback

    function(newElements) {
        // hide new items while they are loading
        var $newElems = $(newElements).css({
            opacity: 0
        });
        // ensure that images load before adding to masonry layout
        $newElems.imagesLoaded(function() {
            // show elems now they're ready
            $newElems.animate({
                opacity: 1
            });
            $container.masonry('appended', $newElems, true);
        });

        $(".gallery a[rel^='prettyPhoto']").prettyPhoto();
    });

    $(".gallery a[rel^='prettyPhoto']").prettyPhoto();

    $("#submit").bind("click",function() {
        function getExt(s) {
            pos = s.lastIndexOf(".");
            return s.substring(pos, s.length);
        }

        var file = $("#file").val();
        if (file == "") {
            alert("请选择文件后上传!");
            return false;
        }

        var ext = getExt(file).toLowerCase();
        if (ext != ".jpg" && ext != ".png" && ext != ".gif") {
            alert("请选择图片进行上传!");
            return false;
        }

        var classId = $("#class_select").val();
        if (classId == "0") {
            alert("请选择图片种类!");
            return false;
        }
    });

    $(".expand").click(function(){
        if($("#upload").html()=="+"){
            $("#tip").hide();
            $("#uploadForm").css({"display":"inline-block"});
            $("#upload_block").css({"border-color":"#00a0de","color":"#00a0de"});
            $("#upload").html("-");
        }else{
            $("#uploadForm").hide();
            $("#tip").css({"display":"inline-block"});
            $("#upload_block").css({"border-color":"#D8D5D2","color":"#858585"});
            $("#upload").html("+");
            $("#file").val("");
        }
    });

    $("#file").change(function(){
        if($("#file").val()!=""){
            $("#filee").html($("#file").val());
        }else{
            $("#filee").html("选择文件");
        }
    });
});